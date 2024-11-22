package web

import (
	"encoding/json"
	"time"

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
	Deskripsi             string          `json:"deskripsi"`
	Maps                  json.RawMessage `json:"maps"`
	Gambar                json.RawMessage `json:"gambar"`  // Field for images
	UserId                int             `json:"user_id"` // Add UserId field
	SektorUsaha           string          `json:"sekto_usaha"`
	StatusTempatUsaha     string          `json:"status_tempat_usaha"`
	KodeProv              string          `json:"kode_prov"`
	KodeKabupaten string `gorm:"column:kode_kabupaten"`
	KodeKec               string          `json:"kode_kec"`
	KodeKel               string          `json:"kode_kelurahan"`
	Rt                    string          `json:"rt"`
	Rw                    string          `json:"rw"`
	KodePos               string          `json:"kode_pos"`
	NoNpwd                string          `json:"no_npwd"`
	BahanBakar            string          `json:"bahan_bakar"`
	TanggalMulaiUsaha     time.Time       `json:"tanggal_mulai_usaha"`
	Kapasitas             string             `json:"kapasitas"`
	TenagaKerjaPria       int             `json:"tenaga_kerja_pria"`
	TenagaKerjaWanita     int             `json:"tenaga_kerja_wanita"`
	NominalAset           decimal.Decimal `json:"nominal_aset"`
	NominalSendiri        decimal.Decimal `json:"nominal_sendiri"`
	EkonomiKreatif        bool            `json:"ekonomi_kreatif"`
	KriteriaUsaha         string          `json:"kriteria_usaha"`
	BentukUsaha           string          `json:"bentuk_usaha"`
	NoNib                 string          `json:"no_nib"`
	JenisUsaha            string          `json:"jenis_usaha"`
	Gaji 				  decimal.Decimal `json:"gaji"`
	KaryawanPria          int 			  `json:"karyawan_pria"`
	KaryawanWanita        int 			  `json:"karyawan_wanita"`
	Omset                 []OmsetRequest  `json:"omset"`
	Dokumen               []DokumenLegal  `json:"dokumen"`
}

type OmsetRequest struct {
	Bulan       string          `json:"bulan"`
	JumlahOmset decimal.Decimal `json:"jumlah_omset"`
}

type DokumenLegal struct {
	DokumenId     int             `validate:"required" json:"dokumen_id"`
	DokumenUpload json.RawMessage `json:"dok_upload"`
}



type Updateumkm struct {
	Name                  string          `json:"name"`
	NoNpwp                string          `json:"no_npwp"`
	Kategori_Umkm_Id      json.RawMessage `json:"kategori_umkm_id"`
	Nama_Penanggung_Jawab string          `json:"nama_penanggung_jawab"`
	Informasi_JamBuka     json.RawMessage `json:"informasi_jambuka"`
	No_Kontak             string          `json:"no_kontak"`
	Lokasi                string          `json:"lokasi"`
	Deskripsi             string          `json:"deskripsi"`
	Maps                  json.RawMessage `json:"maps"`
	Gambar                json.RawMessage `json:"gambar"`
	SektorUsaha           string          `json:"sekto_usaha"`
	StatusTempatUsaha     string          `json:"status_tempat_usaha"`
	KodeProv              string          `json:"kode_prov"`
	KodeKabupaten string `gorm:"column:kode_kabupaten"`
	KodeKec               string          `json:"kode_kec"`
	KodeKel               string          `json:"kode_kelurahan"`
	Rt                    string          `json:"rt"`
	Rw                    string          `json:"rw"`
	KodePos               string          `json:"kode_pos"`
	NoNpwd                string          `json:"no_npwd"`
	BahanBakar            string          `json:"bahan_bakar"`
	TanggalMulaiUsaha     time.Time       `json:"tanggal_mulai_usaha"`
	Kapasitas             string             `json:"kapasitas"`
	TenagaKerjaPria       int             `json:"tenaga_kerja_pria"`
	TenagaKerjaWanita     int             `json:"tenaga_kerja_wanita"`
	NominalAset           decimal.Decimal `json:"nominal_aset"`
	NominalSendiri        decimal.Decimal `json:"nominal_sendiri"`
	EkonomiKreatif        bool            `json:"ekonomi_kreatif"`
	KriteriaUsaha         string          `json:"kriteria_usaha"`
	BentukUsaha           string          `json:"bentuk_usaha"`
	NoNib                 string          `json:"no_nib"`
	Gaji 				  decimal.Decimal `json:"gaji"`
	KaryawanPria          int 			  `json:"karyawan_pria"`
	KaryawanWanita        int 			  `json:"karyawan_wanita"`
}

type UpdateActiveUmkm struct {
	Active int `json:"active"`
}
