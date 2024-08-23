package homepagecontroller

import "github.com/labstack/echo/v4"

type KategoriUmkmController interface {
	Create(c echo.Context) error
	GetTestimonial(c echo.Context) error
	DeleteTestimonial(c echo.Context) error 
	UpdateTestimonial(c echo.Context) error
	GetKategoriId(c echo.Context) error
	GetTestimonialActive(c echo.Context) error
	// UpdateTestimonialActive(c echo.Context) error
}