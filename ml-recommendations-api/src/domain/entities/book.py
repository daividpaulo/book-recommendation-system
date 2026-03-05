from dataclasses import dataclass


@dataclass
class Book:
    id: str
    title: str
    author: str
    category: str
    subject: str
    area: str
    description: str
