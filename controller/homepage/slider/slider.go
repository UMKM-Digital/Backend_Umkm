package slidercontroller

import "github.com/labstack/echo/v4"

type Slider interface {
	Create(c echo.Context) error
	List(c echo.Context) error
	GetSlideId(c echo.Context) error
	DelSlideId(c echo.Context) error
}