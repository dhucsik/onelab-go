package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"practice/config"
	"practice/service"
	"practice/storage"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var (
	getBookJSON = `{
		"id": "2",
		"title": "book2",
		"author": "author2",
		"rent_price": 0
	}`
)

func TestBookHandler_Get(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/book", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h, err := getBookHandler(t)
	if err != nil {
		t.Errorf("couldn't run server")
		return
	}

	if assert.NoError(t, h.GetBook(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, getBookJSON, rec.Body.String())
	}
}

func getBookHandler(t *testing.T) (*Manager, error) {
	viper.AddConfigPath("../..")
	cfg, err := config.New()
	if err != nil {
		return nil, err
	}

	repo, err := storage.New(context.Background(), cfg)
	if err != nil {
		return nil, err
	}

	srvManager, err := service.NewManager(repo)
	if err != nil {
		return nil, err
	}

	return NewManager(srvManager), nil
}
