import os


PORT = int(os.getenv("PORT", "5000"))
MODEL_PATH = os.getenv("MODEL_PATH", "/app/saved_models")
QDRANT_HOST = os.getenv("QDRANT_HOST", "qdrant")
QDRANT_PORT = int(os.getenv("QDRANT_PORT", "6333"))
