package usercontroller

import (
	"net/http"
	"os"
	"time"
	"umkm/helper"
	"umkm/model"
	"umkm/model/web"
	userservice "umkm/service/user"

	"github.com/golang-jwt/jwt/v4"
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

// // UserControllerImpl.go
// func (controller *UserControllerImpl) Logout(c echo.Context) error {
// 	authHeader := c.Request().Header.Get("Authorization")
// 	if authHeader == "" {
// 		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, "Authorization header is required", nil))
// 	}

// 	// Membuat token kosong yang langsung kedaluwarsa
// 	expiredToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
// 		ExpiresAt: time.Now().Add(-1 * time.Second).Unix(), // Token langsung kedaluwarsa
// 	})
// 	tokenString, err := expiredToken.SignedString([]byte(os.Getenv("SECRET_KEY")))
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, "Could not create token", nil))
// 	}

// 	// Kirim token kedaluwarsa sebagai respons
// 	c.Response().Header().Set("Authorization", "Bearer "+tokenString)
// 	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "You have been logged out", nil))
// }

func (controller *UserControllerImpl) View(c echo.Context) error {
	adminID, _ := helper.GetAuthId(c)

	result, err := controller.userService.ViewMe(adminID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(),nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseToJsonOtp(http.StatusOK, "Success", result))
}

func (controller *UserControllerImpl) Logout(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, "Authorization header is required", nil))
	}

	// Create a token with a very short expiration time
	expirationTime := time.Now().Add(0 * time.Second)
	expiredToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	})
	tokenString, err := expiredToken.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, "Could not create token", nil))
	}

	// Return the short-lived token as part of the response
	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "You have been logged out", map[string]string{
		"token": tokenString,
	}))
}

//update
// controller/usercontroller/user_controller_impl.go
// func (controller *UserControllerImpl) Update(c echo.Context) error {
//     userId, _ := helper.GetAuthId(c)
//     request := new(web.UpdateUserRequest)
    
//     // Bind and validate request
//     if err := c.Bind(request); err != nil {
//         return c.JSON(http.StatusBadRequest, helper.ResponseToJsonOtp(http.StatusBadRequest, err.Error(), nil))
//     }
//     if err := c.Validate(request); err != nil {
//         return err
//     }

//     // Handle file upload
//     var profilePicturePath string
//     if request.Picture != nil {
//         src, err := request.Picture.Open()
//         if err != nil {
//             return err
//         }
//         defer src.Close()

//         // Create a destination file
//         profilePicturePath = filepath.Join("uploads", request.Picture.Filename)
//         dst, err := os.Create(profilePicturePath)
//         if err != nil {
//             return err
//         }
//         defer dst.Close()

//         // Copy the uploaded file to the destination
//         if _, err = io.Copy(dst, src); err != nil {
//             return err
//         }
//     }

//     // Call the service to update user info
//     result, err := controller.userService.Update(userId, *request, profilePicturePath)
//     if err != nil {
//         return c.JSON(http.StatusBadRequest, helper.ResponseToJsonOtp(http.StatusBadRequest, err.Error(), nil))
//     }

//     return c.JSON(http.StatusOK, helper.ResponseToJsonOtp(http.StatusOK, "Success", result))
// }
