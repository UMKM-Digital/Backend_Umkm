package query_builder_kategori_produk

import (
	querybuilder "umkm/query_builder"

	"gorm.io/gorm"
)

type KategoriProdukQueryBuilder interface {
	querybuilder.BaseQueryBuilderList
	GetBuilder(filters string, limit int, page int) (*gorm.DB, error)
}

type KategoriProdukQueryBuilderImpl struct {
	querybuilder.BaseQueryBuilderList
	db *gorm.DB
}

func NewKategoriProdukQueryBuilder(db *gorm.DB) *KategoriProdukQueryBuilderImpl {
	return &KategoriProdukQueryBuilderImpl{
		BaseQueryBuilderList: querybuilder.NewBaseQueryBuilderList(db),
		db:                   db,
	}
}

func (transaksiQueryBuilder *KategoriProdukQueryBuilderImpl) GetBuilder(filters string, limit int, page int) (*gorm.DB, error) {
	query := transaksiQueryBuilder.db

	// Implementasi filter di sini
	if filters != "" {
		searchPattern := "%" + filters
		query = query.Where("nama ILIKE ?", searchPattern)
	}

	query, err := transaksiQueryBuilder.GetQueryBuilderList(query, filters, limit, page)
	if err != nil {
		return nil, err
	}

	query = query.Preload("UMKM")

	return query, nil
}
