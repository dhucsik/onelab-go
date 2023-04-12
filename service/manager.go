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
	Update(ctx context.Context, book *models.Book) error
	Delete(ctx context.Context, ID string) error
	Get(ctx context.Context, ID string) (models.Book, error)
	List(ctx context.Context) ([]models.Book, error)
}

type IRecordService interface {
	Create(ctx context.Context, record *models.Record) (string, error)
	Update(ctx context.Context, record *models.Record) error
	Delete(ctx context.Context, ID string) error
	Get(ctx context.Context, ID string) (models.Record, error)
	List(ctx context.Context) ([]models.Record, error)
}

type Manager struct {
	User   IUserService
	Book   IBookService
	Record IRecordService
}

func NewManager(storage *storage.Storage) (*Manager, error) {
	if storage == nil {
		return nil, errors.New("no storage provided")
	}

	uSrv := NewUserService(storage)
	bSrv := NewBookService(storage)
	rSrv := NewRecordService(storage)

	return &Manager{
		User:   uSrv,
		Book:   bSrv,
		Record: rSrv,
	}, nil
}
