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

func ToUmkmEntities(transaksiList []domain.UMKM) []UmkmEntity {
    var umkmEntities []UmkmEntity
    for _, umkm := range transaksiList {
        umkmEntities = append(umkmEntities, ToUmkmEntity(umkm))
    }
    return umkmEntities
}

type UmkmEntityList struct{
	Id uuid.UUID `json:"id"`
	Name string `json:"name"`
}

func ToUmkmEntitiyList(umkmList domain.UMKM) UmkmEntityList {
	return UmkmEntityList{
	 Id: umkmList.IdUmkm,
	 Name: umkmList.Name,
	}
}

func ToUmkmListEntities(umkmList []domain.UMKM) []UmkmEntityList {
    var umkmListEntities []UmkmEntityList
    for _, umkm := range umkmList {
        umkmListEntities = append(umkmListEntities, ToUmkmEntitiyList(umkm))
    }
    return umkmListEntities
}