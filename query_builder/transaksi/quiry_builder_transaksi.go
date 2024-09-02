package general_query_builder

import (
	querybuilder "umkm/query_builder"

	
	"gorm.io/gorm"
)

type TransaksiQueryBuilder interface {
	querybuilder.BaseQueryBuilderList
	GetBuilder(filters string, limit int, page int, status int) (*gorm.DB, error)
}

type TransaksiQueryBuilderImpl struct {
	querybuilder.BaseQueryBuilderList
}

func NewTransaksiQueryBuilder(db *gorm.DB) *TransaksiQueryBuilderImpl {
	return &TransaksiQueryBuilderImpl{
		BaseQueryBuilderList: querybuilder.NewBaseQueryBuilderList(db),
	}
}

func (transaksiQueryBuilder *TransaksiQueryBuilderImpl) GetBuilder(filters string,  limit int, page int, status int) (*gorm.DB, error) {

	query, err := transaksiQueryBuilder.GetQueryBuilderList(filters,limit, page, status)
	if err != nil {
		return nil, err
	}
	query = query.Preload("Umkm")

	return query, nil
}
