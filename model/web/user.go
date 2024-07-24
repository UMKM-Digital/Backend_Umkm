package web

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

type UpdateUserRequest struct {
	Username     string `validate:"required" json:"name"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
}

type OtpRequest struct {
	No_Phone string `validate:"required" json:"no_phone"`
}


