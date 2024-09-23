package sektorusahacontroller

import "github.com/labstack/echo/v4"

type ProdukController interface {
	Create(c echo.Context) error
	GetSektorUsaha(c echo.Context) error
}