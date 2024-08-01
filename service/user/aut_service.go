package userservice

import (
	"umkm/model/entity"
	"umkm/model/web"
)

type AuthUserService interface {
	RegisterRequest(user web.RegisterRequest)(map[string]interface{}, error)
	LoginRequest(username string, password string) (map[string]interface{}, error)
	SendOtp(phone string) (map[string]interface{}, error)
	ViewMe(userId int) (entity.UserEntity, error)
	// Update(userId int, req web.UpdateUserRequest, profilePicturePath string) (helper.ResponseToJson, error
	VerifyOTP(phone string, otp_code string)(map[string]interface{}, error)
}
