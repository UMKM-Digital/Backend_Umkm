package web

type RegisterRequest struct {
	Fullname           string `validate:"required" json:"fullname"`       //
	Email              string `validate:"required,email" json:"email"`    //
	Password           string `validate:"required" json:"password"`       //
	No_Nik             string `validate:"required" json:"no_nik"`         //
	No_Phone           string `validate:"required" json:"no_telp"`        //
	No_KK              string `validate:"required" json:"no_kk"`          //
	// No_Nib             string `json:"no_nib"`         //
	TanggalLahir       string `validate:"required" json:"tgl_lahir"`      //
	JenisKelamin       string `validate:"required" json:"jenis_kelamin"`  //
	StatusMenikah      string `validate:"required" json:"status_menikah"` //
	Alamat             string `validate:"required" json:"alamat"`         //
	Provinsi           string `validate:"required" json:"provinsi"`       //
	Kabupaten          string `validate:"required" json:"kabupaten"`      //
	Kecamatan          string `validate:"required" json:"kecamatan"`      //
	Kelurahan          string `validate:"required" json:"kelurahan"`      //
	Rt                 string `validate:"required" json:"rt"`             //
	Rw                 string `validate:"required" json:"rw"`             //
	PendidikanTerakhir string `validate:"required" json:"pendidikan_terakhir"`
	KodePos            string `validate:"required" json:"kode_pos"`
}

type LoginRequest struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
}

// web/update_user_request.go
type UpdateUserRequest struct {
	Fullname           string  `validate:"required" json:"fullname"`       //o
	Email              string  `validate:"required,email" json:"email"`    //0
	Password           string  `validate:"required" json:"password"`       //0
	No_Nik             string  `validate:"required" json:"no_nik"`         //00
	No_Phone           string  `validate:"required" json:"no_telp"`        //
	No_KK              string  `validate:"required" json:"no_kk"`          //00
	// No_Nib             string  `json:"no_nib"`         //00
	TanggalLahir       string  `validate:"required" json:"tgl_lahir"`      //00
	JenisKelamin       string  `validate:"required" json:"jenis_kelamin"`  //00
	StatusMenikah      string  `validate:"required" json:"status_menikah"` //00
	Alamat             string  `validate:"required" json:"alamat"`         //0
	Provinsi           string  `validate:"required" json:"provinsi"`       //0
	Kabupaten          string  `validate:"required" json:"kabupaten"`      //0
	Kecamatan          string  `validate:"required" json:"kecamatan"`      //0
	Kelurahan          string  `validate:"required" json:"kelurahan"`      //0
	Rt                 string  `validate:"required" json:"rt"`             //0
	Rw                 string  `validate:"required" json:"rw"`             //0
	PendidikanTerakhir string  `validate:"required" json:"pendidikan_terakhir"`
	KodePos            string  `validate:"required" json:"kode_pos"`
	Picture            string  ` json:"picture"`
	Ktp                string `json:"ktp"`
    Kk                string `json:"kk"`

}

type OtpRequest struct {
	No_Phone string `validate:"required" json:"no_phone"`
}

type User struct{
	ID          uint   `json:"id"`
	PhoneNumber string `json:"phone_number"`
	// Tambahkan field lain sesuai kebutuhan
}

type CekPassword struct {
	Password string `validate:"required" json:"password"`
}

type ResetPasswordRequest struct {
	Email string `validate:"required,email" json:"email"`
}

type VerifyOtp struct {
	Phone string `json:"phone_number"`
	OTP   string `json:"otp_code"`
}