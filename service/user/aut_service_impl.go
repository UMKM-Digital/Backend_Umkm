package userservice

import (
	// "encoding/json"
	"errors"
	"strconv"
	"time"
	"umkm/helper"
	"umkm/model/domain"
	"umkm/model/entity"

	"umkm/model/web"
	"umkm/repository/userrepo"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"fmt"
	"net/smtp"
)

type AuthServiceImpl struct {
	authrepository userrepo.AuthUserRepo
	tokenUseCase   helper.TokenUseCase
	db             *gorm.DB // Tambahkan field ini
}

func Newauthservice(authrepository userrepo.AuthUserRepo, token helper.TokenUseCase, db *gorm.DB) *AuthServiceImpl {
	return &AuthServiceImpl{
		authrepository: authrepository,
		tokenUseCase:   token,
		db:             db,
	}
}

// register
// register
func (service *AuthServiceImpl) RegisterRequest(user web.RegisterRequest) (map[string]interface{}, error) {
    // Hash password menggunakan bcrypt
    passHash, errHash := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
    if errHash != nil {
        return nil, errHash
    }
    user.Password = string(passHash)

    // Parse TanggalLahir (Date of Birth) tanpa memanggil .String()
    tanggalLahirParsed, errDate := helper.ParseDateLahir(user.TanggalLahir)
    if errDate != nil {
        return nil, errDate
    }

    // Membuat object User baru
    newUser := domain.Users{
        Fullname:           user.Fullname,
        Password:           user.Password,
        Email:              user.Email,
        Role:               "umkm",
        No_Phone:           user.No_Phone,
        Nik:                user.No_Nik,
        NoKk:               user.No_KK,
        Nib:                user.No_Nib,
        TanggalLahir:       tanggalLahirParsed,
        JenisKelamin:       user.JenisKelamin,
        PendidikanTerakhir: user.PendidikanTerakhir,
        StatusMenikah:      user.StatusMenikah,
        Alamat:             user.Alamat,
        Provinsi:           user.Provinsi,
        Kabupaten:          user.Kabupaten,
        Kecamatan:          user.Kecamatan,
        Kelurahan:          user.Kelurahan,
		KodePos: user.KodePos,
        Rt:                 user.Rt,
        Rw:                 user.Rw,
    }

    // Menyimpan user ke database
    saveUser, errSaveUser := service.authrepository.RegisterRequest(newUser)
    if errSaveUser != nil {
        return nil, errSaveUser
    }

    // Membuat claims untuk token JWT
    claims := helper.JwtCustomClaims{
        ID:      strconv.Itoa(saveUser.IdUser), // Menggunakan ID dari saveUser setelah disimpan ke DB
        Name:    saveUser.Username,
        Email:   saveUser.Email,
        Phone:   saveUser.No_Phone,
        Role:    saveUser.Role,
        Picture: saveUser.Picture,
    }

    // Menghasilkan token JWT
    token, tokenErr := service.tokenUseCase.GenerateAccessToken(claims)
    if tokenErr != nil {
        return nil, tokenErr
    }

    // Menghitung waktu kedaluwarsa token
    expirationTime := time.Now().Add(1 * time.Hour).Format(time.RFC3339)

    // Mengembalikan token dan informasi user
    return map[string]interface{}{
        "token":      token,
        "expires_at": expirationTime, // Sertakan waktu kedaluwarsa yang sebenarnya
    }, nil
}


func (service *AuthServiceImpl) LoginRequest(username string, password string) (map[string]interface{}, error) {
	user, getUserErr := service.authrepository.FindUserByUsername(username)
	if getUserErr != nil {
		return nil, errors.New("username not found")
	}
	if checkPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); checkPassword != nil {
		return nil, errors.New("incorrect password")
	}

	claims := helper.JwtCustomClaims{
		ID:      strconv.Itoa(user.IdUser),
		Name:    user.Username,
		Email:   user.Email,
		Phone:   user.No_Phone,
		Role:    user.Role,
		Picture: user.Picture,
	}

	isDataComplete := helper.IsUserDataComplete(*user)

    
	token, tokenErr := service.tokenUseCase.GenerateAccessToken(claims)
	if tokenErr != nil {
		return nil, tokenErr
	}

	// Hitung waktu kedaluwarsa token
	expirationTime := time.Now().Add(1 * time.Hour).Format(time.RFC3339)
	return map[string]interface{}{
		"token":      token,
		"expires_at": expirationTime, // Sertakan waktu kedaluwarsa yang sebenarnya
		"data": isDataComplete,
	}, nil
}

func (service *AuthServiceImpl) SendOtp(phone string) (map[string]interface{}, error) {
	// Temukan pengguna berdasarkan nomor telepon
	_, err := service.authrepository.FindUserByPhone(phone)
	if err != nil {
		return nil, errors.New("phone number not found")
	}

	// Generate OTP

	// Tentukan waktu kadaluarsa OTP
	expirationTime := time.Now().Add(1 * time.Minute)

	// Kirim OTP melalui WhatsApp dan simpan ke database
	if err := helper.SendWhatsAppOTP(service.db, phone, expirationTime); err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"message":    "OTP sent successfully",
		"expires_at": expirationTime.Format(time.RFC3339),
	}, nil
}

// get profile
func (service *AuthServiceImpl) ViewMe(userId int) (entity.UserEntity, error) {
	user, err := service.authrepository.GetByID(userId)
	if err != nil {
		return entity.UserEntity{}, err
	}

	return entity.ToUserEntity(user), nil
}

// update profile
// func (service *AuthServiceImpl) Update(userId int, req web.UpdateUserRequest, profilePicturePath string) (helper.ResponseToJson, error) {
//     user, errUser := service.authrepository.GetByID(userId)
//     if errUser != nil {
//         return nil, errUser
//     }

//     if req.Username != "" {
//         user.Username = req.Username
//     }
//     if req.Email != "" {
//         user.Email = req.Email
//     }
//     if req.No_Phone != "" {
//         user.No_Phone = req.No_Phone
//     }
//     if req.Password != "" {
//         passHash, errHash := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
//         if errHash != nil {
//             return nil, errHash
//         }
//         user.Password = string(passHash)
//     }
//     if profilePicturePath != "" {
//         user.Picture = profilePicturePath
//     }

//     result, errUpdate := service.authrepository.UpdateId(userId, user)
//     if errUpdate != nil {
//         return nil, errUpdate
//     }

//     data := helper.ResponseToJson{
//         "id":             result.IdUser,
//         "username":       result.Username,
//         "email":          result.Email,
//         "no_phone":       result.No_Phone,
//         "profile_picture": result.Picture,
//     }

//     return data, nil
// }

//

// verify
func (service *AuthServiceImpl) VerifyOTP(phone_number string, otpCode string) (map[string]interface{}, error) {
	// Verifikasi OTP
	isValid, err := helper.VerifyOTP(service.db, phone_number, otpCode)
	if err != nil || !isValid {
		return nil, errors.New("invalid OTP")
	}

	// Temukan pengguna berdasarkan nomor telepon
	user, err := service.authrepository.FindUserByPhone(phone_number)
	if err != nil {
		return nil, errors.New("user not found")
	}

	isDataComplete := helper.IsUserDataComplete(*user)

	// Token
	claims := helper.JwtCustomClaims{
		ID:      strconv.Itoa(user.IdUser),
		Name:    user.Username,
		Email:   user.Email,
		Phone:   user.No_Phone,
		Picture: user.Picture,
		Role:    user.Role,
	}

	token, tokenErr := service.tokenUseCase.GenerateAccessToken(claims)
	if tokenErr != nil {
		return nil, tokenErr
	}

	// Hitung waktu kedaluwarsa token
	expirationTime := time.Now().Add(1 * time.Hour).Format(time.RFC3339)

	// OTP berhasil diverifikasi dan password benar
	return map[string]interface{}{
		"message": "OTP verified",
		// "user":    user,
		"token":        token,
		"expired time": expirationTime,
		"data": isDataComplete,
	}, nil
}

func (service *AuthServiceImpl) SendOtpRegister(phone string) (map[string]interface{}, error) {
	user, err := service.authrepository.FindUserByPhoneRegister(phone)
	if phone == "" {
		return nil, errors.New("no telepon kosong")
	}
	if err != nil {
		return nil, err
	}

	if user != nil {
		return map[string]interface{}{
			"message": "Phone number already registered",
		}, nil
	}

	expirationTime := time.Now().Add(1 * time.Minute)

	if err := helper.SendWhatsAppOTP(service.db, phone, expirationTime); err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"message":    "otp terkirim",
		"expires_at": expirationTime.Format(time.RFC3339),
	}, nil
}

// verify otp register
func (service *AuthServiceImpl) VerifyOTPRegister(otp_code string, phone_code string) (map[string]interface{}, error) {
	// Verifikasi OTP
	isValid, err := helper.VerifyOTP(service.db, otp_code, phone_code)
	if err != nil || !isValid {
		return nil, errors.New("invalid OTP")
	}

	return map[string]interface{}{
		"message": "OTP verified successfully",
	}, nil
}

func (service *AuthServiceImpl) CekInRequest(authID int, password string) (map[string]interface{}, error) {
    user, getUserErr := service.authrepository.CekInPassword(authID)
    if getUserErr != nil {
        return nil, errors.New("user not found")
    }

    if checkPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); checkPassword != nil {
        return nil, errors.New("incorrect password")
    }

    // Return the user details or any necessary information upon successful password check
    return map[string]interface{}{
        "id":    user.IdUser,
        "name":  user.Username,
        "email": user.Email,
    }, nil
}


//updatepasssword
func (service *AuthServiceImpl) ChangePassword(authID int, oldPassword string, newPassword string) error {
    // Dapatkan user berdasarkan authID
    user, getUserErr := service.authrepository.CekInPassword(authID)
    if getUserErr != nil {
        return errors.New("user not found")
    }

    // Verifikasi apakah password lama sesuai
    if checkPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword)); checkPassword != nil {
        return errors.New("incorrect old password")
    }

    // Ubah ke password baru
    if err := service.authrepository.UpdatePassword(authID, newPassword); err != nil {
        return err
    }

    return nil
}

// //login google
// func (service *AuthServiceImpl) VerifyGoogleToken(token string) (*oauth2.Tokeninfo, error) {
//     ctx := context.Background()
//     oauth2Service, err := oauth2.NewService(ctx, option.WithHTTPClient(http.DefaultClient))
//     if err != nil {
//         return nil, err
//     }

//     tokenInfo, err := oauth2Service.Tokeninfo().IdToken(token).Do()
//     if err != nil {
//         return nil, err
//     }

//     return tokenInfo, nil
// }
// // File: service/auth_service.go
// func (service *AuthServiceImpl) LoginWithGoogle(token string) (*domain.Users, string, error) {
//     // Verifikasi token Google
//     tokenInfo, err := service.VerifyGoogleToken(token)
//     if err != nil {
//         return nil, "", errors.New("invalid google token")
//     }

// 	username := tokenInfo.Email

//     // Cari atau buat user berdasarkan googleID
//     user, err := service.authrepository.FindOrCreateUserByGoogleID(tokenInfo.UserId, tokenInfo.Email, username)
//     if err != nil {
//         return nil, "", err
//     }

//     // Buat claims untuk JWT
//     claims := helper.JwtCustomClaims{
//         ID:      strconv.Itoa(user.IdUser),
//         Name:    user.Username,
//         Email:   user.Email,
//         Phone:   user.No_Phone,
//         Picture: user.Picture,
//         Role:    user.Role,
//     }

//     // Panggil GenerateAccessToken untuk membuat JWT token
//     jwtToken, jwtErr := service.tokenUseCase.GenerateAccessToken(claims)
//     if jwtErr != nil {
//         return nil, "", jwtErr // Kembalikan error jika pembuatan token gagal
//     }

//     return user, jwtToken, nil // Kembalikan user dan token
// }


// File: service/auth_service.go


//
// File: service/auth_service.go



// // VerifyGoogleToken untuk memverifikasi token dan mendapatkan tokenInfo
// func (service *AuthServiceImpl) VerifyGoogleToken(idToken string) (map[string]interface{}, error) {
//     ctx := context.Background()

//     // Use idtoken.Validate to validate the Google ID token
//     tokenPayload, err := idtoken.Validate(ctx, idToken, "552769390995-9oggvci6d86aelv4ri7vlirrp08vti52.apps.googleusercontent.com") // Replace with your Google client ID
//     if err != nil {
//         return nil, fmt.Errorf("invalid google token: %w", err)
//     }

//     return tokenPayload.Claims, nil
// }


// // Mendapatkan informasi pengguna dari Userinfo API
// // Mendapatkan informasi pengguna dari Userinfo API
// // Mendapatkan informasi pengguna dari Userinfo API
// func (service *AuthServiceImpl) GetUserInfo(accessToken string) (*oauth2.Userinfo, error) {
//     ctx := context.Background()
//     oauth2Service, err := oauth2.NewService(ctx, option.WithHTTPClient(http.DefaultClient))
//     if err != nil {
//         return nil, err
//     }

//     // Membuat permintaan untuk mendapatkan informasi pengguna
//     userInfo, err := oauth2Service.Userinfo.Get().Context(ctx).Do()
//     if err != nil {
//         return nil, fmt.Errorf("failed to get user info: %w", err)
//     }

//     return userInfo, nil // Mengembalikan pointer ke objek Userinfo
// }



// // LoginWithGoogle untuk login menggunakan Google
// // LoginWithGoogle untuk login menggunakan Google
// func (service *AuthServiceImpl) LoginWithGoogle(token string) (*domain.Users, string, error) {
//     // Verifikasi token Google dan dapatkan klaim
//     claims, err := service.VerifyGoogleToken(token)
//     if err != nil {
//         return nil, "", fmt.Errorf("failed to verify google token: %w", err)
//     }

//     // Ekstrak user ID (sub), email, nama, dan phone dari klaim
//     googleID, _ := claims["sub"].(string)      // 'sub' adalah Google user ID
//     email, _ := claims["email"].(string)       // Ekstrak email
//     name, _ := claims["name"].(string)         // Ekstrak nama

//     // Cari atau buat user baru berdasarkan Google ID dan email di database
//     user, err := service.authrepository.FindOrCreateUserByGoogleID(googleID, email, name)
//     if err != nil {
//         return nil, "", err
//     }

//     // Buat JWT token untuk user
//     jwtToken, jwtErr := service.tokenUseCase.GenerateAccessToken(helper.JwtCustomClaims{
//         ID:    strconv.Itoa(user.IdUser),
//         Name:  user.Username,
//         Email: user.Email,
//         Role:  user.Role,
//     })
//     if jwtErr != nil {
//         return nil, "", jwtErr
//     }

//     return user, jwtToken, nil
// }

func (service *AuthServiceImpl) HandleGoogleLoginOrRegister(googleID string, email string, username string, picture string) (map[string]interface{}, error) {
	// Mencari atau membuat pengguna berdasarkan ID Google
	user, err := service.authrepository.FindOrCreateUserByGoogleID(googleID, email, username, picture)
	if err != nil {
		return nil, err
	}

	// Jika pengguna baru dibuat, kita hash password kosong (opsional)
	if user.GoogleId != "" {
		// Hash password menggunakan bcrypt (opsional)
		passHash, errHash := bcrypt.GenerateFromPassword([]byte("temporaryPassword"), bcrypt.MinCost) // Password sementara
		if errHash != nil {
			return nil, errHash
		}
		user.Password = string(passHash) // Simpan password yang di-hash
	}

	isDataComplete := helper.IsUserDataComplete(*user)


	// Membuat claims untuk token JWT
	claims := helper.JwtCustomClaims{
		ID:      strconv.Itoa(user.IdUser),
		Name:    user.Fullname,
		Email:   user.Email,
		Phone:   user.No_Phone,
		Role:    user.Role,
		Picture: user.Picture,
	}

	// Menghasilkan token JWT
	token, tokenErr := service.tokenUseCase.GenerateAccessToken(claims)
	if tokenErr != nil {
		return nil, tokenErr
	}

	// Menghitung waktu kedaluwarsa token
	expirationTime := time.Now().Add(1 * time.Hour).Format(time.RFC3339)

	// Mengembalikan token dan informasi pengguna
	return map[string]interface{}{
		"token":      token,
		"expires_at": expirationTime, // Sertakan waktu kedaluwarsa yang sebenarnya
		"data": isDataComplete,
	}, nil
}


// SendPasswordResetLink mengirimkan link reset password ke email yang diberikan
func (service *AuthServiceImpl) SendPasswordResetLink(email string) error {
	// Cek apakah email ada di database
	user, err := service.authrepository.ChangePassword(email)
	if err != nil {
		return fmt.Errorf("email not found: %v", err)
	}

	// Buat token reset password (misalnya JWT atau random string)
	resetToken := "23214jadkhb" // Ganti dengan fungsi untuk membuat token yang aman

	// Buat link reset password
	resetLink := fmt.Sprintf("https://yourapp.com/reset-password?token=%s", resetToken)

	// Waktu expired link (misal 1 jam)
	expirationTime := time.Now().Add(1 * time.Hour)

	// Simpan token dan waktu expired ke database jika diperlukan (opsional)

	// Kirim email dengan SMTP
	err = service.sendEmail(user.Email, resetLink, expirationTime)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}

func (service *AuthServiceImpl) sendEmail(recipientEmail, resetLink string, expirationTime time.Time) error {
	from := "adlisantosanaufal@gmail.com"
	password := "bnpp toam ocrd yuid"

	to := []string{recipientEmail}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte(fmt.Sprintf(
		"Subject: Password Reset Request\n\nClick the following link to reset your password:\n%s\n\nThis link will expire at %s.",
		resetLink, expirationTime.Format("02 Jan 2006 15:04")))

	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}