package querybuildertransaksi

import (
	querybuilder "umkm/query_builder"

	"gorm.io/gorm"
)

type TransaksiQueryBuilder interface {
	querybuilder.BaseQueryBuilder
	GetBuilder(filter map[string]string)(*gorm.DB)
	GetAllowedFilters()[]string
}


type TransaksiQueryBuilderImpl struct {
	querybuilder.BaseQueryBuilder
}

func NewTransaksiQueryBuilder(db *gorm.DB) *TransaksiQueryBuilderImpl {
	return &TransaksiQueryBuilderImpl{
		BaseQueryBuilder: querybuilder.NewBaseQueryBuilder(db),
	}
}

func (transaksiQueryBuilder *TransaksiQueryBuilderImpl) GetBuilder(filters map[string]string, sort string, limit int, page int) (*gorm.DB, error) {
	allowedFilters := transaksiQueryBuilder.getAllowedFilters()
	query, err := transaksiQueryBuilder.GetQueryBuilder(filters, allowedFilters)
	if err != nil {
		return nil, err
	}
	query = query.Preload("Umkm")

	return query, nil
}

func (transaksiQueryBuilder *TransaksiQueryBuilderImpl) getAllowedFilters() []string {
	return []string{
		"tanggal",
	}
}
