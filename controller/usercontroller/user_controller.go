package usercontroller

import "github.com/labstack/echo/v4"

type SellerController interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
	SendOtp(c echo.Context) error
	Logout(c echo.Context) error
	View(c echo.Context) error
	// Update(c echo.Context) error
	VerifyOTP(c echo.Context) error
	SendOtpRegister(c echo.Context) error
	VerifyOTPHandlerRegister(c echo.Context) error
}