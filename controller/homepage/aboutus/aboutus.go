package aboutuscontroller

import "github.com/labstack/echo/v4"

type AboutUsController interface {
	Create(c echo.Context) error
	GetAboutUs(c echo.Context) error
	UpdateAboutUs(c echo.Context) error
	GetAboutId(c echo.Context) error
}