package domain

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	// "github.com/shopspring/decimal"
)

type UMKM struct {
    IdUmkm               uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;column:id"`
    Name                 string    `gorm:"column:name"`
    NoNpwp               string    `gorm:"column:no_npwp"`
    KategoriUmkmId       JSONB     `gorm:"column:kategori_umkm_id"`
    NamaPenanggungJawab  string    `gorm:"column:nama_penanggung_jawab"`
    InformasiJambuka     JSONB     `gorm:"column:informasi_jambuka"`
    NoKontak             string    `gorm:"column:no_kontak"`
    Lokasi               string    `gorm:"column:lokasi"`
    Deskripsi               string    `gorm:"column:deskripsi"`
    Maps                 JSONB     `gorm:"column:maps"`
    Images                 JSONB     `gorm:"column:gambar"` // Menggunakan JSONB untuk menyimpan URL gambar // Menyimpan URL gambar
    CreatedAt            time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt            time.Time `gorm:"column:updated_at;autoUpdateTime"`
    SektorUsaha                 string `gorm:"column:sektor_usaha"`//
    StatusTempatUsaha                 string `gorm:"column:status_tempat_usaha"`
    Bentukusaha    string `gorm:"column:bentuk_usaha"`
    KodeProv                 string `gorm:"column:kode_prov"`
    KodeKabupaten                 string `gorm:"column:kode_kab"`
    KodeKecamatan                 string `gorm:"column:kode_kec"`
    KodeKelurahan                 string `gorm:"column:kode_kelurahan"`
    RT string `gorm:"column:rt"`
    Rw string `gorm:"column:rw"`
    KodePos string `gorm:"column:kode_pos"`
    NoNpwd string `gorm:"column:no_npwd"`
    BahanBakar string `gorm:"column:bahan_bakar"`
    TanggalMulaiUsaha time.Time `gorm:"column:tanggal_mulai_usaha"`
    Kapasitas string `gorm:"column:kapasitas"`
    TenagaKerjaPria int `gorm:"column:tenaga_kerja_pria"`
    TenagaKerjaWanita int `gorm:"column:tenaga_kerja_wanita"`
    NominalAset decimal.Decimal `gorm:"column:nominal_aset"`
    NominalSendiri decimal.Decimal `gorm:"column:nominal_sendiri"`
    EkonomiKreatif bool `gorm:"column:ekonomi_kreatif"`
    KriteriaUsaha string `gorm:"column:kriteria_usaha"`
    NoNib         string `gorm:"column:no_nib"`
    Active                 int `gorm:"column:active"`
    JenisUsaha string `gorm:"column:jenis_usaha"`
    BentukUsaha string `gorm:"column:bentuk_usaha"`
    HakAkses []HakAkses  `gorm:"foreignKey:UmkmId;references:IdUmkm"`
    Produk               []Produk  `gorm:"foreignkey:UmkmId"`
    Transaksi            []Transaksi `gorm:"foreignkey:UmkmId"`
    Dokumen            []UmkmDokumen `gorm:"foreignkey:UmkmId"`
    Omset               []Omset      `gorm:"foreignkey:UmkmId"`
}

func (UMKM) TableName() string {
    return "umkm"
}

// JSONB is a custom type to handle PostgreSQL JSONB columns in GORM
type JSONB map[string]interface{}

func (j JSONB) Value() (driver.Value, error) {
    return json.Marshal(j)
}

func (j *JSONB) Scan(value interface{}) error {
    if value == nil {
        *j = JSONB{}
        return nil
    }
    bytes, ok := value.([]byte)
    if !ok {
        return errors.New("failed to scan JSONB value")
    }
    return json.Unmarshal(bytes, j)
}

type JSON []byte
