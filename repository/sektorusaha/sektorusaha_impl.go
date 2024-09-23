package sektorusaharepo

import (
	domain "umkm/model/domain/master"

	"gorm.io/gorm"
)

type SektorUSahaRepoImpl struct {
	db *gorm.DB
}

func NewSektorUsaha(db *gorm.DB) *SektorUSahaRepoImpl {
	return &SektorUSahaRepoImpl{db: db}
}

func (repo *SektorUSahaRepoImpl) CreateSektorUsaha(sektorusaha domain.SektorUsaha) (domain.SektorUsaha, error) {
	err := repo.db.Create(&sektorusaha).Error
	if err != nil {
		return domain.SektorUsaha{}, err
	}
	return sektorusaha, nil
}

func (repo *SektorUSahaRepoImpl) GetSektorUsaha() ([]domain.SektorUsaha, error) {
	var sektorusaha []domain.SektorUsaha
	err := repo.db.Find(&sektorusaha).Error
	if err != nil {
		return nil, err
	}
	return sektorusaha, nil
}