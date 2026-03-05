from __future__ import annotations

import zlib

from qdrant_client import QdrantClient
from qdrant_client.http.exceptions import UnexpectedResponse
from qdrant_client.models import Distance, PointStruct, VectorParams


class VectorDimensionMismatchError(RuntimeError):
    pass


class QdrantStore:
    def __init__(self, host: str, port: int):
        self.client = QdrantClient(host=host, port=port)

    def _ensure_collection(self, collection_name: str, size: int) -> None:
        collections = self.client.get_collections().collections
        names = {c.name for c in collections}
        if collection_name not in names:
            self.client.create_collection(
                collection_name=collection_name,
                vectors_config=VectorParams(size=size, distance=Distance.COSINE),
            )
            return

        current_size = self._get_collection_size(collection_name)
        if current_size != size:
            # Feature changes can alter vector dimensions. Recreate collection so
            # runtime does not fail with dimension mismatch.
            self.client.delete_collection(collection_name=collection_name)
            self.client.create_collection(
                collection_name=collection_name,
                vectors_config=VectorParams(size=size, distance=Distance.COSINE),
            )

    def _get_collection_size(self, collection_name: str) -> int:
        info = self.client.get_collection(collection_name=collection_name)
        vectors = info.config.params.vectors
        if hasattr(vectors, "size"):
            return int(vectors.size)
        if isinstance(vectors, dict):
            first_value = next(iter(vectors.values()), None)
            if first_value and hasattr(first_value, "size"):
                return int(first_value.size)
        return 0

    def upsert_user(self, user_id: str, vector: list[float], payload: dict) -> None:
        self._ensure_collection("users", len(vector))
        self.client.upsert(
            collection_name="users",
            points=[PointStruct(id=self._point_id(user_id), vector=vector, payload=payload)],
        )

    def upsert_book(self, book_id: str, vector: list[float], payload: dict) -> None:
        self._ensure_collection("books", len(vector))
        self.client.upsert(
            collection_name="books",
            points=[PointStruct(id=self._point_id(book_id), vector=vector, payload=payload)],
        )

    def search_books(self, query_vector: list[float], limit: int = 50) -> list[str]:
        self._ensure_collection("books", len(query_vector))
        try:
            results = self.client.search(
                collection_name="books",
                query_vector=query_vector,
                limit=limit,
            )
        except UnexpectedResponse as exc:
            if "Vector dimension error" in str(exc):
                raise VectorDimensionMismatchError(str(exc)) from exc
            raise
        return [str(item.payload.get("id", item.id)) for item in results]

    def _point_id(self, raw_id: str) -> int:
        # Qdrant point id must be uint or UUID. Use stable uint from string id.
        return zlib.crc32(raw_id.encode("utf-8"))
