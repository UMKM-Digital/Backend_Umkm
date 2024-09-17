package general_query_builder

// import (
// 	querybuilder "umkm/query_builder"
// 	"gorm.io/gorm"
// )

// type TransaksiQueryBuilder interface {
// 	querybuilder.BaseQueryBuilderList
// 	GetBuilder(filters string, limit int, page int, status string) (*gorm.DB, error)
// }

// type TransaksiQueryBuilderImpl struct {
// 	querybuilder.BaseQueryBuilderList
// 	db *gorm.DB
// }

// func NewTransaksiQueryBuilder(db *gorm.DB) *TransaksiQueryBuilderImpl {
// 	return &TransaksiQueryBuilderImpl{
// 		BaseQueryBuilderList: querybuilder.NewBaseQueryBuilderList(db),
// 		db: db,
// 	}
// }

// func (transaksiQueryBuilder *TransaksiQueryBuilderImpl) GetBuilder(filters string, limit int, page int, status string) (*gorm.DB, error) {
// 	query := transaksiQueryBuilder.db

// 	// Implementasi filter di sini
// 	if filters != "" {
// 		searchPattern := "%" + filters + "%"
// 		query = query.Where("no_invoice ILIKE ? OR name_client ILIKE ?", searchPattern, searchPattern)
// 	}

// 	if status != "" { 
// 		query = query.Where("status = ?", status)
// 	}

// 	query, err := transaksiQueryBuilder.GetQueryBuilderList(query, filters, limit, page)
// 	if err != nil {
// 		return nil, err
// 	}

// 	query = query.Preload("Umkm")

// 	return query, nil
// }
