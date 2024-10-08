package usercontroller

import (
	"net/http"
	// "strings"
	"umkm/helper"
	"umkm/model"
	"umkm/model/web"
	userservice "umkm/service/user"

	"github.com/labstack/echo/v4"
)

type UserControllerImpl struct {
	userService  userservice.AuthUserService
	tokenUseCase helper.TokenUseCase
}

func NewAuthController(service userservice.AuthUserService, tokenUseCase helper.TokenUseCase) *UserControllerImpl {
	return &UserControllerImpl{
		userService:  service,
		tokenUseCase: tokenUseCase,
	}
}

func (controller *UserControllerImpl) Register(c echo.Context) error {
	user := new(web.RegisterRequest)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, err.Error(), nil))
	}

	if err := c.Validate(user); err != nil {
		return err
	}

	userUser, errSaveuser := controller.userService.RegisterRequest(*user)

	if errSaveuser != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, errSaveuser.Error(), nil))
	}
	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "register berhasil", userUser))
}

func (controller *UserControllerImpl) Login(c echo.Context) error {
	user := new(web.LoginRequest)

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, err.Error(), nil))
	}
	userRes, errLogin := controller.userService.LoginRequest(user.Username, user.Password)

	if errLogin != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, errLogin.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "login berhasil", userRes))
}

// send otp login
func (controller *UserControllerImpl) SendOtp(c echo.Context) error {
	user := new(web.OtpRequest)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, err.Error(), nil))
	}

	otpResponse, err := controller.userService.SendOtp(user.No_Phone)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "berasil send otp login", otpResponse))
}

// melihat isi profile
func (controller *UserControllerImpl) View(c echo.Context) error {
	adminID, _ := helper.GetAuthId(c)

	result, err := controller.userService.ViewMe(adminID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "Profile dapat dilihat", result))
}

//logout

// func (controller *UserControllerImpl) Logout(c echo.Context) error {
//     authHeader := c.Request().Header.Get("Authorization")
//     if authHeader == "" {
//         return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, "Authorization header is required", nil))
//     }

//     // Extract the token from the header
//     tokenString := strings.TrimPrefix(authHeader, "Bearer ")

//     // Blacklist the token
//     err := controller.tokenUseCase.BlacklistAccessToken(tokenString)
//     if err != nil {
//         return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, err.Error(), nil))
//     }

//     // Return a success message without a new token
//     return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "You have been logged out", nil))
// }

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

// verivy otp login
func (controller *UserControllerImpl) VerifyOTPHandler(c echo.Context) error {
	// Bind request body to a struct
	var req struct {
		Phone string `json:"phone_number"`
		OTP   string `json:"otp_code"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "Invalid request",
			"code":    http.StatusBadRequest,
		})
	}

	// Call the AuthService to verify OTP and phone number
	result, err := controller.userService.VerifyOTP(req.Phone, req.OTP)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  false,
			"message": "Verification failed",
			"code":    http.StatusUnauthorized,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "Verification successful",
		"data":    result,
		"code":    http.StatusOK,
	})
}

// sendotp register
func (controller *UserControllerImpl) SendOtpRegister(c echo.Context) error {
	user := new(web.OtpRequest)

	// Bind request body ke user
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, err.Error(), nil))
	}

	// Panggil service untuk mengirim OTP
	otpResponse, err := controller.userService.SendOtpRegister(user.No_Phone)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, err.Error(), nil))
	}

	// Cek apakah nomor telepon sudah terdaftar
	if otpResponse["message"] == "Phone number already registered" {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, "OTP tidak terkirim", otpResponse))
	}

	// Jika OTP berhasil dikirim, kembalikan status 200
	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "OTP terkirim", otpResponse))
}

// verify otp register
func (controller *UserControllerImpl) VerifyOTPHandlerRegister(c echo.Context) error {
	// Bind request body to a struct
	var req struct {
		Phone string `json:"phone_number"`
		OTP   string `json:"otp_code"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "Invalid request",
			"code":    http.StatusBadRequest,
		})
	}

	// Call the AuthService to verify OTP and phone number
	result, err := controller.userService.VerifyOTPRegister(req.Phone, req.OTP)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  false,
			"message": "Verification failed",
			"code":    http.StatusUnauthorized,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "Verification successful",
		"data":    result,
		"code":    http.StatusOK,
	})
}

func (controller *UserControllerImpl) CekPassword(c echo.Context) error {
    user := new(web.CekPassword)

    if err := c.Bind(&user); err != nil {
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, err.Error(), nil))
    }

    // Mendapatkan user ID dari token JWT yang sedang login
    authID, errAuthID := helper.GetAuthId(c)
    if errAuthID != nil {
        return c.JSON(http.StatusUnauthorized, model.ResponseToClient(http.StatusUnauthorized, false, "unauthorized", nil))
    }

    // Memeriksa apakah password sesuai dengan user yang sedang login
    userRes, errLogin := controller.userService.CekInRequest(authID, user.Password)
    if errLogin != nil {
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, errLogin.Error(), nil))
    }

    return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "password correct", userRes))
}

func (controller *UserControllerImpl) ChangePassword(c echo.Context) error {
    // Struct untuk menerima input password lama dan baru dari request
    var passwordChangeRequest struct {
        OldPassword string `json:"old_password"`
        NewPassword string `json:"new_password"`
    }

    // Bind input dari request ke struct
    if err := c.Bind(&passwordChangeRequest); err != nil {
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, err.Error(), nil))
    }

    // Dapatkan authID dari JWT token yang aktif
    authID, errAuthID := helper.GetAuthId(c)
    if errAuthID != nil {
        return c.JSON(http.StatusUnauthorized, model.ResponseToClient(http.StatusUnauthorized, false, "unauthorized", nil))
    }

    // Panggil service untuk mengubah password
    if err := controller.userService.ChangePassword(authID, passwordChangeRequest.OldPassword, passwordChangeRequest.NewPassword); err != nil {
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, err.Error(), nil))
    }

    return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "password changed successfully", nil))
}



// func (controller *UserControllerImpl) LoginWithGoogle(c echo.Context) error {
//     var request struct {
//         GoogleToken string `json:"google_token"`
//     }

//     if err := c.Bind(&request); err != nil {
//         return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "invalid request", nil))
//     }

//     // Verifikasi token Google dan login
//     user, jwtToken, err := controller.userService.LoginWithGoogle(request.GoogleToken)
//     if err != nil {
//         return c.JSON(http.StatusUnauthorized, model.ResponseToClient(http.StatusUnauthorized, false, err.Error(), nil))
//     }

//     return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "login successful", map[string]interface{}{
//         "user":  user,
//         "token": jwtToken,
//     }))
// }


// HandleGoogleLoginOrRegister untuk menangani login atau pendaftaran menggunakan Google
func (controller *UserControllerImpl) HandleGoogleLoginOrRegister(c echo.Context) error {
    var request struct {
        GoogleID string `json:"google_id"` // ID Google pengguna
        Email    string `json:"email"`     // Email pengguna
        Username string `json:"username"`  // Nama pengguna
        Picture   string `json:"picture"`   // URL gambar profil
    }

    // Bind JSON request ke struct
    if err := c.Bind(&request); err != nil {
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "invalid request", nil))
    }

    // Panggil HandleGoogleLoginOrRegister untuk mencari atau membuat pengguna
    userInfo, err := controller.userService.HandleGoogleLoginOrRegister(request.GoogleID, request.Email, request.Username, request.Picture)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, err.Error(), nil))
    }

    return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "login or registration successful", userInfo))
}
