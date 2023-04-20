package service

import (
	"context"
	"practice/models"
	"practice/storage"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *storage.Storage
}

func NewUserService(repo *storage.Storage) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) SignUp(ctx context.Context, user *models.User) (string, error) {
	if err := s.VerifyRole(user.UserRole); err != nil {
		return "", err
	}

	hash, err := s.HashPassword(user.Password)
	if err != nil {
		return "", err
	}

	user.Password = hash

	return s.repo.User.Create(ctx, user)
}

func (s *UserService) VerifyRole(role string) error {
	switch role {
	case models.AdminRole:
		// Nothing to do, verified successfully
	case models.ModeratorRole:
		// Nothing to do, verified successfully
	case models.UserRole:
		// Nothing to do, verified successfully
	default:
		return ErrRoleDoesNotExist
	}

	return nil
}

func (s *UserService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (s *UserService) SignIn(ctx context.Context, user *models.AuthUser) (*models.ContextUserData, error) {
	userFromDB, userErr := s.repo.User.GetByUsername(ctx, user.Username)
	if userErr != nil {
		return nil, userErr
	}

	checkErr := s.CheckPassword(userFromDB.Password, user.Password)
	if checkErr != nil {
		return nil, checkErr
	}

	return &models.ContextUserData{
		Usename:  userFromDB.Username,
		UserRole: userFromDB.UserRole,
	}, nil
}

func (s *UserService) CheckPassword(hashedPwd, inputPwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(inputPwd))
}

func (s *UserService) UpdatePassword(ctx context.Context, req *models.UpdatePasswordReq) error {
	user := ctx.Value(models.ContextKey).(*models.ContextUserData)
	userFromDB, err := s.repo.User.GetByUsername(ctx, user.Usename)
	if err != nil {
		return err
	}

	if err := s.CheckPassword(userFromDB.Password, req.CurrentPassword); err != nil {
		return err
	}

	hash, err := s.HashPassword(req.NewPassword)
	if err != nil {
		return err
	}

	return s.repo.User.UpdatePassword(ctx, userFromDB.ID, hash)
}

func (s *UserService) GetUsersWithBooks(ctx context.Context) ([]models.UserBook, error) {
	return s.repo.User.GetUsersWithBooks(ctx)
}

func (s *UserService) GetUsersWithBooksForMonth(ctx context.Context) ([]models.UserBook, error) {
	return s.repo.User.GetUsersWithBooksForMonth(ctx)
}
