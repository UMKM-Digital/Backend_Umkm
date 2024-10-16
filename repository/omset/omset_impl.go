package omsetrepo

import (
	"umkm/model/domain"

	"gorm.io/gorm"
)

type OmmsetRepoImpl struct {
	db                 *gorm.DB
	
}

func NewomsetRepositoryImpl(db *gorm.DB) *OmmsetRepoImpl {
	return &OmmsetRepoImpl{
		db:                 db,
	}
}

func (repo *OmmsetRepoImpl) CreateRequest(omset domain.Omset) (domain.Omset, error) {
	err := repo.db.Create(&omset).Error
	if err != nil {
		return domain.Omset{}, err
	}

	return omset, nil
}