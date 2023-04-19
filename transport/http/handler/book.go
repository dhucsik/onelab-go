package handler

import (
	"net/http"
	"practice/models"

	"github.com/labstack/echo/v4"
)

// @Summary Create Book
// @Security ApiKeyAuth
// @Tags book
// @Description create book
// @ID create-book
// @Accept  json
// @Produce  json
// @Param input body models.Book true "book info"
// @Success 201 {string} string "id"
// @Failure 400,404 {string} string "message"
// @Failure 500 {string} string "message"
// @Failure default {string} string "message"
// @Router /book [post]
func (h Manager) CreateBook(c echo.Context) error {
	var req models.Book
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ID, err := h.srv.Book.Create(c.Request().Context(), &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"id": ID,
	})
}

// @Summary Update Book
// @Security ApiKeyAuth
// @Tags book
// @Description update book
// @ID update-book
// @Accept json
// @Produce json
// @Param book_id path string true "book id"
// @Param input body models.Book true "book info"
// @Success 200
// @Failure 400,404 {string} string "message"
// @Failure 500 {string} string "message"
// @Failure default {string} string "message"
// @Router /book/{book_id} [put]
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

// @Summary List Books
// @Tags book
// @Description list books
// @ID list-books
// @Accept json
// @Produce json
// @Success 200 {object} []models.Book
// @Failure 400,404 {object} string "message"
// @Failure 500 {object} string "message"
// @Failure default {object} string "message"
// @Router /book [get]
func (h Manager) ListBooks(c echo.Context) error {
	resp, err := h.srv.Book.List(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

// @Summary Get Book
// @Tags book
// @Description get book by ID
// @ID get-book
// @Accept json
// @Produce json
// @Param book_id path string true "book id"
// @Success 200 {object} models.Book
// @Failure 400,404 {object} string "message"
// @Failure 500 {object} string "message"
// @Failure default {object} string "message"
// @Router /book/{book_id} [get]
func (h Manager) GetBook(c echo.Context) error {
	id := c.Param("id")

	resp, err := h.srv.Book.Get(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

// @Summary Delete Book
// @Security ApiKeyAuth
// @Tags book
// @Description delete book by ID
// @ID delete-book
// @Accept json
// @Produce json
// @Param book_id path string true "book id"
// @Success 200
// @Failure 400,404 {object} string "message"
// @Failure 500 {object} string "message"
// @Failure default {object} string "message"
// @Router /book/{book_id} [delete]
func (h Manager) DeleteBook(c echo.Context) error {
	id := c.Param("id")

	err := h.srv.Book.Delete(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

// @Summary GetBooksUsersIncome
// @Tags book
// @Description get books that users have and total income for each
// @ID get-book-users-income
// @Accept json
// @Produce json
// @Success 200 {object} []models.BookUserIncome
// @Failure 400,404 {string} string "message"
// @Failure 500 {string} string "message"
// @Failure default {string} string "message"
// @Router /book-users-income [get]
func (h Manager) GetBooksUsersIncome(c echo.Context) error {
	resp, err := h.srv.Book.GetBooksUsersIncome(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}
