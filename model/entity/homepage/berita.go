package entity

import (
	"time"
	"umkm/model/domain"
)

type BeritaFilterEntity struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Image   string `json:"image"`
	Content string `json:"content"`
	Author  string `json:"author"` // Ubah dari int ke string untuk menyimpan nama
	Created_at time.Time `json:"created_at"`
}

func ToBeritaFilterEntity(berita domain.Berita) BeritaFilterEntity {
	return BeritaFilterEntity{
		Id:      berita.Id,
		Title:   berita.Title,
		Image:   berita.Image,
		Content: berita.Content,
		Author:  berita.User.Username, // Ambil nama dari relasi User
		Created_at: berita.CreatedAt,
	}
}

func ToberitafilterEntities(berita []domain.Berita) []BeritaFilterEntity {
	var beritaEntities []BeritaFilterEntity
	for _, berita := range berita {
		beritaEntities = append(beritaEntities, ToBeritaFilterEntity(berita))
	}
	return beritaEntities
}
