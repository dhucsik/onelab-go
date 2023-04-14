package handler

import (
	"net/http"
	"practice/models"

	"github.com/labstack/echo"
)

func (h Manager) SignUp(c echo.Context) error {
	var req models.User
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	resp, err := h.srv.User.SignUp(c.Request().Context(), &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

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

func (h Manager) GetUsersWithBooks(c echo.Context) error {
	resp, err := h.srv.User.GetUsersWithBooks(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

func (h Manager) GetUsersWithBooksForMonth(c echo.Context) error {
	resp, err := h.srv.User.GetUsersWithBooksForMonth(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}
