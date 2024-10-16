package entity

import "umkm/model/domain"

type KategoriEntity struct {
	Id   int    `json:"id"`
	IdSektorUsaha int `json:"id_sektor_usaha"`
	Name string `json:"name"`
}

func ToKategoriEntity(kategori domain.Kategori_Umkm) KategoriEntity {
	return KategoriEntity{
		Id: kategori.IdKategori,
		IdSektorUsaha: kategori.IdSektorUsaha,
		Name: kategori.Name,
	}
}

func ToKategoriEntities(kategoriList []domain.Kategori_Umkm) []KategoriEntity {
    var kategoriEntities []KategoriEntity
    for _, kategori := range kategoriList {
        kategoriEntities = append(kategoriEntities, ToKategoriEntity(kategori))
    }
    return kategoriEntities
}