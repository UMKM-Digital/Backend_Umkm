package homepagecontroller

import "github.com/labstack/echo/v4"

type KategoriUmkmController interface {
	Create(c echo.Context) error
	GetTransaksiList(c echo.Context) error
	DeleteTestimonial(c echo.Context) error 
}