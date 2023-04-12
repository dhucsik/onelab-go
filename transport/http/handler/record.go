package handler

import (
	"net/http"
	"practice/models"

	"github.com/labstack/echo"
)

func (h Manager) CreateRecord(c echo.Context) error {
	var req models.Record
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	resp, err := h.srv.Record.Create(c.Request().Context(), &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, resp)
}

func (h Manager) UpdateRecord(c echo.Context) error {
	var req models.Record
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := h.srv.Record.Update(c.Request().Context(), &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (h Manager) ListRecords(c echo.Context) error {
	resp, err := h.srv.Record.List(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

func (h Manager) GetRecord(c echo.Context) error {
	id := c.Param("id")

	resp, err := h.srv.Record.Get(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

func (h Manager) DeleteRecord(c echo.Context) error {
	id := c.Param("id")

	err := h.srv.Record.Delete(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
