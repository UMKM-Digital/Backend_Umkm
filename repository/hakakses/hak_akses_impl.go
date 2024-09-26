package hakaksesrepo

import (
	"context"
	"umkm/model/domain"

	"github.com/google/uuid"
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

func (repo *HakAksesRepoUmkmImpl) GetHakAksesByUserId(ctx context.Context, userId int) ([]domain.HakAkses, error){
	var hakAkses []domain.HakAkses
	err := repo.db.Where("user_id = ?", userId).Find(&hakAkses).Error
	if err != nil {
		return nil, err
	}
	return hakAkses, nil
}

func(repo *HakAksesRepoUmkmImpl)DeleteUmkmId(id uuid.UUID) error{
	return repo.db.Where("umkm_id = ?", id).Delete(&domain.HakAkses{}).Error
}