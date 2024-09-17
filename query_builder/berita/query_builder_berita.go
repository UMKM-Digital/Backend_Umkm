package query_builder_berita

import (
	querybuilder "umkm/query_builder"

	"gorm.io/gorm"
)

type BeritaQueryBuilder interface {
	querybuilder.BaseQueryBuilderList
	GetBuilder( limit int, page int) (*gorm.DB, error)
}

type BeritaQueryBuilderImpl struct {
	querybuilder.BaseQueryBuilderList
	db *gorm.DB
}

func NewBeritaQueryBuilder(db *gorm.DB) *BeritaQueryBuilderImpl {
	return &BeritaQueryBuilderImpl{
		BaseQueryBuilderList: querybuilder.NewBaseQueryBuilderList(db),
		db:                   db,
	}
}

func (BeritaQueryBuilder *BeritaQueryBuilderImpl) GetBuilder( limit int, page int) (*gorm.DB, error) {
	query := BeritaQueryBuilder.db

	// Set default limit jika limit == 0
	if limit <= 0 {
		limit = 15
	}

	// Hitung offset
	offset := (page - 1) * limit
	query = query.Offset(offset).Limit(limit)

	// Apply pagination

	return query, nil
}
