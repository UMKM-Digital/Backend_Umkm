package produkcontroller

import "github.com/labstack/echo/v4"

type ProdukController interface {
	CreateProduk(c echo.Context) error
}