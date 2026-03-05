from __future__ import annotations

import io
import json
import os
from urllib.parse import urlparse

import numpy as np
import tensorflow as tf
from minio import Minio
from minio.error import S3Error


class TensorFlowModelRepository:
    def __init__(self, path: str, minio_endpoint: str, minio_access_key: str, minio_secret_key: str, minio_bucket: str):
        self.path = path
        os.makedirs(path, exist_ok=True)
        self.model_file = os.path.join(path, "book-recommendation.keras")
        self.context_file = os.path.join(path, "encoding-context.json")
        self.model_object = "book-recommendation.keras"
        self.context_object = "encoding-context.json"
        self.minio_bucket = minio_bucket
        self.minio_client = self._build_minio_client(
            minio_endpoint=minio_endpoint,
            minio_access_key=minio_access_key,
            minio_secret_key=minio_secret_key,
        )
        self._ensure_bucket()

    def build_and_train(self, xs, ys, input_dim: int):
        x_tensor = tf.convert_to_tensor(xs, dtype=tf.float32)
        y_tensor = tf.convert_to_tensor(ys, dtype=tf.float32)

        model = tf.keras.Sequential([
            tf.keras.layers.Input(shape=(input_dim,)),
            tf.keras.layers.Dense(units=64, activation="relu"),
            tf.keras.layers.Dense(units=32, activation="relu"),
            tf.keras.layers.Dense(units=16, activation="relu"),
            tf.keras.layers.Dense(units=1, activation="sigmoid"),
        ])
        model.compile(
            optimizer=tf.keras.optimizers.Adam(0.001),
            loss=tf.keras.losses.BinaryCrossentropy(label_smoothing=0.05),
            metrics=["accuracy"],
        )
        early_stopping = tf.keras.callbacks.EarlyStopping(
            monitor="loss",
            patience=6,
            restore_best_weights=True,
        )
        model.fit(
            x_tensor,
            y_tensor,
            epochs=40,
            batch_size=16,
            shuffle=True,
            verbose=0,
            callbacks=[early_stopping],
        )
        return model

    def save(self, model) -> None:
        model.save(self.model_file)
        self.minio_client.fput_object(self.minio_bucket, self.model_object, self.model_file)

    def load(self):
        if not self._object_exists(self.model_object):
            return None
        self.minio_client.fget_object(self.minio_bucket, self.model_object, self.model_file)
        return tf.keras.models.load_model(self.model_file)

    def predict(self, model, xs):
        x_tensor = tf.convert_to_tensor(xs, dtype=tf.float32)
        scores = model.predict(x_tensor, verbose=0)
        return np.asarray(scores).flatten().tolist()

    def save_context(self, context: dict) -> None:
        payload = json.dumps(context).encode("utf-8")
        self.minio_client.put_object(
            self.minio_bucket,
            self.context_object,
            io.BytesIO(payload),
            length=len(payload),
            content_type="application/json",
        )

    def load_context(self) -> dict | None:
        if not self._object_exists(self.context_object):
            return None
        response = self.minio_client.get_object(self.minio_bucket, self.context_object)
        try:
            return json.loads(response.read().decode("utf-8"))
        finally:
            response.close()
            response.release_conn()

    def _ensure_bucket(self) -> None:
        if not self.minio_client.bucket_exists(self.minio_bucket):
            self.minio_client.make_bucket(self.minio_bucket)

    def _object_exists(self, object_name: str) -> bool:
        try:
            self.minio_client.stat_object(self.minio_bucket, object_name)
            return True
        except S3Error as exc:
            if exc.code in {"NoSuchKey", "NoSuchObject", "NoSuchBucket"}:
                return False
            raise

    def _build_minio_client(self, minio_endpoint: str, minio_access_key: str, minio_secret_key: str) -> Minio:
        parsed = urlparse(minio_endpoint)
        if parsed.scheme:
            endpoint = parsed.netloc
            secure = parsed.scheme == "https"
        else:
            endpoint = minio_endpoint
            secure = False
        return Minio(
            endpoint,
            access_key=minio_access_key,
            secret_key=minio_secret_key,
            secure=secure,
        )
