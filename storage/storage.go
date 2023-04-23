package storage

import (
	"context"
	"practice/config"
	"practice/logging"
	"practice/models"
	"practice/storage/postgre"
)

type IUserRepository interface {
	Create(ctx context.Context, user *models.User) (string, error)
	GetByUsername(ctx context.Context, username string) (models.User, error)
	UpdatePassword(ctx context.Context, ID string, password string) error
	GetUsersWithBooks(ctx context.Context) ([]models.UserBook, error)
	GetUsersWithBooksForMonth(ctx context.Context) ([]models.UserBook, error)
}

type IBookRepository interface {
	Create(ctx context.Context, book *models.Book, logger *logging.Logger) (string, error)
	Update(ctx context.Context, ID string, book *models.Book, logger *logging.Logger) error
	Get(ctx context.Context, ID string, logger *logging.Logger) (models.Book, error)
	Delete(ctx context.Context, ID string, logger *logging.Logger) error
	List(ctx context.Context, logger *logging.Logger) ([]models.Book, error)
	GetBooksUsersIncome(ctx context.Context, logger *logging.Logger) ([]models.BookUserIncome, error)
}

type IBookRentRepository interface {
	Create(ctx context.Context, bookRent *models.BookRent) (string, error)
	Get(ctx context.Context, ID string) (models.BookRent, error)
	Update(ctx context.Context, ID string, bookRent *models.BookRent) error
	List(ctx context.Context) ([]models.BookRent, error)
	Delete(ctx context.Context, ID string) error
}

type Storage struct {
	User     IUserRepository
	Book     IBookRepository
	BookRent IBookRentRepository
}

func New(ctx context.Context, cfg *config.Config) (*Storage, error) {
	pgDB, err := postgre.Dial(ctx, cfg.PgURL)
	if err != nil {
		return nil, err
	}

	userRepo := postgre.NewUserRepository(pgDB)
	bookRepo := postgre.NewBookRepository(pgDB)
	rentRepo := postgre.NewBookRentRepository(pgDB)

	storage := Storage{
		User:     userRepo,
		Book:     bookRepo,
		BookRent: rentRepo,
	}
	return &storage, nil
}
