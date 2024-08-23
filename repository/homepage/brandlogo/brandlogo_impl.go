package brandrepo

import (
	domain "umkm/model/domain/homepage"

	"gorm.io/gorm"
	"errors"
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

//delete produk
func (repo *BrandLogoRepoImpl) DeleteLogoId(id int) error {
    if err := repo.db.Delete(&domain.Brandlogo{}, id).Error; err != nil {
        return err
    }
    return nil
}

//
func (repo *BrandLogoRepoImpl) FindById(id int) (domain.Brandlogo, error) {
	var produk domain.Brandlogo
	if err := repo.db.First(&produk, "id = ?", id).Error; err != nil {
		return produk, err
	}
	return produk, nil
}

func (repo *BrandLogoRepoImpl)  UpdateBrandLogoId(id int, brandlogo domain.Brandlogo) (domain.Brandlogo, error) {
   
    var existingBrandLogo domain.Brandlogo
    if err := repo.db.First(&existingBrandLogo, id).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return domain.Brandlogo{}, errors.New("brandlogo not found")
        }
        return domain.Brandlogo{}, err
    }

    // Lakukan pembaruan
    if err := repo.db.Model(&existingBrandLogo).Updates(brandlogo).Error; err != nil {
        return domain.Brandlogo{}, errors.New("failed to update brandlogo")
    }

    return brandlogo, nil
}