package kategoriprodukcontroller

import "github.com/labstack/echo/v4"

type KategoriUmkmController interface {
	Create(c echo.Context) error
	GetKategoriList(c echo.Context) error
	GetKategoriId(c echo.Context) error
	UpdateKategoriProduk(c echo.Context) error
	Delete( c echo.Context) error
}
