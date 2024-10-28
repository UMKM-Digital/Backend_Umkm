package omsetcontroller

import "github.com/labstack/echo/v4"

type omsetcontroller interface {
	createOmsetcontroller(c echo.Context) error
	LisOmsetController(c echo.Context) error
	GetOmsetController(c echo.Context) error
}
