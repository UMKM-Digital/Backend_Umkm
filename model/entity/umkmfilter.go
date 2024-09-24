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
	PenanggunagJawab string `json:"nama_penanggung_jawab"`
	KategoriUmkm domain.JSONB `json:"kategori_umkm_id"`
	 TotalProduk        int           `json:"total_produk"`
}

func ToUmkmFilterEntity(umkm domain.UMKM, products []domain.Produk) UmkmFilterEntity {
	totalProduk := CalculateTotalProdukByUmkm(umkm.IdUmkm, products)
	return UmkmFilterEntity{
	Id: umkm.IdUmkm,
	Name: umkm.Name,
	Gambar: umkm.Images,
	Lokasi: umkm.Lokasi,
	PenanggunagJawab: umkm.NamaPenanggungJawab,
	KategoriUmkm: umkm.KategoriUmkmId,
	TotalProduk:        totalProduk, // Menambahkan total produk
	}
}

func ToUmkmfilterEntities(umkmList []domain.UMKM, products []domain.Produk) []UmkmFilterEntity {
    var umkmEntities []UmkmFilterEntity
    for _, umkm := range umkmList {
        umkmEntity := ToUmkmFilterEntity(umkm, products)
        umkmEntities = append(umkmEntities, umkmEntity)
    }
    return umkmEntities
}


//menghitung produk
func CalculateTotalProdukByUmkm(umkmId uuid.UUID, products []domain.Produk) int {
    total := 0
    for _, produk := range products {
        if produk.UmkmId == umkmId {
            total++
        }
    }
    return total
}

