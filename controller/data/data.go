package datacontroller

import "github.com/labstack/echo/v4"

type SellerController interface {
	CountData(c echo.Context) error
}