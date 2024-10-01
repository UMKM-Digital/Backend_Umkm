package aboutusrepo

import (
	"errors"
	domain "umkm/model/domain/homepage"
	"gorm.io/gorm"
)

type AboutUsRepoImpl struct {
	db *gorm.DB
}

func NewAboutUS(db *gorm.DB) *AboutUsRepoImpl {
	return &AboutUsRepoImpl{db: db}
}

func (repo *AboutUsRepoImpl) CreatedAboutUs(aboutUs domain.AboutUs) (domain.AboutUs, error) {
	err := repo.db.Create(&aboutUs).Error
	if err != nil {
		return domain.AboutUs{}, err
	}
	return aboutUs, nil
}

func (repo *AboutUsRepoImpl) GetAboutUs() ([]domain.AboutUs, error) {
	var aboutUs []domain.AboutUs
	err := repo.db.Order("id ASC").Find(&aboutUs).Error
	if err != nil {
		return nil, err
	}
	return aboutUs, nil
}

func (repo *AboutUsRepoImpl) FindByAboutId(id int) (domain.AboutUs, error) {
	var aboutUs domain.AboutUs
	if err := repo.db.First(&aboutUs, "id = ?", id).Error; err != nil {
		return aboutUs, err
	}
	return aboutUs, nil
}

func (repo *AboutUsRepoImpl) UpdateAboutUsId(id int, aboutUs domain.AboutUs) (domain.AboutUs, error) {
	var existingAboutUs domain.AboutUs
	if err := repo.db.First(&existingAboutUs, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.AboutUs{}, errors.New("about us not found")
		}
		return domain.AboutUs{}, err
	}

	if err := repo.db.Model(&existingAboutUs).Updates(aboutUs).Error; err != nil {
		return domain.AboutUs{}, errors.New("failed to update AboutUs")
	}

	return aboutUs, nil
}
