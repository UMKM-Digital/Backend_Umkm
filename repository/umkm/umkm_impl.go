package umkmrepo

import (
	"context"
	"umkm/model/domain"
	query_builder_umkm "umkm/query_builder/umkm"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RepoUmkmImpl struct {
	db *gorm.DB
	umkmQueryBuilder query_builder_umkm.UmkmQueryBuilder
}

func NewUmkmRepositoryImpl(db *gorm.DB, umkmQueryBuilder query_builder_umkm.UmkmQueryBuilder) *RepoUmkmImpl {
	return &RepoUmkmImpl{
		db: db,
		umkmQueryBuilder: umkmQueryBuilder,
	}
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

func (repo *RepoUmkmImpl) GetUmkmListByIds(ctx context.Context, umkmIDs []uuid.UUID, filters string, limit int, page int) ([]domain.UMKM, int, error) {
	var umkm []domain.UMKM
	var totalcount int64

	query, err := repo.umkmQueryBuilder.GetBuilder(filters, limit, page)
	if err != nil {
		return nil, 0, err
	}

	err = query.Where("id IN (?)", umkmIDs).Find(&umkm).Error
	if err != nil {
		return nil, 0, err
	}

	umkmQueryBuilder, err := repo.umkmQueryBuilder.GetBuilder(filters, 0, 0)
	if err != nil {
		return nil, 0, err
	}
	
	err = umkmQueryBuilder.Model(&domain.UMKM{}).Where("id IN (?)", umkmIDs).Count(&totalcount).Error
	if err != nil {	
		return nil, 0, err
	}

	return umkm, int(totalcount), nil
}

//
func(repo *RepoUmkmImpl) GetUmkmFilterName(ctx context.Context,  umkmIDs []uuid.UUID)([]domain.UMKM, error){
	var umkm []domain.UMKM
	err := repo.db.Where("id IN (?)", umkmIDs).Find(&umkm).Error
	if err != nil {
		return []domain.UMKM{}, err
	}
	return umkm, nil
}

func(repo *RepoUmkmImpl) GetUmkmListWeb(ctx context.Context, umkmIds []uuid.UUID)([]domain.UMKM, error){
	var umkm []domain.UMKM

	err := repo.db.Where("id IN (?)", umkmIds).Find(&umkm).Error
	if err != nil{
		return []domain.UMKM{}, err
	}
	return umkm, nil
}