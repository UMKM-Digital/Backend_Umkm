package sliderrepo

import (
	"errors"
	domain "umkm/model/domain/homepage"

	"gorm.io/gorm"
)

type SliderRepoImpl struct {
	db *gorm.DB
}

func NewSlider(db *gorm.DB) *SliderRepoImpl{
	return &SliderRepoImpl{db: db}
}

func (repo *SliderRepoImpl) Created(slider domain.Slider) (domain.Slider, error){
	err := repo.db.Create(&slider).Error
	if err != nil {
		return domain.Slider{}, err
	}
	return slider, nil
}

func (repo *SliderRepoImpl) GetSlider() ([]domain.Slider, error){
	var slider []domain.Slider
	err := repo.db.Find(&slider).Error
	if err != nil {
		return nil, err
	}

	return slider, nil
}

func (repo *SliderRepoImpl) GetSliderId(id int) (domain.Slider, error){
	var slide domain.Slider

    // Gunakan `First` untuk mendapatkan satu entri berdasarkan ID
    err := repo.db.First(&slide, "id = ?", id).Error

    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return domain.Slider{}, errors.New("slide not found")
        }
        return domain.Slider{}, err
    }

    return slide, nil
}