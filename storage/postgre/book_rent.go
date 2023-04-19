package postgre

import (
	"context"
	"practice/models"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type BookRent struct {
	ID           uint           `gorm:"primaryKey"`
	CreatedAt    time.Time      ``
	UpdatedAt    time.Time      ``
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	UserID       uint
	BookID       uint
	DateRented   time.Time
	DateReturned *time.Time
	DueDate      time.Time
	RentPrice    float64
}

type BookRentRepository struct {
	db *gorm.DB
}

func NewBookRentRepository(db *gorm.DB) *BookRentRepository {
	return &BookRentRepository{db: db}
}

func (r *BookRentRepository) Create(ctx context.Context, bookRent *models.BookRent) (string, error) {
	model, err := toPostgreBookRent(bookRent)
	if err != nil {
		return "", err
	}

	result := r.db.WithContext(ctx).Omit("deleted_at", "date_returned").Create(&model)
	return strconv.FormatUint(uint64(model.ID), 10), result.Error
}

func (r *BookRentRepository) Update(ctx context.Context, ID string, bookRent *models.BookRent) error {
	id, err := strconv.ParseUint(ID, 10, 32)
	if err != nil {
		return err
	}

	model, err := toPostgreBookRent(bookRent)
	if err != nil {
		return err
	}

	model.ID = uint(id)
	return r.db.Save(&model).Error
}

func (r *BookRentRepository) Get(ctx context.Context, ID string) (models.BookRent, error) {
	bookRent := new(BookRent)
	err := r.db.WithContext(ctx).Where("id = ?", ID).First(bookRent).Error
	if err != nil {
		return models.BookRent{}, err
	}

	return toBookRentModel(bookRent), nil
}

func (r *BookRentRepository) Delete(ctx context.Context, ID string) error {
	return r.db.WithContext(ctx).Delete(&BookRent{}, ID).Error
}

func (r *BookRentRepository) List(ctx context.Context) ([]models.BookRent, error) {
	var bookRents []BookRent
	err := r.db.WithContext(ctx).Find(&bookRents)
	return toBookRentModels(bookRents), err.Error
}

func toPostgreBookRent(r *models.BookRent) (BookRent, error) {
	userID, err := strconv.ParseUint(r.UserID, 10, 32)
	if err != nil {
		return BookRent{}, err
	}

	bookID, err := strconv.ParseUint(r.BookID, 10, 32)
	if err != nil {
		return BookRent{}, err
	}

	return BookRent{
		UserID:       uint(userID),
		BookID:       uint(bookID),
		DateRented:   r.DateRented,
		DateReturned: r.DateReturned,
		DueDate:      r.DueDate,
	}, nil
}

func toBookRentModel(r *BookRent) models.BookRent {
	return models.BookRent{
		ID:           strconv.FormatUint(uint64(r.ID), 10),
		UserID:       strconv.FormatUint(uint64(r.UserID), 10),
		BookID:       strconv.FormatUint(uint64(r.BookID), 10),
		DateRented:   r.DateRented,
		DateReturned: r.DateReturned,
		DueDate:      r.DueDate,
	}
}

func toBookRentModels(bookRents []BookRent) []models.BookRent {
	out := make([]models.BookRent, len(bookRents))

	for i, bookRent := range bookRents {
		out[i] = toBookRentModel(&bookRent)
	}

	return out
}
