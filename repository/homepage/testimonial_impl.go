package testimonialrepo

import (
	domain "umkm/model/domain/homepage"

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

func (repo *TestimonalRepoImpl) GetTestimonial()([]domain.Testimonal, error) {
    var testimonal []domain.Testimonal
    err := repo.db.Find(&testimonal).Error
    if err != nil {
        return nil, err
    }
    return testimonal, nil
}