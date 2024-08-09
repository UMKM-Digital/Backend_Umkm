package entity

import (
	"umkm/model/domain"

	"github.com/google/uuid"
)

type UmkmEntity struct {
	IdUmkm uuid.UUID `json:"id"`
	Name string `json:"name"`
	Images domain.JSONB `json:"gambar"`
	Lokasi string `json:"lokasi"`
	KategoriUmkm domain.JSONB `json:"kategori_umkm_id"`
}

func ToUmkmEntity(umkm domain.UMKM) UmkmEntity {
	return UmkmEntity{
	 IdUmkm: umkm.IdUmkm,
	 Name: umkm.Name,
	 Images: umkm.Images,
	 Lokasi: umkm.Lokasi,
	 KategoriUmkm: umkm.KategoriUmkmId,
	}
}

// func ToUmkmEntities(UmkmList []domain.Transaksi) []UmkmEntity {
//     var UmkmEntities []UmkmEntity
//     for _, UmkmList := range UmkmList {
//         UmkmEntities = append(UmkmEntities, ToUmkmEntity(UmkmList))
//     }
//     return UmkmEntities
// }

func ToUmkmEntities(transaksiList []domain.UMKM) []UmkmEntity {
    var umkmEntities []UmkmEntity
    for _, umkm := range transaksiList {
        umkmEntities = append(umkmEntities, ToUmkmEntity(umkm))
    }
    return umkmEntities
}

