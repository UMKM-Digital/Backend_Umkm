package entity

import (
	"umkm/model/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Struct UmkmEntity yang mencakup total produk
type UmkmEntity struct {
	IdUmkm              uuid.UUID    `json:"id"`
	Name                string       `json:"name"`
	Images              domain.JSONB `json:"gambar"`
	Lokasi              string       `json:"lokasi"`
	KategoriUmkm        domain.JSONB `json:"kategori_umkm_id"`
	NamaPenanggungJawab string       `json:"nama_penanggung_jawab"`
	TotalProduk         int          `json:"total_produk"`
}

// Fungsi untuk menghitung jumlah produk berdasarkan umkm_id
func CountProductsByUmkm(db *gorm.DB, umkmID uuid.UUID) (int, error) {
	var count int64
	err := db.Model(&domain.Produk{}).Where("umkm_id = ?", umkmID).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

// Fungsi untuk mengonversi domain.UMKM ke UmkmEntity termasuk menghitung total produk
func ToUmkmEntity(umkm domain.UMKM, db *gorm.DB) (UmkmEntity, error) {
	totalProduk, err := CountProductsByUmkm(db, umkm.IdUmkm)
	if err != nil {
		return UmkmEntity{}, err
	}

	return UmkmEntity{
		IdUmkm:              umkm.IdUmkm,
		Name:                umkm.Name,
		Images:              umkm.Images,
		Lokasi:              umkm.Lokasi,
		KategoriUmkm:        umkm.KategoriUmkmId,
		NamaPenanggungJawab: umkm.NamaPenanggungJawab,
		TotalProduk:         totalProduk,
	}, nil
}

// Fungsi untuk mengonversi daftar domain.UMKM menjadi daftar UmkmEntity termasuk total produk
func ToUmkmEntities(umkmList []domain.UMKM, db *gorm.DB) ([]UmkmEntity, error) {
	var umkmEntities []UmkmEntity
	for _, umkm := range umkmList {
		umkmEntity, err := ToUmkmEntity(umkm, db)
		if err != nil {
			return nil, err
		}
		umkmEntities = append(umkmEntities, umkmEntity)
	}
	return umkmEntities, nil
}

// Struct untuk UmkmEntityList (versi list singkat tanpa detail total produk)
type UmkmEntityList struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// Fungsi untuk mengonversi domain.UMKM ke UmkmEntityList (list sederhana)
func ToUmkmEntityList(umkm domain.UMKM) UmkmEntityList {
	return UmkmEntityList{
		Id:   umkm.IdUmkm,
		Name: umkm.Name,
	}
}

// Fungsi untuk mengonversi daftar domain.UMKM ke daftar UmkmEntityList (versi sederhana)
func ToUmkmListEntities(umkmList []domain.UMKM) []UmkmEntityList {
	var umkmListEntities []UmkmEntityList
	for _, umkm := range umkmList {
		umkmListEntities = append(umkmListEntities, ToUmkmEntityList(umkm))
	}
	return umkmListEntities
}
