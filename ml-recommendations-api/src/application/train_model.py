class TrainModelUseCase:
    def __init__(self, encoder, model_repository):
        self.encoder = encoder
        self.model_repository = model_repository

    def execute(self, users, books):
        context = self.encoder.build_context(users, books)
        xs, ys, input_dim = self.encoder.create_training_data(users, books, context)
        if not xs:
            return {"trained": False, "reason": "no training data"}

        model = self.model_repository.build_and_train(xs, ys, input_dim=input_dim)
        self.model_repository.save(model)
        self.model_repository.save_context(context)
        return {"trained": True, "samples": len(xs), "input_dim": input_dim, "context": context}
