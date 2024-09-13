package dokumenumkmcontroller

import "github.com/labstack/echo/v4"

type DokumenController interface {
	Create(c echo.Context) error
}