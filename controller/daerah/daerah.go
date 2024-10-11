package daerahcontoller

import "github.com/labstack/echo/v4"

type KategoriUmkmController interface {
	GetDaerah(c echo.Context) error
	GetKabupaten(c echo.Context) error
	GetKecamatan(c echo.Context) error
	GetKelurahan(c echo.Context) error
}