package masterdokumenlegalrepo

import (
	"umkm/model/domain"

	"gorm.io/gorm"
)

type MasterDokumenLegalRepoImpl struct {
	db *gorm.DB
}

func NewDokumenLegalRepoImpl(db *gorm.DB) *MasterDokumenLegalRepoImpl {
	return &MasterDokumenLegalRepoImpl{
		db:                 db,
	}
}

func (repo *MasterDokumenLegalRepoImpl) Created(dokumen domain.MasterDokumenLegal) (domain.MasterDokumenLegal, error) {
	err := repo.db.Create(&dokumen).Error
	if err != nil {
		return domain.MasterDokumenLegal{}, err
	}

	return dokumen, nil
}