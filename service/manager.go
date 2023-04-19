package service

import (
	"context"
	"errors"
	"practice/models"
	"practice/storage"
)

type IUserService interface {
	SignUp(ctx context.Context, user *models.User) (string, error)
	SignIn(ctx context.Context, auth *models.AuthUser) (*models.ContextUserData, error)
	HashPassword(password string) (string, error)
	CheckPassword(hashedPwd, inputPwd string) error
	VerifyRole(role string) error
	UpdatePassword(ctx context.Context, req *models.UpdatePasswordReq) error
	GetUsersWithBooks(ctx context.Context) ([]models.UserBook, error)
	GetUsersWithBooksForMonth(ctx context.Context) ([]models.UserBook, error)
}

type IBookService interface {
	Create(ctx context.Context, book *models.Book) (string, error)
	Update(ctx context.Context, ID string, book *models.Book) error
	Delete(ctx context.Context, ID string) error
	Get(ctx context.Context, ID string) (models.Book, error)
	List(ctx context.Context) ([]models.Book, error)
	GetBooksUsersIncome(ctx context.Context) ([]models.BookUserIncome, error)
}

type IBookRentService interface {
	Create(ctx context.Context, bookRent *models.BookRent) (string, error)
	Update(ctx context.Context, ID string, bookRent *models.BookRent) error
	Delete(ctx context.Context, ID string) error
	Get(ctx context.Context, ID string) (models.BookRent, error)
	List(ctx context.Context) ([]models.BookRent, error)
}

type Manager struct {
	User     IUserService
	Book     IBookService
	BookRent IBookRentService
}

func NewManager(storage *storage.Storage) (*Manager, error) {
	if storage == nil {
		return nil, errors.New("no storage provided")
	}

	uSrv := NewUserService(storage)
	bSrv := NewBookService(storage)
	rSrv := NewBookRentService(storage)

	return &Manager{
		User:     uSrv,
		Book:     bSrv,
		BookRent: rSrv,
	}, nil
}
