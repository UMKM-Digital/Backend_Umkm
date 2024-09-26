package produkcontroller

import "github.com/labstack/echo/v4"

type ProdukController interface {
	CreateProduk(c echo.Context) error
	DeleteProdukId(c echo.Context) error
	GetProdukWebId(c echo.Context) error
	GetprodukListUmkm(c echo.Context) error 
	UpdateProduk(c echo.Context) error
	GetProdukList(c echo.Context) error
	GetProdukListWeb(c echo.Context) error

}