package userservice

import (
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
		db: db,
	}
}

// register
func (service *AuthServiceImpl) RegisterRequest(user web.RegisterRequest) (map[string]interface{}, error) {
	// Hash password menggunakan bcrypt
	passHash, errHash := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if errHash != nil {
		return nil, errHash
	}
	user.Password = string(passHash)

	// Membuat object User baru
	newUser := domain.Users{
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		Role:     "umkm",
		No_Phone: user.No_Phone,
	}

	// Menyimpan user ke database
	saveUser, errSaveUser := service.authrepository.RegisterRequest(newUser)
	if errSaveUser != nil {
		return nil, errSaveUser
	}

	// Membuat claims untuk token JWT
	claims := helper.JwtCustomClaims{
		ID:      strconv.Itoa(saveUser.IdUser),  // Menggunakan ID dari saveUser setelah disimpan ke DB
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
		ID:    strconv.Itoa(user.IdUser),
		Name:  user.Username,
		Email: user.Email,
		Phone: user.No_Phone,
		Role: user.Role,
		Picture: user.Picture,
	}

	token, tokenErr := service.tokenUseCase.GenerateAccessToken(claims)
	if tokenErr != nil {
		return nil, tokenErr
	}

	// Hitung waktu kedaluwarsa token
	expirationTime := time.Now().Add(1 * time.Hour).Format(time.RFC3339)
	return map[string]interface{}{
		"token":      token,
		"expires_at": expirationTime, // Sertakan waktu kedaluwarsa yang sebenarnya
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

//verify
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

	// Token
	claims := helper.JwtCustomClaims{
		ID:    strconv.Itoa(user.IdUser),
		Name:  user.Username,
		Email: user.Email,
		Phone: user.No_Phone,
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
		"token": token,
		"expired time": expirationTime,
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

//verify otp register
func (service *AuthServiceImpl) VerifyOTPRegister(otp_code string, phone_code string)(map[string]interface{}, error) {
    // Verifikasi OTP
    isValid, err := helper.VerifyOTP(service.db, otp_code, phone_code)
    if err != nil || !isValid {
        return nil, errors.New("invalid OTP")
    }

    return map[string]interface{}{
        "message": "OTP verified successfully",
    }, nil
}
