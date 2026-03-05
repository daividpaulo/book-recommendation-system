package ports

import (
	"context"

	"book-recommendation-system/recommendations-api/internal/domain/entities"
)

type BookRepository interface {
	CreateBook(ctx context.Context, book entities.Book) error
	ListBooks(ctx context.Context) ([]entities.Book, error)
	BookExists(ctx context.Context, bookID string) (bool, error)
}

type UserRepository interface {
	CreateUser(ctx context.Context, user entities.User) error
	ListUsers(ctx context.Context) ([]entities.User, error)
	UserExists(ctx context.Context, userID string) (bool, error)
}

type PurchaseRepository interface {
	CreatePurchase(ctx context.Context, purchase entities.Purchase) error
	ListPurchases(ctx context.Context) ([]entities.Purchase, error)
	ListPurchasesByUser(ctx context.Context, userID string) ([]entities.Purchase, error)
}

type MLGateway interface {
	EmbedBook(book entities.Book)
	EmbedUser(user entities.User)
	Train(payload map[string]any) (map[string]any, error)
	Recommend(userID string) (map[string]any, error)
}
