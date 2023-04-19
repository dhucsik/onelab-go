package handler

import (
	"net/http"
	"practice/models"

	"github.com/labstack/echo/v4"
)

// @Summary Create Book Rent
// @Security ApiKeyAuth
// @Tags rent
// @Description create book rent
// @ID create-rent
// @Accept  json
// @Produce  json
// @Param input body models.BookRent true "rent info"
// @Success 201 {string} string "id"
// @Failure 400,404 {string} string "message"
// @Failure 500 {string} string "message"
// @Failure default {string} string "message"
// @Router /rent [post]
func (h Manager) CreateBookRent(c echo.Context) error {
	var req models.BookRent
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	resp, err := h.srv.BookRent.Create(c.Request().Context(), &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, resp)
}

// @Summary Update Book Rent
// @Security ApiKeyAuth
// @Tags rent
// @Description update book rent
// @ID update-rent
// @Accept json
// @Produce json
// @Param rent_id path string true "rent id"
// @Param input body models.BookRent true "rent info"
// @Success 200
// @Failure 400,404 {string} string "message"
// @Failure 500 {string} string "message"
// @Failure default {string} string "message"
// @Router /rent/{rent_id} [put]
func (h Manager) UpdateBookRent(c echo.Context) error {
	id := c.Param("id")

	var req models.BookRent
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := h.srv.BookRent.Update(c.Request().Context(), id, &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

// @Summary List Book Rents
// @Tags rent
// @Description list book rents
// @ID list-rents
// @Accept json
// @Produce json
// @Success 200 {object} []models.BookRent
// @Failure 400,404 {object} string "message"
// @Failure 500 {object} string "message"
// @Failure default {object} string "message"
// @Router /rent [get]
func (h Manager) ListBookRents(c echo.Context) error {
	resp, err := h.srv.BookRent.List(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

// @Summary Get Book Rent
// @Tags rent
// @Description get book rent by ID
// @ID get-rent
// @Accept json
// @Produce json
// @Param rent_id path string true "rent id"
// @Success 200 {object} models.BookRent
// @Failure 400,404 {object} string "message"
// @Failure 500 {object} string "message"
// @Failure default {object} string "message"
// @Router /rent/{rent_id} [get]
func (h Manager) GetBookRent(c echo.Context) error {
	id := c.Param("id")

	resp, err := h.srv.BookRent.Get(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

// @Summary Delete Book Rent
// @Security ApiKeyAuth
// @Tags rent
// @Description delete book rent by ID
// @ID delete-rent
// @Accept json
// @Produce json
// @Param rent_id path string true "rent id"
// @Success 200
// @Failure 400,404 {object} string "message"
// @Failure 500 {object} string "message"
// @Failure default {object} string "message"
// @Router /rent/{rent_id} [delete]
func (h Manager) DeleteBookRent(c echo.Context) error {
	id := c.Param("id")

	err := h.srv.BookRent.Delete(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
