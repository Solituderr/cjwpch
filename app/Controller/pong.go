package Controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetRes(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}
