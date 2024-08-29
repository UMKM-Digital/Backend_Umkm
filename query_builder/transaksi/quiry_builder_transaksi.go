package general_query_builder

import (
	"gorm.io/gorm"
	"umkm/query_builder"
)

type EventQueryBuilder interface {
	querybuilder.BaseQueryBuilderTransaksi
	GetQueryBuilder(filters map[string]string, limit string, page int) (*gorm.DB, error)
	getAllowedFilters() []string
}

type EventQueryBuilderImpl struct {
	*querybuilder.BaseQueryBuilderMain
}

func NewEventQueryBuilder(db *gorm.DB) *EventQueryBuilderImpl {
	return &EventQueryBuilderImpl{
		BaseQueryBuilderMain: querybuilder.NewBaseQueryBuilderMain(db),
	}
}

// Implementasi GetQueryBuilder sesuai dengan signature interface
func (impl *EventQueryBuilderImpl) GetQueryBuilder(filters map[string]string, limit string, page int) (*gorm.DB, error) {
	allowedFilters := impl.getAllowedFilters()
	return impl.BaseQueryBuilderMain.GetQueryBuilder(filters, limit, page, allowedFilters)
}

func (eventQueryBuilder *EventQueryBuilderImpl) getAllowedFilters() []string {
	return []string{
		"year",  // Filter berdasarkan tahun
		"month", // Filter berdasarkan bulan
		"day",   // Filter berdasarkan hari
	}
}
