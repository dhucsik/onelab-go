package service

import (
	"context"
	"errors"
	"practice/models"
	"practice/storage"
	"time"
)

type RecordService struct {
	repo *storage.Storage
}

func NewRecordService(repo *storage.Storage) *RecordService {
	return &RecordService{
		repo: repo,
	}
}

func (s *RecordService) Create(ctx context.Context, record *models.Record) (string, error) {
	user := ctx.Value(models.ContextKey).(*models.ContextUserData)

	userFromDB, err := s.repo.User.GetByUsername(ctx, user.Usename)
	if err != nil {
		return "", err
	}

	record.UserID = userFromDB.ID
	record.DateBorrowed = time.Now()
	record.DueDate = record.DateBorrowed.AddDate(0, 0, 10)

	return s.repo.Record.Create(ctx, record)
}

func (s *RecordService) Update(ctx context.Context, ID string, record *models.Record) error {
	user := ctx.Value(models.ContextKey).(*models.ContextUserData)

	if user.UserRole != models.AdminRole && user.UserRole != models.ModeratorRole {
		return errors.New("permission denied")
	}

	return s.repo.Record.Update(ctx, ID, record)
}

func (s *RecordService) Delete(ctx context.Context, ID string) error {
	user := ctx.Value(models.ContextKey).(*models.ContextUserData)

	if user.UserRole != models.AdminRole && user.UserRole != models.ModeratorRole {
		return errors.New("permission denied")
	}

	return s.repo.Record.Delete(ctx, ID)
}

func (s *RecordService) Get(ctx context.Context, ID string) (models.Record, error) {
	return s.repo.Record.Get(ctx, ID)
}

func (s *RecordService) List(ctx context.Context) ([]models.Record, error) {
	return s.repo.Record.List(ctx)
}
