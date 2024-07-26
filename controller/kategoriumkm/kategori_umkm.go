package kategoriumkmcontroller

import "github.com/labstack/echo/v4"

type KategoriUmkmController interface {
	Create(c echo.Context) error
	GetKategoriList(c echo.Context) error
	GetKategoriId(c echo.Context) error
	UpdateKategoriId(c echo.Context) error
}