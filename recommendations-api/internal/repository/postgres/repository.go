package postgres

import (
	"context"
	"database/sql"

	"book-recommendation-system/recommendations-api/internal/domain/entities"
	"github.com/lib/pq"
)

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateBook(ctx context.Context, book entities.Book) error {
	_, err := r.db.ExecContext(
		ctx,
		`INSERT INTO books (id, title, author, category, subject, area, description)
		 VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		book.ID, book.Title, book.Author, book.Category, book.Subject, book.Area, book.Description,
	)
	return err
}

func (r *Repository) ListBooks(ctx context.Context) ([]entities.Book, error) {
	rows, err := r.db.QueryContext(
		ctx,
		`SELECT id, title, author, category, subject, area, description
		 FROM books ORDER BY created_at DESC`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []entities.Book
	for rows.Next() {
		var book entities.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Category, &book.Subject, &book.Area, &book.Description); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (r *Repository) BookExists(ctx context.Context, bookID string) (bool, error) {
	var exists bool
	err := r.db.QueryRowContext(
		ctx,
		`SELECT EXISTS(SELECT 1 FROM books WHERE id = $1)`,
		bookID,
	).Scan(&exists)
	return exists, err
}

func (r *Repository) CreateUser(ctx context.Context, user entities.User) error {
	_, err := r.db.ExecContext(
		ctx,
		`INSERT INTO users (id, name, age, profession, interest_areas)
		 VALUES ($1, $2, $3, $4, $5)`,
		user.ID, user.Name, user.Age, user.Profession, pq.Array(user.InterestAreas),
	)
	return err
}

func (r *Repository) ListUsers(ctx context.Context) ([]entities.User, error) {
	rows, err := r.db.QueryContext(
		ctx,
		`SELECT id, name, age, profession, interest_areas
		 FROM users ORDER BY created_at DESC`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entities.User
	for rows.Next() {
		var user entities.User
		var interests []string
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Profession, pq.Array(&interests)); err != nil {
			return nil, err
		}
		user.InterestAreas = interests
		users = append(users, user)
	}
	return users, nil
}

func (r *Repository) UserExists(ctx context.Context, userID string) (bool, error) {
	var exists bool
	err := r.db.QueryRowContext(
		ctx,
		`SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)`,
		userID,
	).Scan(&exists)
	return exists, err
}

func (r *Repository) CreatePurchase(ctx context.Context, purchase entities.Purchase) error {
	_, err := r.db.ExecContext(
		ctx,
		`INSERT INTO purchases (id, user_id, book_id, quantity)
		 VALUES ($1, $2, $3, $4)`,
		purchase.ID, purchase.UserID, purchase.BookID, purchase.Quantity,
	)
	return err
}

func (r *Repository) ListPurchases(ctx context.Context) ([]entities.Purchase, error) {
	rows, err := r.db.QueryContext(
		ctx,
		`SELECT id, user_id, book_id, quantity
		 FROM purchases ORDER BY purchased_at DESC`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var purchases []entities.Purchase
	for rows.Next() {
		var purchase entities.Purchase
		if err := rows.Scan(&purchase.ID, &purchase.UserID, &purchase.BookID, &purchase.Quantity); err != nil {
			return nil, err
		}
		purchases = append(purchases, purchase)
	}
	return purchases, nil
}

func (r *Repository) ListPurchasesByUser(ctx context.Context, userID string) ([]entities.Purchase, error) {
	rows, err := r.db.QueryContext(
		ctx,
		`SELECT id, user_id, book_id, quantity
		 FROM purchases
		 WHERE user_id = $1
		 ORDER BY purchased_at DESC`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var purchases []entities.Purchase
	for rows.Next() {
		var purchase entities.Purchase
		if err := rows.Scan(&purchase.ID, &purchase.UserID, &purchase.BookID, &purchase.Quantity); err != nil {
			return nil, err
		}
		purchases = append(purchases, purchase)
	}
	return purchases, nil
}
