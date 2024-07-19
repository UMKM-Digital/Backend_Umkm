package usercontroller

import "github.com/labstack/echo/v4"

type SellerController interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
	
}