package userservice

import (
	
	"umkm/model/web"
)

type AuthUserService interface {
	RegisterRequest(user web.RegisterRequest)(map[string]interface{}, error)
	LoginRequest(email string, password string, no_phone string) (map[string]interface{}, error)
	SendOtp(phone string) (map[string]interface{}, error)
}
