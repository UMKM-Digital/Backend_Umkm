package web

import (
	"mime/multipart" // Tambahkan ini
	"time"
	// Paket lain yang diperlukan
)


type RegisterRequest struct {
    Fullname string `validate:"required" json:"fullname"`//
	Email    string `validate:"required,email" json:"email"`//
	Password string `validate:"required" json:"password"`//
    No_Nik string `validate:"required" json:"no_nik"`//
	No_Phone string `validate:"required" json:"no_telp"`//
    No_KK string    `validate:"required" json:"no_kk"`//
    No_Nib string `validate:"requireed" json:"no_nib"`//
    TanggalLahir time.Time `validate:"requireed" json:"tgl_lahir"`//
    JeniKelamin string `validate:"requireed" json:"jenis_kelamin"`//
    StatusMenikah string `validate:"requireed" json:"status_menikah"`//
    Alamat string `validate:"requireed" json:"alamat"`//
    Provinsi string `validate:"requireed" json:"provinsi"`
    Kabupaten string `validate:"requireed" json:"kabupaten"`
    Kecamatan string `validate:"requireed" json:"kecamatan"`
    Kelurahan string `validate:"requireed" json:"kelurahan"`
    Rt string `validate:"requireed" json:"rt"`
    Rw string `validate:"requireed" json:"rw"`
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
