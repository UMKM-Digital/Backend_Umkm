package helper

import (
	"encoding/json"
	"umkm/model/domain"
)

// IsUserDataComplete memeriksa apakah semua field yang diperlukan pada pengguna tidak kosong
func IsUserDataComplete(user domain.Users) bool {
	 // Cek apakah KTP dan Kartu Keluarga adalah JSONB yang tidak kosong
	   // Cek apakah KTP dan Kartu Keluarga adalah JSONB yang tidak kosong
	  // Cek apakah KTP dan Kartu Keluarga adalah JSONB yang tidak kosong
	  var ktpData domain.JSONB
	  var kartuKeluargaData domain.JSONB
  
	  // Mengonversi JSONB ke []byte dan kemudian unmarshal
	  ktpBytes, _ := user.Ktp.Value() // Dapatkan nilai KTP
	  kartuKeluargaBytes, _ := user.KartuKeluarga.Value() // Dapatkan nilai Kartu Keluarga
  
	  // Unmarshal untuk memeriksa apakah JSONB tidak kosong
	  if err := json.Unmarshal(ktpBytes.([]byte), &ktpData); err != nil {
		  ktpData = nil // Atur ke nil jika unmarshalling gagal
	  }
	  if err := json.Unmarshal(kartuKeluargaBytes.([]byte), &kartuKeluargaData); err != nil {
		  kartuKeluargaData = nil // Atur ke nil jika unmarshalling gagal
	  }
  
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
		   ktpData != nil &&               // Periksa apakah KTP tidak kosong
           kartuKeluargaData != nil &&  
		   user.Alamat != "" &&
		   user.Provinsi != "" &&
		   user.Kabupaten != "" &&
		   user.Kecamatan != "" &&
		   user.Kelurahan != "" &&
		   user.PendidikanTerakhir != "" 
    // Tambahkan kondisi lainnya jika diperlukan
}
