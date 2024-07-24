package usercontroller

import (
	"net/http"
	"umkm/helper"
	"umkm/model"
	"umkm/model/web"
	userservice "umkm/service/user"

	"github.com/labstack/echo/v4"
)

type UserControllerImpl struct {
	userService userservice.AuthUserService
}

func NewAuthController(service userservice.AuthUserService) *UserControllerImpl {
	return &UserControllerImpl{
		userService: service,
	}
}

func (controller *UserControllerImpl) Register(c echo.Context) error {
	user := new(web.RegisterRequest)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(user); err != nil {
		return err
	}

	userUser, errSaveuser := controller.userService.RegisterRequest(*user)

	if errSaveuser != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, errSaveuser.Error(), nil))
	}
	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Register Success", userUser))
}


func (controller *UserControllerImpl) Login(c echo.Context) error {
    user := new(web.LoginRequest)

    if err := c.Bind(&user); err != nil {
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
    }
    userRes, errLogin := controller.userService.LoginRequest(user.Email, user.Password, user.No_Phone)

    if errLogin != nil {
        
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, errLogin.Error(), nil))
    }

    return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "login berhasil", userRes))
}


func (controller *UserControllerImpl) SendOtp(c echo.Context) error {
	user := new(web.OtpRequest)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseToJsonOtp(http.StatusBadRequest, err.Error(), nil))
	}

	otpResponse, err := controller.userService.SendOtp(user.No_Phone)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseToJsonOtp(http.StatusInternalServerError, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseToJsonOtp(http.StatusOK, "login berhasil", otpResponse))
}