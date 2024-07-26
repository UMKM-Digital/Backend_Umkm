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
)

type AuthServiceImpl struct {
	authrepository userrepo.AuthUserRepo
	tokenUseCase   helper.TokenUseCase
}

func Newauthservice(authrepository userrepo.AuthUserRepo, token helper.TokenUseCase) *AuthServiceImpl {
	return &AuthServiceImpl{
		authrepository: authrepository,
		tokenUseCase:   token,
	}
}

// register
func (service *AuthServiceImpl) RegisterRequest(user web.RegisterRequest) (map[string]interface{}, error) {
	passHash, errHash := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if errHash != nil {
		return nil, errHash
	}

	user.Password = string(passHash)
	newUser := domain.Users{
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		Role:     user.Role,
		No_Phone: user.No_Phone,
	}

	saveUser, errSaveUser := service.authrepository.RegisterRequest(newUser)
	if errSaveUser != nil {
		return nil, errSaveUser
	}

	return helper.ResponseToJson{"username": saveUser.Username, "email": saveUser.Email}, nil
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
	_, err := service.authrepository.FindUserByPhone(phone)
	if err != nil {
		return nil, errors.New("phone number not found")
	}

	otp, otpErr := helper.GenerateOTP()
	if otpErr != nil {
		return nil, otpErr
	}

	if err := helper.SendWhatsAppOTP(phone, otp); err != nil {
		return nil, err
	}

	expirationTime := time.Now().Add(1 * time.Minute)
	return map[string]interface{}{
		"message":    "OTP sent successfully",
		"expires_at": expirationTime.Format(time.RFC3339),
		"otp":        otp,
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
