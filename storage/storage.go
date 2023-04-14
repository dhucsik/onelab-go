package storage

import (
	"context"
	"practice/config"
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
	Create(ctx context.Context, book *models.Book) (string, error)
	Update(ctx context.Context, ID string, book *models.Book) error
	Get(ctx context.Context, ID string) (models.Book, error)
	Delete(ctx context.Context, ID string) error
	List(ctx context.Context) ([]models.Book, error)
}

type IRecordRepository interface {
	Create(ctx context.Context, record *models.Record) (string, error)
	Get(ctx context.Context, ID string) (models.Record, error)
	Update(ctx context.Context, ID string, record *models.Record) error
	List(ctx context.Context) ([]models.Record, error)
	Delete(ctx context.Context, ID string) error
}

type Storage struct {
	User   IUserRepository
	Book   IBookRepository
	Record IRecordRepository
}

func New(ctx context.Context, cfg *config.Config) (*Storage, error) {
	pgDB, err := postgre.Dial(ctx, cfg.PgURL)
	if err != nil {
		return nil, err
	}

	userRepo := postgre.NewUserRepository(pgDB)
	bookRepo := postgre.NewBookRepository(pgDB)
	recordRepo := postgre.NewRecordRepository(pgDB)

	storage := Storage{
		User:   userRepo,
		Book:   bookRepo,
		Record: recordRepo,
	}
	return &storage, nil
}
