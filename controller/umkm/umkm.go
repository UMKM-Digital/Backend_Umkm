package umkmcontroller

import "github.com/labstack/echo/v4"

type UmkmController interface {
	Create(c echo.Context) error
}