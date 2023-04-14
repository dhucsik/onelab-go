package postgre

import (
	"context"
	"practice/models"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      ``
	UpdatedAt time.Time      ``
	DeletedAt gorm.DeletedAt `gorm:"index"`
	FirstName string         ``
	LastName  string         ``
	Username  string         `gorm:"unique"`
	Email     string         `gorm:"unique"`
	Password  string         ``
	UserRole  string         ``
	Records   []Record       ``
	Books     []Book         `gorm:"many2many:records"`
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) (string, error) {
	model := toPostgreUser(user)
	result := r.db.WithContext(ctx).Omit("deleted_at").Create(&model)
	return strconv.FormatUint(uint64(model.ID), 10), result.Error
}

func (r *UserRepository) GetByUsername(ctx context.Context, username string) (models.User, error) {
	user := new(User)
	err := r.db.WithContext(ctx).Where("username = ?", username).First(user).Error
	if err != nil {
		return models.User{}, err
	}
	return toUserModel(user), nil
}

func (r *UserRepository) GetUsersWithBooks(ctx context.Context) ([]models.UserBook, error) {
	var users []User
	err := r.db.Model(&User{}).Preload("Records").Preload("Books").Find(&users).Where("records.date_returned IS NULL").Error

	return toUserBookModels(users), err
}

func (r *UserRepository) UpdatePassword(ctx context.Context, ID string, password string) error {
	return r.db.Model(&User{}).Where("id = ?", ID).Update("password", password).Error
}

func (r *UserRepository) GetUsersWithBooksForMonth(ctx context.Context) ([]models.UserBook, error) {
	var users []User
	err := r.db.Model(&User{}).Preload("Records").Preload("Books").Find(&users).Where("records.date_borrowed BETWEEN DATE_TRUNC('month', NOW() - INTERVAL '1 month') AND DATE_TRUNC('month', NOW())").Error

	return toUserBookModels(users), err
}

func toPostgreUser(u *models.User) User {
	return User{
		Username:  u.Username,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Password:  u.Password,
		UserRole:  u.UserRole,
	}
}

func toUserModel(u *User) models.User {
	return models.User{
		ID:        strconv.FormatUint(uint64(u.ID), 10),
		Username:  u.Username,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Password:  u.Password,
		UserRole:  u.UserRole,
	}
}

func toUserBookModels(users []User) []models.UserBook {
	out := make([]models.UserBook, len(users))

	for i, user := range users {
		out[i] = toUserBookModel(&user)
	}

	return out
}

func toUserBookModel(u *User) models.UserBook {
	return models.UserBook{
		ID:       strconv.FormatUint(uint64(u.ID), 10),
		Username: u.Username,
		Firsname: u.FirstName,
		LastName: u.LastName,
		Email:    u.Email,
		Books:    toBookModels(u.Books),
	}
}
