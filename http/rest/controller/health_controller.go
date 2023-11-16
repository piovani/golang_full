package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthController struct{}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (h *HealthController) Health(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
}
