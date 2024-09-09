package produkcontroller

import "github.com/labstack/echo/v4"

type ProdukController interface {
	CreateProduk(c echo.Context) error
	DeleteProdukId(c echo.Context) error
	// GetProdukId(c echo.Context) error
	GetprodukList(c echo.Context) error 
	// UpdateProduk(c echo.Context) error
}