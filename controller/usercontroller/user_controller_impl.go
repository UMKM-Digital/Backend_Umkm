package usercontroller

import (
	"log"
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
	userRes, errLogin := controller.userService.LoginRequest(user.Username,  user.Password)

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
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, err.Error(), nil))
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

//update
func (controller *UserControllerImpl) Update(c echo.Context) error {
    // Get the user ID from the JWT token
    userId, _ := helper.GetAuthId(c)

    // Parse form-data values
    name := c.FormValue("fullname")
    email := c.FormValue("email")
    phoneNumber := c.FormValue("phone_number")
	Password := c.FormValue("password")
	NoNik    := c.FormValue("no_nik")
	no_kk    := c.FormValue("no_kk")
	// no_nib    := c.FormValue("no_nib")
	tanggalLahir    := c.FormValue("tanggal_lahir")
	jeniskelamin    := c.FormValue("jenis_kelamin")
	statusmenikah   := c.FormValue("status_menikah")
	provinsi    := c.FormValue("provinsi")
	kabupaten    := c.FormValue("kabupaten")
	kelurahan    := c.FormValue("kelurahan")
	rt    := c.FormValue("rt")
	rw    := c.FormValue("rw")
	pendidikanterakhir    := c.FormValue("pendidikan_terakhir")
	kodepos    := c.FormValue("kode_pos")
	kecamatan    := c.FormValue("kecamatan")
    address := c.FormValue("alamat")

		file, err := c.FormFile("potoprofile")
		if err == http.ErrMissingFile {
			file = nil // jika file tidak ada, atur nil
		} else if err != nil {
			return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Failed to get the uploaded file ktp", nil))
		}
		


		ktp, err := c.FormFile("ktp")
		if err == http.ErrMissingFile {
			ktp = nil // jika file tidak ada, atur nil
		} else if err != nil {
			return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Failed to get the uploaded file ktp", nil))
		}
		
		// Handle KK file
		kk, err := c.FormFile("kk")
		if err == http.ErrMissingFile {
			kk = nil // jika file tidak ada, atur nil
		} else if err != nil {
			return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Failed to get the uploaded file kk", nil))
		}

    // Create the request object manually
    request := web.UpdateUserRequest{
        Fullname:         name,
        Email:        email,
        No_Phone:  phoneNumber,
        Alamat:      address,
		Password: Password,
		No_Nik: NoNik,
		No_KK: no_kk,
		// No_Nib: no_nib,
		TanggalLahir: tanggalLahir,
		JenisKelamin: jeniskelamin,
		StatusMenikah: statusmenikah,
		Provinsi: provinsi,
		Kabupaten: kabupaten,
		Kecamatan: kecamatan,
		Kelurahan: kelurahan,
		Rt: rt,
		Rw: rw,
		PendidikanTerakhir: pendidikanterakhir,
		KodePos: kodepos,
    }

    // Call the service to update user info
    result, err := controller.userService.Update(userId, request, file,ktp, kk)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, helper.ResponseToJsonOtp(http.StatusInternalServerError, err.Error(), nil))
    }

    // Return success response
    return c.JSON(http.StatusOK, helper.ResponseToJsonOtp(http.StatusOK, "Successfully updated", result))
}

// verivy otp login
func (controller *UserControllerImpl) VerifyOTPHandler(c echo.Context) error {
	// Bind request body to a struct
	user := new(web.VerifyOtp)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, err.Error(), nil))
	}

	otpResponse, err := controller.userService.VerifyOTP(user.Phone, user.OTP)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "berasil send otp login", otpResponse))
}

// sendotp register
func (controller *UserControllerImpl) SendOtpRegister(c echo.Context) error {
    user := new(web.OtpRequest)

    if err := c.Bind(user); err != nil {
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, err.Error(), nil))
    }

    otpResponse, err := controller.userService.SendOtpRegister(user.No_Phone)
    if err != nil {
        // Cek pesan error apakah nomor telepon sudah terdaftar
        if err.Error() == "No Telepon telah terdaftar" {
            return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, err.Error(), nil))
        }
        return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, err.Error(), nil))
    }

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
			"message": "OTP tidak sesuai",
			"code":    http.StatusUnauthorized,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "Kode OTP sesuai",
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


// HandleGoogleLoginOrRegister untuk menangani login atau pendaftaran menggunakan Google
func (controller *UserControllerImpl) HandleGoogleLoginOrRegister(c echo.Context) error {
    var request struct {
        GoogleID string `json:"google_id"`
        Email    string `json:"email"`     
        Username string `json:"username"`  
        Picture   string `json:"picture"`   
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


//sendpasswor
// HandlePassworduser menangani permintaan reset password
func (controller *UserControllerImpl) HandlePasswordResetRequest(c echo.Context) error {
	// Buat variabel untuk menangkap request body
	var user web.ResetPasswordRequest

	// Bind request data ke struct
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request format",
			"error":   err.Error(),
		})
	}

	// Validasi email menggunakan validator
	if err := c.Validate(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid email format",
			"error":   err.Error(),
		})
	}

	// Panggil service untuk mengirimkan link reset password
	err := controller.userService.SendPasswordResetLink(user.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to send reset password link",
			"error":   err.Error(),
		})
	}

	// Sukses mengirim email
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Password reset link has been sent to your email",
	})
}


func (controller *UserControllerImpl) GetUser(c echo.Context) error {
	getUser, errGetUser := controller.userService.GetListUser()

	if errGetUser != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, errGetUser.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "success", getUser))
}


func (controller *UserControllerImpl) GetUserCountByGender(c echo.Context) error {
    // Memanggil service untuk menghitung jumlah pengguna berdasarkan gender
    result, err := controller.userService.CountUser()
    if err != nil {
        // Mengembalikan respons error jika terjadi kesalahan
        return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, err.Error(), nil))
    }

    // Mengembalikan respons sukses dengan data yang diperoleh
    return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "success", result))
}


func (controller *UserControllerImpl)DeleteUser(c echo.Context) error{
	log.Println("Checking userService:", controller.userService) // Debugging
	userId, _ := helper.GetAuthId(c)

	if errDeleteUser := controller.userService.DeleteUser(userId); errDeleteUser != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, errDeleteUser.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "Delete all data user Success", nil))
}