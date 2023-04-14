package handler

import (
	"net/http"
	"practice/models"

	"github.com/labstack/echo/v4"
)

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body models.User true "account info"
// @Success 200 {string} string "id"
// @Failure 400,404 {string} string "message"
// @Failure 500 {string} string "message"
// @Failure default {string} string "message"
// @Router /sign-up [post]
func (h Manager) SignUp(c echo.Context) error {
	var req models.User
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ID, err := h.srv.User.SignUp(c.Request().Context(), &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"id": ID,
	})
}

// @Summary SignIn
// @Tags auth
// @Description login
// @ID login
// @Accept  json
// @Produce  json
// @Param input body models.AuthUser true "credentials"
// @Success 200 {string} string "token"
// @Failure 400,404 {string} string "message"
// @Failure 500 {string} string "message"
// @Failure default {string} string "message"
// @Router /sign-in [post]
func (h Manager) SignIn(c echo.Context) error {
	var req models.AuthUser
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userData, err := h.srv.User.SignIn(c.Request().Context(), &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Set(string(models.ContextKey), userData)

	return nil
}

// @Summary UpdatePassword
// @Security ApiKeyAuth
// @Tags auth
// @Description update password
// @ID update-password
// @Accept json
// @Produce json
// @Param input body models.UpdatePasswordReq true "current and new password"
// @Success 200
// @Failure 400,404 {string} string "message"
// @Failure 500 {string} string "message"
// @Failure default {string} string "message"
// @Router /update-password [post]
func (h Manager) UpdatePassword(c echo.Context) error {
	var req models.UpdatePasswordReq
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.srv.User.UpdatePassword(c.Request().Context(), &req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

// @Summary GetUsersWithBooks
// @Tags user
// @Description get users with books they borrowed
// @ID get-users-with-books
// @Accept json
// @Produce json
// @Success 200 {object} []models.UserBook
// @Failure 400,404 {string} string "message"
// @Failure 500 {string} string "message"
// @Failure default {string} string "message"
// @Router /users-with-books [get]
func (h Manager) GetUsersWithBooks(c echo.Context) error {
	resp, err := h.srv.User.GetUsersWithBooks(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

// @Summary GetUsersWithBooksForMonth
// @Tags user
// @Description get users with books they borrowed for last month
// @ID get-users-with-books-for-month
// @Accept json
// @Produce json
// @Success 200 {object} []models.UserBook
// @Failure 400,404 {string} string "message"
// @Failure 500 {string} string "message"
// @Failure default {string} string "message"
// @Router /users-with-books-for-month [get]
func (h Manager) GetUsersWithBooksForMonth(c echo.Context) error {
	resp, err := h.srv.User.GetUsersWithBooksForMonth(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}
