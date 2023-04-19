package service

import (
	"context"
	"errors"
	"practice/models"
	"practice/storage"
	"time"
)

type BookRentService struct {
	repo *storage.Storage
}

func NewBookRentService(repo *storage.Storage) *BookRentService {
	return &BookRentService{
		repo: repo,
	}
}

func (s *BookRentService) Create(ctx context.Context, bookRent *models.BookRent) (string, error) {
	user := ctx.Value(models.ContextKey).(*models.ContextUserData)

	userFromDB, err := s.repo.User.GetByUsername(ctx, user.Usename)
	if err != nil {
		return "", err
	}

	bookRent.UserID = userFromDB.ID
	bookRent.DateRented = time.Now()
	bookRent.DueDate = bookRent.DateRented.AddDate(0, 0, 10)

	return s.repo.BookRent.Create(ctx, bookRent)
}

func (s *BookRentService) Update(ctx context.Context, ID string, bookRent *models.BookRent) error {
	user := ctx.Value(models.ContextKey).(*models.ContextUserData)

	if user.UserRole != models.AdminRole && user.UserRole != models.ModeratorRole {
		return errors.New("permission denied")
	}

	return s.repo.BookRent.Update(ctx, ID, bookRent)
}

func (s *BookRentService) Delete(ctx context.Context, ID string) error {
	user := ctx.Value(models.ContextKey).(*models.ContextUserData)

	if user.UserRole != models.AdminRole && user.UserRole != models.ModeratorRole {
		return errors.New("permission denied")
	}

	return s.repo.BookRent.Delete(ctx, ID)
}

func (s *BookRentService) Get(ctx context.Context, ID string) (models.BookRent, error) {
	return s.repo.BookRent.Get(ctx, ID)
}

func (s *BookRentService) List(ctx context.Context) ([]models.BookRent, error) {
	return s.repo.BookRent.List(ctx)
}
