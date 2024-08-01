package umkmrepo

import (
	"umkm/model/domain"

	"gorm.io/gorm"
)

type RepoUmkmImpl struct {
	db *gorm.DB
}

func NewUmkmRepositoryImpl(db *gorm.DB) *RepoUmkmImpl {
	return &RepoUmkmImpl{db: db}
}

func (repo *RepoUmkmImpl) CreateRequest(umkm domain.UMKM)(domain.UMKM, error) {
	err := repo.db.Create(&umkm).Error
	if err != nil {
		return domain.UMKM{}, err
	}

	return umkm, nil
}