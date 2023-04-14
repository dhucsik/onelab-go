package handler

import (
	"net/http"
	"practice/models"

	"github.com/labstack/echo/v4"
)

// @Summary Create Record
// @Security ApiKeyAuth
// @Tags record
// @Description create record
// @ID create-record
// @Accept  json
// @Produce  json
// @Param input body models.Record true "record info"
// @Success 201 {string} string "id"
// @Failure 400,404 {string} string "message"
// @Failure 500 {string} string "message"
// @Failure default {string} string "message"
// @Router /record [post]
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

// @Summary Update Record
// @Security ApiKeyAuth
// @Tags record
// @Description update record
// @ID update-record
// @Accept json
// @Produce json
// @Param record_id path string true "record id"
// @Param input body models.Record true "record info"
// @Success 200
// @Failure 400,404 {string} string "message"
// @Failure 500 {string} string "message"
// @Failure default {string} string "message"
// @Router /record/{record_id} [put]
func (h Manager) UpdateRecord(c echo.Context) error {
	id := c.Param("id")

	var req models.Record
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := h.srv.Record.Update(c.Request().Context(), id, &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

// @Summary List Records
// @Tags record
// @Description list records
// @ID list-records
// @Accept json
// @Produce json
// @Success 200 {object} []models.Record
// @Failure 400,404 {object} string "message"
// @Failure 500 {object} string "message"
// @Failure default {object} string "message"
// @Router /record [get]
func (h Manager) ListRecords(c echo.Context) error {
	resp, err := h.srv.Record.List(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

// @Summary Get Record
// @Tags record
// @Description get record by ID
// @ID get-record
// @Accept json
// @Produce json
// @Param record_id path string true "record id"
// @Success 200 {object} models.Record
// @Failure 400,404 {object} string "message"
// @Failure 500 {object} string "message"
// @Failure default {object} string "message"
// @Router /record/{record_id} [get]
func (h Manager) GetRecord(c echo.Context) error {
	id := c.Param("id")

	resp, err := h.srv.Record.Get(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

// @Summary Delete record
// @Security ApiKeyAuth
// @Tags record
// @Description delete record by ID
// @ID delete-record
// @Accept json
// @Produce json
// @Param record_id path string true "record id"
// @Success 200
// @Failure 400,404 {object} string "message"
// @Failure 500 {object} string "message"
// @Failure default {object} string "message"
// @Router /record/{record_id} [delete]
func (h Manager) DeleteRecord(c echo.Context) error {
	id := c.Param("id")

	err := h.srv.Record.Delete(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
