package domain

import (
    "time"
    "database/sql/driver"
    "encoding/json"
    "errors"
    "github.com/google/uuid"
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
    Maps                 JSONB     `gorm:"column:maps"`
    Images                 JSONB     `gorm:"column:gambar"` // Menggunakan JSONB untuk menyimpan URL gambar // Menyimpan URL gambar
    CreatedAt            time.Time `gorm:"column:created_at"`
    UpdatedAt            time.Time `gorm:"column:updated_at"`
    HakAkses             []HakAkses `gorm:"foreignKey:umkm_id"`
    KategoriProduk        []KategoriProduk `gorm:"foreignkey:umkm_id"`
    Produk               []Produk  `gorm:"foreignkey:UmkmId"`
    Transaksi            []Transaksi `gorm:"foreignkey:UmkmId"`
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