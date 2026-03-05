from src.domain.entities.book import Book


class EmbedBookUseCase:
    def __init__(self, encoder, vector_store):
        self.encoder = encoder
        self.vector_store = vector_store

    def execute(self, book: Book, context: dict) -> list[float]:
        vector = self.encoder.encode_book(book, context)
        self.vector_store.upsert_book(book.id, vector, payload=book.__dict__)
        return vector
