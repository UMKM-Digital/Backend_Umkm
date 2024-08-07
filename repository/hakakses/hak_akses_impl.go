package hakaksesrepo

import (
	"umkm/model/domain"

	"gorm.io/gorm"
)

type HakAksesRepoUmkmImpl struct {
	db *gorm.DB
}

func NewHakAksesRepositoryImpl(db *gorm.DB) *HakAksesRepoUmkmImpl{
	return &HakAksesRepoUmkmImpl{db:db}
}


func (repo *HakAksesRepoUmkmImpl) CreateHakAkses(hakAkses *domain.HakAkses) error {
	return repo.db.Create(hakAkses).Error
}
