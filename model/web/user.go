package web

import (
    "mime/multipart" // Tambahkan ini
    // Paket lain yang diperlukan
)


type RegisterRequest struct {
	Username string `validate:"required" json:"username"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
	Role string `validate:"required" json:"role"`
	No_Phone string `validate:"required" json:"no_phone"`
}
type LoginRequest struct {
	No_Phone string `validate:"required" json:"no_phone"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`	
}

// web/update_user_request.go
type UpdateUserRequest struct {
    Username     string                `validate:"required" json:"username"`
    Email        string                `validate:"required,email" json:"email"`
    Password     string                `validate:"required" json:"password"`
    No_Phone     string                `validate:"required" json:"no_phone"`
    Picture      *multipart.FileHeader `json:"profile_picture,omitempty" form:"profile_picture"`
}


type OtpRequest struct {
	No_Phone string `validate:"required" json:"no_phone"`
}


