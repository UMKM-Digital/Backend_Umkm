package brandlogo

import "github.com/labstack/echo/v4"

type BrandLogoController interface {
	Create(c echo.Context) error
	// GetTransaksiList(c echo.Context) error
	// DeleteTestimonial(c echo.Context) error
	// UpdateTestimonial(c echo.Context) error
	// GetKategoriId(c echo.Context) error
}