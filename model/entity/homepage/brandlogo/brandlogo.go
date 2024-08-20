package entity

import (
	"time"
	domain "umkm/model/domain/homepage"
)

type BrandLogoEntity struct {
	BrandName string    `json:"brand_name"`
	BrandLogo string    `json:"brand_logo"`
	Created   time.Time `json:"created_at"`
}

// Mengonversi satu instance dari domain.Brandlogo ke BrandLogoEntity
func ToBrandEntity(brandlogo domain.Brandlogo) BrandLogoEntity {
	return BrandLogoEntity{
		BrandName: brandlogo.BrandName,
		BrandLogo: brandlogo.BrandLogo,
		Created:   brandlogo.Created_at,
	}
}

// Mengonversi slice dari domain.Brandlogo ke slice dari BrandLogoEntity
func ToBrandLogoEntities(brandlogos []domain.Brandlogo) []BrandLogoEntity {
	var brandLogoEntities []BrandLogoEntity
	for _, brandlogo := range brandlogos {
		brandLogoEntities = append(brandLogoEntities, ToBrandEntity(brandlogo))
	}
	return brandLogoEntities
}
