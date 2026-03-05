from __future__ import annotations

import json
import os

import numpy as np
import tensorflow as tf


class TensorFlowModelRepository:
    def __init__(self, path: str):
        self.path = path
        os.makedirs(path, exist_ok=True)
        self.model_file = os.path.join(path, "book-recommendation.keras")
        self.context_file = os.path.join(path, "encoding-context.json")

    def build_and_train(self, xs, ys, input_dim: int):
        x_tensor = tf.convert_to_tensor(xs, dtype=tf.float32)
        y_tensor = tf.convert_to_tensor(ys, dtype=tf.float32)

        model = tf.keras.Sequential([
            tf.keras.layers.Dense(input_shape=[input_dim], units=128, activation="relu"),
            tf.keras.layers.Dense(units=64, activation="relu"),
            tf.keras.layers.Dense(units=32, activation="relu"),
            tf.keras.layers.Dense(units=1, activation="sigmoid"),
        ])
        model.compile(
            optimizer=tf.keras.optimizers.Adam(0.01),
            loss="binary_crossentropy",
            metrics=["accuracy"],
        )
        model.fit(x_tensor, y_tensor, epochs=100, batch_size=32, shuffle=True, verbose=0)
        return model

    def save(self, model) -> None:
        model.save(self.model_file)

    def load(self):
        if not os.path.exists(self.model_file):
            return None
        return tf.keras.models.load_model(self.model_file)

    def predict(self, model, xs):
        x_tensor = tf.convert_to_tensor(xs, dtype=tf.float32)
        scores = model.predict(x_tensor, verbose=0)
        return np.asarray(scores).flatten().tolist()

    def save_context(self, context: dict) -> None:
        with open(self.context_file, "w", encoding="utf-8") as file:
            json.dump(context, file)

    def load_context(self) -> dict | None:
        if not os.path.exists(self.context_file):
            return None
        with open(self.context_file, "r", encoding="utf-8") as file:
            return json.load(file)
