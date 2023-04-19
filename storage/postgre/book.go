package postgre

import (
	"context"
	"practice/models"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      ``
	UpdatedAt time.Time      ``
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Title     string         ``
	Author    string         ``
	RentPrice float64        ``
	BookRents []BookRent     ``
	Users     []User         `gorm:"many2many:book_rents"`
}

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) Create(ctx context.Context, book *models.Book) (string, error) {
	model := toPostgreBook(book)
	result := r.db.WithContext(ctx).Omit("deleted_at").Create(&model)
	return strconv.FormatUint(uint64(model.ID), 10), result.Error
}

func (r *BookRepository) Update(ctx context.Context, ID string, book *models.Book) error {
	id, err := strconv.ParseUint(ID, 10, 32)
	if err != nil {
		return err
	}

	model := toPostgreBook(book)
	model.ID = uint(id)
	return r.db.Save(&model).Error
}

func (r *BookRepository) Get(ctx context.Context, ID string) (models.Book, error) {
	book := new(Book)
	err := r.db.WithContext(ctx).Where("id = ?", ID).First(book).Error
	if err != nil {
		return models.Book{}, err
	}

	return toBookModel(book), nil
}

func (r *BookRepository) Delete(ctx context.Context, ID string) error {
	return r.db.WithContext(ctx).Delete(&Book{}, ID).Error
}

func (r *BookRepository) List(ctx context.Context) ([]models.Book, error) {
	var books []Book
	err := r.db.WithContext(ctx).Find(&books)
	return toBookModels(books), err.Error
}

func (r *BookRepository) GetBooksUsersIncome(ctx context.Context) ([]models.BookUserIncome, error) {
	var books []Book
	err := r.db.Model(&Book{}).Preload("BookRents").Preload("Users").Find(&books).Error

	return toBookUserIncomeModels(books), err
}

func toPostgreBook(b *models.Book) Book {
	return Book{
		Title:     b.Title,
		Author:    b.Author,
		RentPrice: b.RentPrice,
	}
}

func toBookModel(b *Book) models.Book {
	return models.Book{
		ID:        strconv.FormatUint(uint64(b.ID), 10),
		Title:     b.Title,
		Author:    b.Author,
		RentPrice: b.RentPrice,
	}
}

func toBookModels(books []Book) []models.Book {
	out := make([]models.Book, len(books))

	for i, book := range books {
		out[i] = toBookModel(&book)
	}

	return out
}

func toBookUserIncomeModels(books []Book) []models.BookUserIncome {
	out := make([]models.BookUserIncome, len(books))

	for i, book := range books {
		out[i] = toBookUserIncomeModel(&book)
	}

	return out
}

func toBookUserIncomeModel(book *Book) models.BookUserIncome {
	return models.BookUserIncome{
		ID:          strconv.FormatUint(uint64(book.ID), 10),
		Title:       book.Title,
		Author:      book.Author,
		Users:       toUserRespModels(book.Users),
		TotalIncome: float64(len(book.Users)) * book.RentPrice,
	}
}
