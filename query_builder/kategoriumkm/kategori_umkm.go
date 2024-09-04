package query_builder_kategori_umkm

import (
	querybuilder "umkm/query_builder"
	"gorm.io/gorm"
)

type KategoriUmkmQueryBuilder interface {
	querybuilder.BaseQueryBuilderList
	GetBuilder(filters string, limit int, page int) (*gorm.DB, error)
}

type KategoriUmkmQueryBuilderImpl struct {
	querybuilder.BaseQueryBuilderList
	db *gorm.DB
}

func NewKategoriUmkmQueryBuilder(db *gorm.DB) *KategoriUmkmQueryBuilderImpl {
	return &KategoriUmkmQueryBuilderImpl{
		BaseQueryBuilderList: querybuilder.NewBaseQueryBuilderList(db),
		db: db,
	}
}

func (transaksiQueryBuilder *KategoriUmkmQueryBuilderImpl) GetBuilder(filters string, limit int, page int) (*gorm.DB, error) {
	query := transaksiQueryBuilder.db

	// Implementasi filter di sini
	if filters != "" {
		searchPattern := "%" + filters 
		query = query.Where("name ILIKE ?", searchPattern)
	}

	query, err := transaksiQueryBuilder.GetQueryBuilderList(query, filters, limit, page)
	if err != nil {
		return nil, err
	}

	// query = query.Debug().Preload("kategori_umkm")

	return query, nil
}
