package helper

import (
	"umkm/model/domain"
)

// IsUserDataComplete memeriksa apakah semua field yang diperlukan pada pengguna tidak kosong
func IsUserDataComplete(user domain.Users) bool {
  
    return user.Fullname != "" && 
           user.No_Phone != "" && 
           user.NoKk != "" && 
           user.Nik != "" && 
           user.Nib != "" &&
		   user.Email != "" &&
			user.Picture != "" &&
		   !user.TanggalLahir.IsZero() && 
		   user.JenisKelamin != "" &&
		   user.Rt != "" &&
		   user.Rw != "" &&
		   user.KodePos != "" &&
		   user.Ktp != "" &&               // Periksa apakah KTP tidak kosong
           user.KartuKeluarga != "" &&  
		   user.Alamat != "" &&
		   user.Provinsi != "" &&
		   user.Kabupaten != "" &&
		   user.Kecamatan != "" &&
		   user.Kelurahan != "" &&
		   user.PendidikanTerakhir != "" 
    // Tambahkan kondisi lainnya jika diperlukan
}
