from dataclasses import dataclass
from dataclasses import field


@dataclass
class User:
    id: str
    name: str
    age: int
    profession: str
    interest_areas: list[str]
    purchase_count: int = 0
    purchased_book_ids: list[str] = field(default_factory=list)
