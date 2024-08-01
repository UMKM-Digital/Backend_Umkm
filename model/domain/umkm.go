package domain

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type UMKM struct {
    IdUmkm               uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;column:id"`
    Name                 string    `gorm:"column:name"`
    NoNpwp               string    `gorm:"column:no_npwp"`
    Gambar               JSONB     `gorm:"column:gambar"`
    KategoriUmkmId       JSONB     `gorm:"column:kategori_umkm_id"`
    NamaPenanggungJawab  string    `gorm:"column:nama_penanggung_jawab"`
    InformasiJambuka     JSONB     `gorm:"column:informasi_jambuka"`
    NoKontak             string    `gorm:"column:no_kontak"`
    Lokasi               string    `gorm:"column:lokasi"`
    Maps                 JSONB    `gorm:"column:maps"`
    Kategori_Umkms       []Kategori_Umkm
    CreatedAt            time.Time `gorm:"column:created_at"`
    UpdatedAt            time.Time `gorm:"column:updated_at"`
}

// JSONB is a custom type to handle PostgreSQL JSONB columns in GORM
type JSONB map[string]interface{}

func (j JSONB) Value() (driver.Value, error) {
    return json.Marshal(j)
}

func (j *JSONB) Scan(value interface{}) error {
    if err := json.Unmarshal(value.([]byte), &j); err != nil {
        return err
    }
    return nil
}

func (UMKM) TableName() string {
    return "umkm"
}