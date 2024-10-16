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
//bentukusaha
func (repo *SektorUSahaRepoImpl) GetBentukUsaha() ([]domain.BentukUsaha, error) {
	var bentukusaha []domain.BentukUsaha
	err := repo.db.Find(&bentukusaha).Error
	if err != nil {
		return nil, err
	}
	return bentukusaha, nil
}
//statustempatusaha
func (repo *SektorUSahaRepoImpl) GetStatusTempatUsaha() ([]domain.StatusTempatUsaha, error) {
	var statustempatusaha []domain.StatusTempatUsaha
	err := repo.db.Find(&statustempatusaha).Error
	if err != nil {
		return nil, err
	}
	return statustempatusaha, nil
}