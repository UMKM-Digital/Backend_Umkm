package web

import (
    "mime/multipart" // Tambahkan ini
    // Paket lain yang diperlukan
)


type RegisterRequest struct {
	Username string `validate:"required" json:"username"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
	// Role string `validate:"required" json:"role"`
	No_Phone string `validate:"required" json:"no_phone"`
}
type LoginRequest struct {
	Username    string `validate:"required,username" json:"username"`
	Password string `validate:"required" json:"password"`	
}

// web/update_user_request.go
type UpdateUserRequest struct {
    Username     string                `validate:"required" json:"username"`
    Email        string                `validate:"required,email" json:"email"`
    No_Phone     string                `validate:"required" json:"no_phone"`
    Picture      *multipart.FileHeader `json:"profile_picture,omitempty" form:"profile_picture"`
}


type OtpRequest struct {
	No_Phone string `validate:"required" json:"no_phone"`
}

type User struct {
    ID          uint   `json:"id"`
    PhoneNumber string `json:"phone_number"`
    // Tambahkan field lain sesuai kebutuhan
}

type CekPassword struct{
    Password string `validate:"required" json:"password"`	
}

type ResetPasswordRequest struct {
		Email    string `validate:"required,email" json:"email"`
}
