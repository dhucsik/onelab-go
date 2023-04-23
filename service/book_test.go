package service

import (
	"context"
	"errors"
	"practice/logging"
	"practice/models"
	"practice/storage"
	repoMock "practice/storage/mocks"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGet(t *testing.T) {
	logger := logging.GetLogger()

	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		prepare func(f *repoMock.MockIBookRepository)
		name    string
		args    args
		want    models.Book
		wantErr bool
	}{
		{
			name:    "success",
			args:    args{context.Background(), "12"},
			want:    models.Book{ID: "12", Title: "Book12", Author: "Author12"},
			wantErr: false,
			prepare: func(f *repoMock.MockIBookRepository) {
				f.EXPECT().Get(gomock.Any(), "12", logger).Return(models.Book{
					ID: "12", Title: "Book12", Author: "Author12",
				}, nil)
			},
		},
		{
			name:    "fail",
			args:    args{context.Background(), "0"},
			want:    models.Book{},
			wantErr: true,
			prepare: func(f *repoMock.MockIBookRepository) {
				f.EXPECT().Get(gomock.Any(), "0", logger).Return(models.Book{}, errors.New("record not found"))
			},
		},
	}

	for _, tCase := range tests {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		bookRepo := repoMock.NewMockIBookRepository(ctrl)
		tCase.prepare(bookRepo)
		s := NewBookService(&storage.Storage{Book: bookRepo})
		t.Run(tCase.name, func(t *testing.T) {
			resp, err := s.Get(tCase.args.ctx, tCase.args.id, logger)
			if (err != nil) != tCase.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tCase.wantErr)
				return
			}
			if resp != tCase.want {
				t.Errorf("Get() resp = %v, want %v", resp, tCase.want)
			}
		})
	}
}
