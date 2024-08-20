package usercontroller

import (
	"net/http"
	"strings"
	"umkm/helper"
	"umkm/model"
	"umkm/model/web"
	userservice "umkm/service/user"

	"github.com/labstack/echo/v4"
)

type UserControllerImpl struct {
	userService userservice.AuthUserService
	tokenUseCase helper.TokenUseCase

}

func NewAuthController(service userservice.AuthUserService, tokenUseCase helper.TokenUseCase) *UserControllerImpl {
	return &UserControllerImpl{
		userService: service,
		tokenUseCase: tokenUseCase,
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
    userRes, errLogin := controller.userService.LoginRequest(user.Username, user.Password,)

    if errLogin != nil {
        
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, errLogin.Error(), nil))
    }

    return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "login berhasil", userRes))
}

//send otp
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

//melihat isi profile
func (controller *UserControllerImpl) View(c echo.Context) error {
	adminID, _ := helper.GetAuthId(c)

	result, err := controller.userService.ViewMe(adminID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(),nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseToJsonOtp(http.StatusOK, "Success", result))
}

//logout

func (controller *UserControllerImpl) Logout(c echo.Context) error {
    authHeader := c.Request().Header.Get("Authorization")
    if authHeader == "" {
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, "Authorization header is required", nil))
    }

    // Extract the token from the header
    tokenString := strings.TrimPrefix(authHeader, "Bearer ")

    // Blacklist the token
    err := controller.tokenUseCase.BlacklistAccessToken(tokenString)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, err.Error(), nil))
    }

    // Return a success message without a new token
    return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "You have been logged out", nil))
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


//verivy otp
func (controller *UserControllerImpl) VerifyOTPHandler(c echo.Context) error {
	// Bind request body to a struct
	var req struct {
		Phone      string `json:"phone_number"`
		OTP        string `json:"otp_code"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request",
			"error":   err.Error(),
		})
	}

	// Call the AuthService to verify OTP and password
	result, err := controller.userService.VerifyOTP(req.Phone, req.OTP)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "Verification failed",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}


//sendotp register
func (controller *UserControllerImpl) SendOtpRegister(c echo.Context) error {
    user := new(web.OtpRequest)

    // Bind request body ke user
    if err := c.Bind(user); err != nil {
        return c.JSON(http.StatusBadRequest, helper.ResponseToJsonOtp(http.StatusBadRequest, err.Error(), nil))
    }

    // Panggil service untuk mengirim OTP
    otpResponse, err := controller.userService.SendOtpRegister(user.No_Phone)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, helper.ResponseToJsonOtp(http.StatusInternalServerError, err.Error(), nil))
    }

    // Cek apakah nomor telepon sudah terdaftar
    if otpResponse["message"] == "Phone number already registered" {
        return c.JSON(http.StatusInternalServerError, helper.ResponseToJsonOtp(http.StatusInternalServerError, "OTP tidak terkirim", otpResponse))
    }

    // Jika OTP berhasil dikirim, kembalikan status 200
    return c.JSON(http.StatusOK, helper.ResponseToJsonOtp(http.StatusOK, "OTP terkirim", otpResponse))
}



