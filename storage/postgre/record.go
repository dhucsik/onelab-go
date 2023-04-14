package postgre

import (
	"context"
	"practice/models"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type Record struct {
	ID           uint           `gorm:"primaryKey"`
	CreatedAt    time.Time      ``
	UpdatedAt    time.Time      ``
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	UserID       uint
	BookID       uint
	DateBorrowed time.Time
	DateReturned *time.Time
	DueDate      time.Time
}

type RecordRepository struct {
	db *gorm.DB
}

func NewRecordRepository(db *gorm.DB) *RecordRepository {
	return &RecordRepository{db: db}
}

func (r *RecordRepository) Create(ctx context.Context, record *models.Record) (string, error) {
	model, err := toPostgreRecord(record)
	if err != nil {
		return "", err
	}

	result := r.db.WithContext(ctx).Omit("deleted_at", "date_returned").Create(&model)
	return strconv.FormatUint(uint64(model.ID), 10), result.Error
}

func (r *RecordRepository) Update(ctx context.Context, ID string, record *models.Record) error {
	id, err := strconv.ParseUint(ID, 10, 32)
	if err != nil {
		return err
	}

	model, err := toPostgreRecord(record)
	if err != nil {
		return err
	}

	model.ID = uint(id)
	return r.db.Save(&model).Error
}

func (r *RecordRepository) Get(ctx context.Context, ID string) (models.Record, error) {
	record := new(Record)
	err := r.db.WithContext(ctx).Where("id = ?", ID).First(record).Error
	if err != nil {
		return models.Record{}, err
	}

	return toRecordModel(record), nil
}

func (r *RecordRepository) Delete(ctx context.Context, ID string) error {
	return r.db.WithContext(ctx).Delete(&Record{}, ID).Error
}

func (r *RecordRepository) List(ctx context.Context) ([]models.Record, error) {
	var records []Record
	err := r.db.WithContext(ctx).Find(&records)
	return toRecordModels(records), err.Error
}

func toPostgreRecord(r *models.Record) (Record, error) {
	userID, err := strconv.ParseUint(r.UserID, 10, 32)
	if err != nil {
		return Record{}, err
	}

	bookID, err := strconv.ParseUint(r.BookID, 10, 32)
	if err != nil {
		return Record{}, err
	}

	return Record{
		UserID:       uint(userID),
		BookID:       uint(bookID),
		DateBorrowed: r.DateBorrowed,
		DateReturned: r.DateReturned,
		DueDate:      r.DueDate,
	}, nil
}

func toRecordModel(r *Record) models.Record {
	return models.Record{
		ID:           strconv.FormatUint(uint64(r.ID), 10),
		UserID:       strconv.FormatUint(uint64(r.UserID), 10),
		BookID:       strconv.FormatUint(uint64(r.BookID), 10),
		DateBorrowed: r.DateBorrowed,
		DateReturned: r.DateReturned,
		DueDate:      r.DueDate,
	}
}

func toRecordModels(records []Record) []models.Record {
	out := make([]models.Record, len(records))

	for i, record := range records {
		out[i] = toRecordModel(&record)
	}

	return out
}
