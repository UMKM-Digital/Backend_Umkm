package entity

import (
	"time"
	"umkm/model/domain"
)

type UserEntity struct {
	Id        int    `json:"id"`
	Fullname  string `json:"fullname"`
	Email     string `json:"email"`
	Picture string `json:"picture"`
	Role string `json:"role"`
	NoKk string `json:"no_kk"`
	NoNik string `json:"no_nik"`
	NoNib string `json:"no_nib"`
	TanggalLahir time.Time `json:"tanggal_lahir"`
	JenisKelamin string `json:"jenis_kelamin"`
	PendidikanTerakhir string `json:"pendidikan_terakhir"`
	StatusMenikah string `json:"status_menikah"`
	Kecamatan string `json:"kecamatan"`
	Kelurahan string `json:"kelurahan"`
	KodePos string `json:"kode_pos"`
	Rt string `json:"rt"`
	Rw string `json:"rw"`
	Alamat string `json:"alamat"`
	NoHp string `json:"no_hp"`
	Ktp domain.JSONB `json:"ktp"`
	Kk domain.JSONB `json:"kk"`
}

func ToUserEntity(user domain.Users) UserEntity {
	return UserEntity{
		Id: user.IdUser,
	    Fullname: user.Fullname,
		Email:   user.Email,
		Picture: user.Picture,
		Role: user.Role,
	}
}