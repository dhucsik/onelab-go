package postgre

import (
	"context"
	"log"
	"practice/config"
	"practice/logging"
	"practice/models"
	"testing"

	"github.com/spf13/viper"
)

func TestBookRepo_Create(t *testing.T) {
	logger := logging.GetLogger()

	type args struct {
		ctx  context.Context
		book *models.Book
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Success",
			args:    args{context.Background(), &models.Book{Title: "NewBook", Author: "NewAuthor", RentPrice: 19.19}},
			want:    "6",
			wantErr: false,
		},
	}
	repo, err := CreateBookRepo(t)
	if err != nil {
		t.Errorf("couldn't connect to database")
		return
	}

	for _, tCase := range tests {
		t.Run(tCase.name, func(t *testing.T) {
			resp, err := repo.Create(tCase.args.ctx, tCase.args.book, logger)
			if (err != nil) != tCase.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tCase.wantErr)
				return
			}
			if resp != tCase.want {
				t.Errorf("Create() resp = %v, want %v", resp, tCase.want)
			}

			book, err := repo.Get(context.Background(), resp, logger)
			if err != nil {
				t.Errorf("Create() resp = %v, want %v", resp, tCase.want)
			}

			log.Print(book)
		})
	}
}

func CreateBookRepo(t *testing.T) (*BookRepository, error) {
	viper.AddConfigPath("../..")
	cfg, err := config.New()
	if err != nil {
		return nil, err
	}

	db, err := Dial(context.Background(), cfg.PgURL)
	if err != nil {
		return nil, err
	}
	return NewBookRepository(db), nil
}
