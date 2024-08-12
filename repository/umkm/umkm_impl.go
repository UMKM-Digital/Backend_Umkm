package umkmrepo

import (
	"context"
	"umkm/model/domain"

	"github.com/google/uuid"
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

// func (repo *RepoUmkmImpl) GetUmkmList(ctx context.Context, umkmIDs []uuid.UUID)([]domain.UMKM, error){
// 	var umkm []domain.UMKM
// 	err := repo.db.Where("id IN (?))", umkmIDs).Find(&umkm).Error
// 	if err != nil{
// 		return []domain.UMKM{},err
// 	}
// 	return umkm, nil
// }
func (repo *RepoUmkmImpl) GetUmkmListByIds(ctx context.Context, umkmIDs []uuid.UUID) ([]domain.UMKM, error) {
	var umkm []domain.UMKM
	err := repo.db.Where("id IN (?)", umkmIDs).Find(&umkm).Error
	if err != nil {
		return []domain.UMKM{}, err
	}
	return umkm, nil
}