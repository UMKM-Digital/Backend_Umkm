package testimonialrepo

import (
	domain "umkm/model/domain/homepage"

	"errors"

	"gorm.io/gorm"
)

type TestimonalRepoImpl struct {
	db *gorm.DB
}

func NewTestimonal(db *gorm.DB) *TestimonalRepoImpl {
	return &TestimonalRepoImpl{db: db}
}

func (repo *TestimonalRepoImpl) CreateTestimonial(testimonal domain.Testimonal) (domain.Testimonal, error) {
	err := repo.db.Create(&testimonal).Error
	if err != nil {
		return domain.Testimonal{}, err
	}
	return testimonal, nil
}

func (repo *TestimonalRepoImpl) GetTestimonial() ([]domain.Testimonal, error) {
	var testimonal []domain.Testimonal
	err := repo.db.Order("id ASC").Find(&testimonal).Error
	if err != nil {
		return nil, err
	}
	return testimonal, nil
}

// delete
func (repo *TestimonalRepoImpl) DelTestimonial(id int) error {
	if err := repo.db.Delete(&domain.Testimonal{}, id).Error; err != nil {
		return err
	}
	return nil
}

// id
func (repo *TestimonalRepoImpl) GetTransaksiByid(id int) (domain.Testimonal, error) {
    var transaksidata domain.Testimonal

    // Gunakan `First` untuk mendapatkan satu entri berdasarkan ID
    err := repo.db.First(&transaksidata, "id = ?", id).Error

    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return domain.Testimonal{}, errors.New("testimonial not found")
        }
        return domain.Testimonal{}, err
    }

    return transaksidata, nil
}

func (repo *TestimonalRepoImpl) UpdateTestimonialId(id int, testimonial domain.Testimonal) (domain.Testimonal, error) {
    // Periksa apakah testimonial dengan ID ini ada
    var existingTestimonial domain.Testimonal
    if err := repo.db.First(&existingTestimonial, id).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return domain.Testimonal{}, errors.New("testimonial not found")
        }
        return domain.Testimonal{}, err
    }

    // Lakukan pembaruan
    if err := repo.db.Model(&existingTestimonial).Updates(testimonial).Error; err != nil {
        return domain.Testimonal{}, errors.New("failed to update testimonial")
    }

    return testimonial, nil
}

func (repo *TestimonalRepoImpl) GetTestimonialActive(active int)([]domain.Testimonal, error) {
    var testimonal []domain.Testimonal
    err := repo.db.Raw("SELECT * FROM homepage.testimonial WHERE active = ?", active).Scan(&testimonal).Error
    if err != nil {
        return nil, err
    }
    return testimonal, nil
}

// TestimonalRepoImpl adalah implementasi dari repository untuk testimonial
func (repo *TestimonalRepoImpl) UpdateActiveId(idTestimonial int, active int) error {
    if err := repo.db.Model(&domain.Testimonal{}).Where("id = ?", idTestimonial).Update("active", active).Error; err != nil {
        return errors.New("failed to update testimonial active status")
    }
    return nil
}

