from __future__ import annotations

from fastapi import FastAPI, HTTPException
from pydantic import BaseModel

from config import (
    MINIO_ACCESS_KEY,
    MINIO_BUCKET,
    MINIO_ENDPOINT,
    MINIO_SECRET_KEY,
    MODEL_PATH,
    QDRANT_HOST,
    QDRANT_PORT,
)
from src.application.embed_book import EmbedBookUseCase
from src.application.embed_user import EmbedUserUseCase
from src.application.recommend import RecommendUseCase
from src.application.train_model import TrainModelUseCase
from src.domain.entities.book import Book
from src.domain.entities.user import User
from src.infrastructure.encoders.tensorflow_encoder import TensorFlowEncoder
from src.infrastructure.model_repository.tensorflow_model import TensorFlowModelRepository
from src.infrastructure.vector_store.qdrant_store import QdrantStore, VectorDimensionMismatchError


class UserSchema(BaseModel):
    id: str
    name: str
    age: int
    profession: str = ""
    interest_areas: list[str] = []
    purchase_count: int = 0
    purchased_book_ids: list[str] = []


class BookSchema(BaseModel):
    id: str
    title: str
    author: str = ""
    category: str = ""
    subject: str = ""
    area: str = ""
    description: str = ""


class TrainRequest(BaseModel):
    users: list[UserSchema]
    books: list[BookSchema]


class RecommendRequest(BaseModel):
    user_id: str


app = FastAPI(title="ml-recommendations-api")

encoder = TensorFlowEncoder()
model_repository = TensorFlowModelRepository(
    path=MODEL_PATH,
    minio_endpoint=MINIO_ENDPOINT,
    minio_access_key=MINIO_ACCESS_KEY,
    minio_secret_key=MINIO_SECRET_KEY,
    minio_bucket=MINIO_BUCKET,
)
vector_store = QdrantStore(host=QDRANT_HOST, port=QDRANT_PORT)

train_usecase = TrainModelUseCase(encoder=encoder, model_repository=model_repository)
embed_user_usecase = EmbedUserUseCase(encoder=encoder, vector_store=vector_store)
embed_book_usecase = EmbedBookUseCase(encoder=encoder, vector_store=vector_store)
recommend_usecase = RecommendUseCase(
    encoder=encoder,
    model_repository=model_repository,
    vector_store=vector_store,
)

# In-memory registry (bootstrap scaffold).
USERS: dict[str, User] = {}
BOOKS: dict[str, Book] = {}
DEFAULT_CONTEXT: dict = {
    "min_age": 0,
    "max_age": 100,
    "min_purchase_count": 0,
    "max_purchase_count": 1,
    "professions": [],
    "interests": [],
    "areas": [],
    "categories": [],
    "subjects": [],
}
CURRENT_CONTEXT: dict = model_repository.load_context() or DEFAULT_CONTEXT.copy()


@app.get("/health")
def health():
    return {
        "status": "ok",
        "service": "ml-recommendations-api",
        "model_trained": model_repository.load() is not None,
        "context_loaded": bool(CURRENT_CONTEXT.get("areas") or CURRENT_CONTEXT.get("interests")),
    }


@app.post("/embed/users")
def embed_user(payload: UserSchema):
    user = User(**payload.model_dump())
    USERS[user.id] = user
    _refresh_context()
    vector = embed_user_usecase.execute(user, CURRENT_CONTEXT)
    return {"status": "embedded", "entity": "user", "id": user.id, "vector_dim": len(vector)}


@app.post("/embed/books")
def embed_book(payload: BookSchema):
    book = Book(**payload.model_dump())
    BOOKS[book.id] = book
    _refresh_context()
    vector = embed_book_usecase.execute(book, CURRENT_CONTEXT)
    return {"status": "embedded", "entity": "book", "id": book.id, "vector_dim": len(vector)}


@app.post("/train")
def train(payload: TrainRequest):
    global CURRENT_CONTEXT
    users = [User(**u.model_dump()) for u in payload.users]
    books = [Book(**b.model_dump()) for b in payload.books]
    USERS.update({u.id: u for u in users})
    BOOKS.update({b.id: b for b in books})

    result = train_usecase.execute(users=users, books=books)
    if result.get("trained"):
        CURRENT_CONTEXT = result.get("context", CURRENT_CONTEXT)
        _reindex_vectors()
    return result


@app.post("/recommend")
def recommend(payload: RecommendRequest):
    user = USERS.get(payload.user_id)
    if not user:
        raise HTTPException(status_code=404, detail="user not found")
    books = list(BOOKS.values())
    if not books:
        return {"recommendations": []}

    if CURRENT_CONTEXT == DEFAULT_CONTEXT:
        loaded_context = model_repository.load_context()
        if loaded_context:
            CURRENT_CONTEXT.update(loaded_context)
        else:
            _refresh_context()

    try:
        ranking = recommend_usecase.execute(user=user, books=books, context=CURRENT_CONTEXT)
    except VectorDimensionMismatchError:
        # When vector dimensions evolve, rebuild index vectors and retry.
        _reindex_vectors()
        ranking = recommend_usecase.execute(user=user, books=books, context=CURRENT_CONTEXT)
    except ValueError as exc:
        raise HTTPException(status_code=400, detail=str(exc)) from exc
    return {"user_id": payload.user_id, "recommendations": ranking}


def _refresh_context():
    global CURRENT_CONTEXT
    CURRENT_CONTEXT = encoder.build_context(list(USERS.values()), list(BOOKS.values()))


def _reindex_vectors():
    for user in USERS.values():
        embed_user_usecase.execute(user, CURRENT_CONTEXT)
    for book in BOOKS.values():
        embed_book_usecase.execute(book, CURRENT_CONTEXT)
