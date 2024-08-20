package brandrepo

import (
	domain "umkm/model/domain/homepage"

	"gorm.io/gorm"
)

type BrandLogoRepoImpl struct {
	db *gorm.DB
}

func NewBrandlogo(db *gorm.DB) *BrandLogoRepoImpl {
	return &BrandLogoRepoImpl{db: db}
}

func (repo *BrandLogoRepoImpl) CreatedBrandLogo(brandlogo domain.Brandlogo) (domain.Brandlogo, error) {
	err := repo.db.Create(&brandlogo).Error
	if err != nil {
		return domain.Brandlogo{}, err
	}
	return brandlogo, nil
}

func (repo *BrandLogoRepoImpl) GetBrandLogo() ([]domain.Brandlogo, error) {
    var brandlogo []domain.Brandlogo
    err := repo.db.Find(&brandlogo).Error
    if err != nil {
        return nil, err
    }
    return brandlogo, nil
}