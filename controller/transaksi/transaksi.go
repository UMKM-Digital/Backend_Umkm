package transaksicontroller

import "github.com/labstack/echo/v4"

type UmkmController interface {
	Create(c echo.Context) error
	GetKategoriId(c echo.Context) error
	GetTransaksiFilterList(c echo.Context) error
}