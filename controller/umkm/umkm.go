package umkmcontroller

import "github.com/labstack/echo/v4"

type UmkmController interface {
	Create(c echo.Context) error
	GetUmkmList(c echo.Context) error
	GetUmkmFilter(c echo.Context) error 
	GetUmkmListWeb(c echo.Context) error
}