from __future__ import annotations

from src.domain.entities.book import Book
from src.domain.entities.user import User


def _normalize(value: float, min_value: float, max_value: float) -> float:
    denominator = max(max_value - min_value, 1e-9)
    return (value - min_value) / denominator


class TensorFlowEncoder:
    def build_context(self, users: list[User], books: list[Book]) -> dict:
        professions = sorted({u.profession for u in users if u.profession})
        interests = sorted({i for u in users for i in u.interest_areas})
        areas = sorted({b.area for b in books if b.area})
        categories = sorted({b.category for b in books if b.category})
        subjects = sorted({b.subject for b in books if b.subject})
        ages = [u.age for u in users] or [0, 1]
        purchase_counts = [u.purchase_count for u in users] or [0, 1]

        return {
            "min_age": min(ages),
            "max_age": max(ages),
            "min_purchase_count": min(purchase_counts),
            "max_purchase_count": max(purchase_counts),
            "professions": professions,
            "interests": interests,
            "areas": areas,
            "categories": categories,
            "subjects": subjects,
        }

    def encode_user(self, user: User, context: dict) -> list[float]:
        min_age = context.get("min_age", 0)
        max_age = context.get("max_age", 100)
        min_purchase_count = context.get("min_purchase_count", 0)
        max_purchase_count = context.get("max_purchase_count", 1)
        professions = context.get("professions", [])
        interests = context.get("interests", [])

        age_norm = _normalize(user.age, min_age, max_age)
        purchase_count_norm = _normalize(
            float(user.purchase_count),
            float(min_purchase_count),
            float(max_purchase_count),
        )
        profession_idx = professions.index(user.profession) if user.profession in professions else 0
        profession_norm = _normalize(profession_idx, 0, max(len(professions) - 1, 1))
        interests_one_hot = [1.0 if i in user.interest_areas else 0.0 for i in interests]
        return [age_norm, purchase_count_norm, profession_norm, *interests_one_hot]

    def encode_book(self, book: Book, context: dict) -> list[float]:
        area_hot = [1.0 if a == book.area else 0.0 for a in context["areas"]]
        category_hot = [1.0 if c == book.category else 0.0 for c in context["categories"]]
        subject_hot = [1.0 if s == book.subject else 0.0 for s in context["subjects"]]
        return [*area_hot, *category_hot, *subject_hot]

    def encode_user_book_query(self, user: User, context: dict) -> list[float]:
        interests = set(user.interest_areas)
        area_hot = [1.0 if a in interests else 0.0 for a in context["areas"]]
        category_hot = [1.0 if c in interests else 0.0 for c in context["categories"]]
        subject_hot = [1.0 if s in interests else 0.0 for s in context["subjects"]]
        return [*area_hot, *category_hot, *subject_hot]

    def encode_user_book_pair(self, user: User, book: Book) -> list[float]:
        purchased_ids = set(user.purchased_book_ids)
        bought_this_book = 1.0 if book.id in purchased_ids else 0.0
        return [bought_this_book]

    def create_training_data(self, users: list[User], books: list[Book], context: dict):
        xs: list[list[float]] = []
        ys: list[list[float]] = []

        for user in users:
            user_vector = self.encode_user(user, context)
            interest_set = set(user.interest_areas)
            purchased_ids = set(user.purchased_book_ids)
            for book in books:
                book_vector = self.encode_book(book, context)
                pair_vector = self.encode_user_book_pair(user, book)
                label = 1.0 if (
                    book.area in interest_set
                    or book.category in interest_set
                    or book.subject in interest_set
                    or book.id in purchased_ids
                ) else 0.0

                xs.append([*user_vector, *book_vector, *pair_vector])
                ys.append([label])

        input_dim = len(xs[0]) if xs else 0
        return xs, ys, input_dim
