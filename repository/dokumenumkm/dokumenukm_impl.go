package dokumenumkmrepo

import (
	"umkm/model/domain"
	"gorm.io/gorm"
)

type UmkmDokumenImpl struct {
	db *gorm.DB
}

func NewDokumenRepositoryImpl(db *gorm.DB) *UmkmDokumenImpl {
	return &UmkmDokumenImpl{
		db: db,
	}
}

func (repo *UmkmDokumenImpl) CreateRequest(dokumenumkm domain.UmkmDokumen) (domain.UmkmDokumen, error) {
	err := repo.db.Create(&dokumenumkm).Error
	if err != nil {
		return domain.UmkmDokumen{}, err
	}

	return dokumenumkm, nil
}

