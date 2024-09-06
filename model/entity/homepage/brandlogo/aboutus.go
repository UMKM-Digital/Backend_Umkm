package entity

import (
	"time"
	domain "umkm/model/domain/homepage"
)

type AboutUsEntity struct {
	ID int `json:"id"`
	Image string `json:"image"`
	Description string `json:"description"`
	Created time.Time `json:"created_at"`
	Update time.Time `json:"updated_at"`
}

func ToAboutUsEntity(about domain.AboutUs) AboutUsEntity {
	return AboutUsEntity{
		ID: about.Id,
		Image: about.Image,
		Description: about.Description,
		Created: about.CreatedAt,
		Update: about.UpdatedAt,
	}
}

func ToAboutEntities(about []domain.AboutUs) []AboutUsEntity {
	var aboutiEntities []AboutUsEntity
	for _, about := range about {
		aboutiEntities = append(aboutiEntities, ToAboutUsEntity(about))
	}
	return aboutiEntities
}