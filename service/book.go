package service

import (
	"context"
	"errors"
	"practice/models"
	"practice/storage"
)

type BookService struct {
	repo *storage.Storage
}

func NewBookService(repo *storage.Storage) *BookService {
	return &BookService{
		repo: repo,
	}
}

func (s *BookService) Create(ctx context.Context, book *models.Book) (string, error) {
	user := ctx.Value(models.ContextKey).(*models.ContextUserData)

	if user.UserRole != models.AdminRole && user.UserRole != models.ModeratorRole {
		return "", errors.New("permission denied")
	}

	return s.repo.Book.Create(ctx, book)
}

func (s *BookService) Update(ctx context.Context, book *models.Book) error {
	user := ctx.Value(models.ContextKey).(*models.ContextUserData)

	if user.UserRole != models.AdminRole && user.UserRole != models.ModeratorRole {
		return errors.New("permission denied")
	}

	return s.repo.Book.Update(ctx, book)
}

func (s *BookService) Delete(ctx context.Context, ID string) error {
	user := ctx.Value(models.ContextKey).(*models.ContextUserData)

	if user.UserRole != models.AdminRole && user.UserRole != models.ModeratorRole {
		return errors.New("permission denied")
	}

	return s.repo.Book.Delete(ctx, ID)
}

func (s *BookService) Get(ctx context.Context, ID string) (models.Book, error) {
	return s.repo.Book.Get(ctx, ID)
}

func (s *BookService) List(ctx context.Context) ([]models.Book, error) {
	return s.repo.Book.List(ctx)
}
