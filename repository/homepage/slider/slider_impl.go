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
	err := repo.db.Order("id ASC").Find(&slider).Error
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

func (repo *SliderRepoImpl) DelSlider(id int) error{
	if err := repo.db.Delete(&domain.Slider{}, id).Error; err != nil {
		return err
	}
	return nil
}

func(repo *SliderRepoImpl) UpdateSliderId(id int, slider domain.Slider) (domain.Slider, error){
	var existingSlider domain.Slider
    if err := repo.db.First(&existingSlider, id).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return domain.Slider{}, errors.New("testimonial not found")
        }
        return domain.Slider{}, err
    }

    // Lakukan pembaruan
    if err := repo.db.Model(&existingSlider).Updates(slider).Error; err != nil {
        return domain.Slider{}, errors.New("failed to update slider")
    }

    return slider, nil
}

// TestimonalRepoImpl adalah implementasi dari repository untuk testimonial
func (repo *SliderRepoImpl) UpdateActiveId(idSlider int, active int) error {
    if err := repo.db.Model(&domain.Slider{}).Where("id = ?", idSlider).Update("active", active).Error; err != nil {
        return errors.New("failed to update slider active status")
    }
    return nil
}

func (repo *SliderRepoImpl) GetSliderActive(active int)([]domain.Slider, error) {
    var slider []domain.Slider
    err := repo.db.Raw("SELECT * FROM homepage.sleder WHERE active = ?", active).Scan(&slider).Error
    if err != nil {
        return nil, err
    }
    return slider, nil
}

