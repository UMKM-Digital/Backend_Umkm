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
	Ktp string `json:"ktp"`
	Kk string `json:"kk"`
}

func ToUserEntity(user domain.Users) UserEntity {
	return UserEntity{
		Id: user.IdUser,
	    Fullname: user.Fullname,
		Email:   user.Email,
		Picture: user.Picture,
		Role: user.Role,
		NoKk: user.NoKk,
		NoNik: user.Nik,
		TanggalLahir: user.TanggalLahir,
		JenisKelamin: user.JenisKelamin,
		PendidikanTerakhir: user.PendidikanTerakhir,
		StatusMenikah: user.StatusMenikah,
		Kecamatan: user.Kecamatan,
		Kelurahan: user.Kelurahan,
		KodePos: user.KodePos,
		Rt: user.Rt,
		Rw: user.Rw,
		Alamat: user.Alamat,
		NoHp: user.No_Phone,
		Ktp: user.Ktp,
		Kk: user.KartuKeluarga,
	}
}

//untuk list user

type UserEntityList struct {
	Id        int    `json:"id"`
	Fullname  string `json:"fullname"`
	NoPhone string `json:"no_phone"`
	Role string `json:"role"`
}

func ToUserEntityList(user domain.Users) UserEntityList {
	return UserEntityList{
		Id: user.IdUser,
	    Fullname: user.Fullname,
		NoPhone: user.No_Phone,
		Role: user.Role,
	}
}

func ToUserEntitiesList(userList []domain.Users) []UserEntityList {
	var userListEntities []UserEntityList
	for _, user := range userList {
		userListEntities = append(userListEntities, ToUserEntityList(user))
	}
	return userListEntities
}
