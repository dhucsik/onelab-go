package handler

import (
	"net/http"
	"practice/models"

	"github.com/labstack/echo"
)

func (h Manager) CreateBook(c echo.Context) error {
	var req models.Book
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	resp, err := h.srv.Book.Create(c.Request().Context(), &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, resp)
}

func (h Manager) UpdateBook(c echo.Context) error {
	id := c.Param("id")

	var req models.Book
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := h.srv.Book.Update(c.Request().Context(), id, &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (h Manager) ListBooks(c echo.Context) error {
	resp, err := h.srv.Book.List(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

func (h Manager) GetBook(c echo.Context) error {
	id := c.Param("id")

	resp, err := h.srv.Book.Get(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

func (h Manager) DeleteBook(c echo.Context) error {
	id := c.Param("id")

	err := h.srv.Book.Delete(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
