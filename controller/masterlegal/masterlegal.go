package masterlegalcontroller

import "github.com/labstack/echo/v4"

type ProdukController interface {
	CreateProduk(c echo.Context) error
	GetMasterLegalList(c echo.Context) error
	Delete( c echo.Context) error
	GetIdMasterLegalId(c echo.Context) error
}
