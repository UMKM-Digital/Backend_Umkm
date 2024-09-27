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
		searchPattern := "%" + filters + "%"
		query = query.Where("nama ILIKE ?", searchPattern)
	}

	if limit <= 0 {
        limit = 15
    }

    // Hitung offset
    offset := (page - 1) * limit
    query = query.Offset(offset).Limit(limit)

	query, err := transaksiQueryBuilder.GetQueryBuilderList(query, limit, page)
	if err != nil {
		return nil, err
	}

	return query, nil
}