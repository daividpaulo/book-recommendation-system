package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"book-recommendation-system/recommendations-api/internal/domain/entities"
	"book-recommendation-system/recommendations-api/internal/domain/ports"
)

type Service struct {
	bookRepo     ports.BookRepository
	userRepo     ports.UserRepository
	purchaseRepo ports.PurchaseRepository
	ml           ports.MLGateway
}

func New(bookRepo ports.BookRepository, userRepo ports.UserRepository, purchaseRepo ports.PurchaseRepository, ml ports.MLGateway) *Service {
	return &Service{
		bookRepo:     bookRepo,
		userRepo:     userRepo,
		purchaseRepo: purchaseRepo,
		ml:           ml,
	}
}

func (s *Service) CreateBook(ctx context.Context, book entities.Book) (entities.Book, error) {
	if book.ID == "" {
		book.ID = fmt.Sprintf("book-%d", time.Now().UnixNano())
	}
	if err := s.bookRepo.CreateBook(ctx, book); err != nil {
		return entities.Book{}, err
	}
	go s.ml.EmbedBook(book)
	return book, nil
}

func (s *Service) ListBooks(ctx context.Context) ([]entities.Book, error) {
	return s.bookRepo.ListBooks(ctx)
}

func (s *Service) CreateUser(ctx context.Context, user entities.User) (entities.User, error) {
	if user.ID == "" {
		user.ID = fmt.Sprintf("user-%d", time.Now().UnixNano())
	}
	if err := s.userRepo.CreateUser(ctx, user); err != nil {
		return entities.User{}, err
	}
	go s.ml.EmbedUser(user)
	return user, nil
}

func (s *Service) ListUsers(ctx context.Context) ([]entities.User, error) {
	return s.userRepo.ListUsers(ctx)
}

func (s *Service) TriggerTraining(ctx context.Context) (map[string]any, error) {
	users, err := s.userRepo.ListUsers(ctx)
	if err != nil {
		return nil, err
	}
	purchases, err := s.purchaseRepo.ListPurchases(ctx)
	if err != nil {
		return nil, err
	}
	books, err := s.bookRepo.ListBooks(ctx)
	if err != nil {
		return nil, err
	}

	booksByUser := map[string][]string{}
	purchaseCountByUser := map[string]int{}
	for _, purchase := range purchases {
		booksByUser[purchase.UserID] = append(booksByUser[purchase.UserID], purchase.BookID)
		purchaseCountByUser[purchase.UserID] += max(1, purchase.Quantity)
	}
	for i := range users {
		users[i].PurchasedBookIDs = booksByUser[users[i].ID]
		users[i].PurchaseCount = purchaseCountByUser[users[i].ID]
	}

	resp, err := s.ml.Train(map[string]any{
		"users": users,
		"books": books,
	})
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrDownstreamFailure, err)
	}
	return resp, nil
}

func (s *Service) GetRecommendations(ctx context.Context, userID string) (map[string]any, error) {
	exists, err := s.userRepo.UserExists(ctx, userID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, ErrUserNotFound
	}
	resp, err := s.ml.Recommend(userID)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrDownstreamFailure, err)
	}
	return resp, nil
}

func (s *Service) CreatePurchase(ctx context.Context, purchase entities.Purchase) (entities.Purchase, error) {
	if purchase.ID == "" {
		purchase.ID = fmt.Sprintf("purchase-%d", time.Now().UnixNano())
	}
	if purchase.Quantity <= 0 {
		purchase.Quantity = 1
	}

	userExists, err := s.userRepo.UserExists(ctx, purchase.UserID)
	if err != nil {
		return entities.Purchase{}, err
	}
	if !userExists {
		return entities.Purchase{}, ErrUserNotFound
	}

	bookExists, err := s.bookRepo.BookExists(ctx, purchase.BookID)
	if err != nil {
		return entities.Purchase{}, err
	}
	if !bookExists {
		return entities.Purchase{}, ErrBookNotFound
	}

	if err := s.purchaseRepo.CreatePurchase(ctx, purchase); err != nil {
		return entities.Purchase{}, err
	}
	return purchase, nil
}

func (s *Service) ListPurchasesByUser(ctx context.Context, userID string) ([]entities.Purchase, error) {
	userExists, err := s.userRepo.UserExists(ctx, userID)
	if err != nil {
		return nil, err
	}
	if !userExists {
		return nil, ErrUserNotFound
	}
	return s.purchaseRepo.ListPurchasesByUser(ctx, userID)
}

var ErrUserNotFound = fmt.Errorf("user not found")
var ErrBookNotFound = fmt.Errorf("book not found")
var ErrDownstreamFailure = errors.New("downstream dependency failure")
