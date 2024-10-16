package web

import (
	"encoding/json"
	// "time"

	"github.com/shopspring/decimal"
)

type UmkmRequest struct {
	Name                  string          `json:"name"`
	NoNpwp                string          `json:"no_npwp"`
	Kategori_Umkm_Id      json.RawMessage `json:"kategori_umkm_id"`
	Nama_Penanggung_Jawab string          `json:"nama_penanggung_jawab"`
	Informasi_JamBuka     json.RawMessage `json:"informasi_jambuka"`
	No_Kontak             string          `json:"no_kontak"`
	Lokasi                string          `json:"lokasi"`
	Deskripsi                string          `json:"deskripsi"`
	Maps                  json.RawMessage `json:"maps"`
	Gambar                json.RawMessage `json:"gambar"`  // Field for images
	UserId                int             `json:"user_id"` // Add UserId field
	// SektorUsaha           string          `json:"sekto_usaha"`
	// StatusTempatUsaha     string		  `json:"status_tempat_usaha"`
	// KodeProv              string		  `json:"kode_prov"`
	// KodeKec               string 		  `json:"kode_kec"`
	// KodeKel               string		  `json:"kode_kelurahan"`
	// Rt 					  string          `json:"rt"`
	// Rw					  string		  `json:"rw"`
	// KodePos               string 		  `json:"kode_pos"`
	// NoNpwd                string    	  `json:"no_npwd"`
	// BahanBakar       	  string 		  `json:"bahan_bakar"`
	// TanggalMulaiUsaha    time.Time 		  `json:"tanggal_mulai_usaha"`
	// Kapasitas			  string		  `json:"kapasitas"`
	// TenagaKerjaPria       int 			  `json:"tenaga_kerja_pria"`
	// TenagaKerjaWanita     int             `json:"tenaga_kerja_wanita"`
	// NominalAset           float64         `json:"nominal_aset"`
	// NominalSendiri           float64         `json:"nominal_sendiri"`
	// EkonomiKreatif           bool 		  `json:"ekonomi_kreatif"`
	// KriteriaUsaha            string	      `json:"kriteria_usaha"`
	Omset                 []OmsetRequest  `json:"omset"` 
}

type OmsetRequest struct {
	Tahun      string          `json:"tahun"`
	Bulan      string          `json:"bulan"`
	JumlahOmset decimal.Decimal `json:"jumlah_omset"`
}

type Updateumkm struct{
	Name                  string          `json:"name"`
	NoNpwp                string          `json:"no_npwp"`
	Kategori_Umkm_Id      json.RawMessage `json:"kategori_umkm_id"`
	Nama_Penanggung_Jawab string          `json:"nama_penanggung_jawab"`
	Informasi_JamBuka     json.RawMessage `json:"informasi_jambuka"`
	No_Kontak             string          `json:"no_kontak"`
	Lokasi                string          `json:"lokasi"`
	Deskripsi                string          `json:"deskripsi"`
	Maps                  json.RawMessage `json:"maps"`
	Gambar                json.RawMessage `json:"gambar"`
}

type UpdateActiveUmkm struct{
	Active int `json:"active"`
}