from dataclasses import dataclass


@dataclass
class Recommendation:
    book_id: str
    score: float
