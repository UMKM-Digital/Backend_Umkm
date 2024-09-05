package query_builder_umkm

import (
	querybuilder "umkm/query_builder"
	"gorm.io/gorm"
)

type UmkmQueryBuilder interface {
	querybuilder.BaseQueryBuilderList
	GetBuilder(filters string,limit int, page int) (*gorm.DB, error)
}

type UmkmQueryBuilderImpl struct {
	querybuilder.BaseQueryBuilderList
	db *gorm.DB
}

func NewUmkmQueryBuilder(db *gorm.DB) *UmkmQueryBuilderImpl {
	return &UmkmQueryBuilderImpl{
		BaseQueryBuilderList: querybuilder.NewBaseQueryBuilderList(db),
		db: db,
	}
}

func (umkmQueryBuilder *UmkmQueryBuilderImpl) GetBuilder( filters string,  limit int, page int,) (*gorm.DB, error) {
	query := umkmQueryBuilder.db

	if filters != "" {
		searchPattern := "%" + filters 
		query = query.Where("name ILIKE ?", searchPattern)
	}

	query, err := umkmQueryBuilder.GetQueryBuilderList(query, filters, limit, page,)
	if err != nil {
		return nil, err
	}

	query = query.Preload("HakAkses")

	return query, nil
}
