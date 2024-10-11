package web

import (
	"mime/multipart" // Tambahkan ini
	"time"
	// Paket lain yang diperlukan
)


type RegisterRequest struct {
    Fullname string `validate:"required" json:"fullname"`//o
	Email    string `validate:"required,email" json:"email"`//0
	Password string `validate:"required" json:"password"`//0
    No_Nik string `validate:"required" json:"no_nik"`//00
	No_Phone string `validate:"required" json:"no_telp"`//
    No_KK string    `validate:"required" json:"no_kk"`//00
    No_Nib string `validate:"requireed" json:"no_nib"`//00
    TanggalLahir time.Time `validate:"requireed" json:"tgl_lahir"`//00
    JeniKelamin string `validate:"requireed" json:"jenis_kelamin"`//00
    StatusMenikah string `validate:"requireed" json:"status_menikah"`//00
    Alamat string `validate:"requireed" json:"alamat"`//0
    Provinsi string `validate:"requireed" json:"provinsi"`//0
    Kabupaten string `validate:"requireed" json:"kabupaten"`//0
    Kecamatan string `validate:"requireed" json:"kecamatan"`//0
    Kelurahan string `validate:"requireed" json:"kelurahan"`//0
    Rt string `validate:"requireed" json:"rt"`//0
    Rw string `validate:"requireed" json:"rw"`//0
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
