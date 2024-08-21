package brandlogo

import "github.com/labstack/echo/v4"

type BrandLogoController interface {
	Create(c echo.Context) error
	GetBrandLogoList(c echo.Context) error
	DeleteProdukId(c echo.Context) error
	// DeleteTestimonial(c echo.Context) error
	// UpdateTestimonial(c echo.Context) error
	// GetKategoriId(c echo.Context) error
}