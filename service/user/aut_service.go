package userservice

import (
	"time"
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
	SendOtpRegister(phone string) (map[string]interface{}, error)
	VerifyOTPRegister(otp_code string, phone_number string)(map[string]interface{}, error)
	CekInRequest(authID int, password string) (map[string]interface{}, error) 
	ChangePassword(authID int, oldPassword string, newPassword string) error
	// VerifyGoogleToken(idToken string) (map[string]interface{}, error)
	
	// LoginWithGoogle(token string) (*domain.Users, string, error) 
	HandleGoogleLoginOrRegister(googleID string, email string, username string, picture string) (map[string]interface{}, error)
	SendPasswordResetLink(email string) error
	sendEmail(recipientEmail, resetLink string, expirationTime time.Time) error
}
