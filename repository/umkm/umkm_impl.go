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

func (repo *RepoUmkmImpl) GetUmkmListByIds(ctx context.Context, umkmIDs []uuid.UUID, filters string, limit int, page int) ([]domain.UMKM, int, int, int, *int, *int, error) {
    var umkm []domain.UMKM
    var totalcount int64

    // Set default limit jika limit == 0
    if limit <= 0 {
        limit = 15
    }

    // Dapatkan query dengan filter dan pagination
    query, err := repo.umkmQueryBuilder.GetBuilder(filters, limit, page)
    if err != nil {
        return nil, 0, 0, 0, nil, nil, err
    }

    // Filter berdasarkan umkmIDs
    err = query.Where("id IN (?)", umkmIDs).Find(&umkm).Error
    if err != nil {
        return nil, 0, 0, 0, nil, nil, err
    }

    // Hitung total records dari hasil pencarian, tanpa pagination
    totalQuery, err := repo.umkmQueryBuilder.GetBuilder(filters, 0, 0) // Tanpa pagination
    if err != nil {
        return nil, 0, 0, 0, nil, nil, err
    }

    // Hitung jumlah total records
    err = totalQuery.Model(&domain.UMKM{}).Where("id IN (?)", umkmIDs).Count(&totalcount).Error
    if err != nil {
        return nil, 0, 0, 0, nil, nil, err
    }

    // Hitung total pages
    totalPages := 1
    if limit > 0 {
        totalPages = int((totalcount + int64(limit) - 1) / int64(limit))
    }

    // Jika page > totalPages, return kosong
    if page > totalPages {
        return nil, int(totalcount), page, totalPages, nil, nil, nil
    }

    currentPage := page

    // Tentukan nextPage dan prevPage
    var nextPage *int
    if currentPage < totalPages {
        np := currentPage + 1
        nextPage = &np
    }

    var prevPage *int
    if currentPage > 1 {
        pp := currentPage - 1
        prevPage = &pp
    }

    return umkm, int(totalcount), currentPage, totalPages, nextPage, prevPage, nil
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