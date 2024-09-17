package beritacontroller

import "github.com/labstack/echo/v4"

type UmkmController interface {
	Create(c echo.Context) error
	LIst(c echo.Context) error
	Delete( c echo.Context) error
	GetId(c echo.Context) error
	Update(c echo.Context) error
}
