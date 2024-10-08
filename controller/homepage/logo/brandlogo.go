package brandlogo

import "github.com/labstack/echo/v4"

type BrandLogoController interface {
	Create(c echo.Context) error
	GetBrandLogoList(c echo.Context) error
	DeleteProdukId(c echo.Context) error
	GetBrandLogoId(c echo.Context) error
	UpdateBrandLogo(c echo.Context) error
}