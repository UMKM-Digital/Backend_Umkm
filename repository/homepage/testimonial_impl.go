package testimonialrepo

import (
	domain "umkm/model/domain/homepage"

	"gorm.io/gorm"
	"errors"
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

func (repo *TestimonalRepoImpl) GetTestimonial()([]domain.Testimonal, error) {
    var testimonal []domain.Testimonal
    err := repo.db.Find(&testimonal).Error
    if err != nil {
        return nil, err
    }
    return testimonal, nil
}

//delete
func (repo *TestimonalRepoImpl) DelTransaksi(id int)error{
	if err := repo.db.Delete(&domain.Testimonal{}, id).Error; err != nil {
        return err
    }
    return nil
}

//id
func(repo *TestimonalRepoImpl)  GetTransaksiByid(id int) (domain.Testimonal, error){
	var transaksidata domain.Testimonal

	err := repo.db.Find(&transaksidata, "id = ?", id).Error

	if err != nil {
		return domain.Testimonal{},errors.New("kategori tidak ditemukan")
	}

	return transaksidata, nil
}

func (repo *TestimonalRepoImpl) UpdateTransaksiId(id int, testimonal domain.Testimonal) (domain.Testimonal, error){
    if err := repo.db.Model(&domain.Testimonal{}).Where("id = ?", id).Updates(testimonal).Error; err != nil {
        return domain.Testimonal{}, errors.New("failed to update testimonial")
    }
    return testimonal, nil
}