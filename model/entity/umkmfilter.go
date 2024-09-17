package entity

import (
	"umkm/model/domain"

	"github.com/google/uuid"
)

type UmkmFilterEntity struct {
	Id uuid.UUID `json:"id"`
	Name string `json:"name"`
	Gambar domain.JSONB `json:"gambar"`
	Lokasi string `json:"lokasi"`
	// KategoriUmkm domain.JSONB `json:"kategori_umkm_id"`
}

func ToUmkmFilterEntity(umkm domain.UMKM) UmkmFilterEntity {
	return UmkmFilterEntity{
	Id: umkm.IdUmkm,
	Name: umkm.Name,
	Gambar: umkm.Images,
	Lokasi: umkm.Lokasi,
	// KategoriUmkm: umkm.KategoriUmkmId,
	}
}

func ToUmkmfilterEntities(umkm []domain.UMKM) []UmkmFilterEntity {
    var umkmEntities []UmkmFilterEntity
    for _, umkm := range umkm {
        umkmEntities = append(umkmEntities, ToUmkmFilterEntity(umkm))
    }
    return umkmEntities
}
