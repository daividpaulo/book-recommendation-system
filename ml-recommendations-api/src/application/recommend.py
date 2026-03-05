class RecommendUseCase:
    def __init__(self, encoder, model_repository, vector_store):
        self.encoder = encoder
        self.model_repository = model_repository
        self.vector_store = vector_store

    def execute(self, user, books, context):
        model = self.model_repository.load()
        if model is None:
            raise ValueError("model not trained yet")

        user_vector = self.encoder.encode_user(user, context)
        user_book_query_vector = self.encoder.encode_user_book_query(user, context)
        candidate_ids = set(self.vector_store.search_books(user_book_query_vector, limit=200))
        candidate_books = [book for book in books if book.id in candidate_ids]
        if not candidate_books:
            candidate_books = books

        xs = []
        for book in candidate_books:
            book_vector = self.encoder.encode_book(book, context)
            pair_vector = self.encoder.encode_user_book_pair(user, book)
            xs.append([*user_vector, *book_vector, *pair_vector])

        scores = self.model_repository.predict(model, xs)
        ranking = [
            {"book_id": book.id, "title": book.title, "score": float(scores[i])}
            for i, book in enumerate(candidate_books)
        ]
        ranking.sort(key=lambda item: item["score"], reverse=True)
        return ranking
