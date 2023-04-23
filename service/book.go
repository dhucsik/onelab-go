package service

import (
	"context"
	"errors"
	"practice/logging"
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

func (s *BookService) Create(ctx context.Context, book *models.Book, logger *logging.Logger) (string, error) {
	logger.Infof("Service level\n")

	user := ctx.Value(models.ContextKey).(*models.ContextUserData)

	logger.Infof("Get user %v\n", user)

	if user.UserRole != models.AdminRole && user.UserRole != models.ModeratorRole {
		return "", errors.New("permission denied")
	}

	return s.repo.Book.Create(ctx, book, logger)
}

func (s *BookService) Update(ctx context.Context, ID string, book *models.Book, logger *logging.Logger) error {
	logger.Infof("Service level\n")

	user := ctx.Value(models.ContextKey).(*models.ContextUserData)

	logger.Infof("Get user %v\n", user)

	if user.UserRole != models.AdminRole && user.UserRole != models.ModeratorRole {
		return errors.New("permission denied")
	}

	return s.repo.Book.Update(ctx, ID, book, logger)
}

func (s *BookService) Delete(ctx context.Context, ID string, logger *logging.Logger) error {
	logger.Infof("Service level\n")

	user := ctx.Value(models.ContextKey).(*models.ContextUserData)

	logger.Infof("Get user %v\n", user)

	if user.UserRole != models.AdminRole && user.UserRole != models.ModeratorRole {
		return errors.New("permission denied")
	}

	return s.repo.Book.Delete(ctx, ID, logger)
}

func (s *BookService) Get(ctx context.Context, ID string, logger *logging.Logger) (models.Book, error) {
	logger.Infof("Service level\n")

	return s.repo.Book.Get(ctx, ID, logger)
}

func (s *BookService) List(ctx context.Context, logger *logging.Logger) ([]models.Book, error) {
	logger.Infof("Service level\n")

	return s.repo.Book.List(ctx, logger)
}

func (s *BookService) GetBooksUsersIncome(ctx context.Context, logger *logging.Logger) ([]models.BookUserIncome, error) {
	logger.Infof("Service level\n")

	return s.repo.Book.GetBooksUsersIncome(ctx, logger)
}
