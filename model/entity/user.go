package entity

import (
	"time"
	"umkm/model/domain"
	"umkm/repository/userrepo"
)

type Wilayah struct {
    Nama string `json:"nama"`
    Id   string `json:"id"`
}


type UserEntity struct {
	Id        int    `json:"id"`
	Fullname  string `json:"fullname"`
	Email     string `json:"email"`
	Picture string `json:"picture"`
	Role string `json:"role"`
	NoKk string `json:"no_kk"`
	NoNik string `json:"no_nik"`
	TanggalLahir time.Time `json:"tanggal_lahir"`
	JenisKelamin string `json:"jenis_kelamin"`
	PendidikanTerakhir string `json:"pendidikan_terakhir"`
	StatusMenikah string `json:"status_menikah"`
	Provinsi          Wilayah  `json:"provinsi"`
    Kabupaten         Wilayah  `json:"kabupaten"`
    Kecamatan         Wilayah  `json:"kecamatan"`
    Kelurahan         Wilayah  `json:"kelurahan"`
	KodePos string `json:"kode_pos"`
	Rt string `json:"rt"`
	Rw string `json:"rw"`
	Alamat string `json:"alamat"`
	NoHp string `json:"no_hp"`
	Ktp string `json:"ktp"`
	Kk string `json:"kk"`
}

func ToUserEntity(user domain.Users, repo userrepo.AuthUserRepo) UserEntity {
	provinsi, kabupaten, kecamatan, kelurahan, err := repo.GetNamaWilayah(user.Provinsi, user.Kabupaten, user.Kecamatan, user.Kelurahan)
    if err != nil {
        // Jika gagal, gunakan kode wilayah sebagai fallback
        provinsi = user.Provinsi
        kabupaten = user.Kabupaten
        kecamatan = user.Kecamatan
        kelurahan = user.Kelurahan
    }


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
		Provinsi: Wilayah{Nama: provinsi, Id: user.Provinsi},
        Kabupaten: Wilayah{Nama: kabupaten, Id: user.Kabupaten},
        Kecamatan: Wilayah{Nama: kecamatan, Id: user.Kecamatan},
        Kelurahan: Wilayah{Nama: kelurahan, Id: user.Kelurahan},
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
