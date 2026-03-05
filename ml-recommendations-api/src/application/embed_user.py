from src.domain.entities.user import User


class EmbedUserUseCase:
    def __init__(self, encoder, vector_store):
        self.encoder = encoder
        self.vector_store = vector_store

    def execute(self, user: User, context: dict) -> list[float]:
        vector = self.encoder.encode_user(user, context)
        self.vector_store.upsert_user(user.id, vector, payload=user.__dict__)
        return vector
