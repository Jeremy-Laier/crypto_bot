package server

import "github.com/labstack/echo/v4"

type Server interface {
	HandlePing(c echo.Context) error
}
