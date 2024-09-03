// package querybuilder

// import (


// 	"gorm.io/gorm"
// )

// type BaseQueryBuilderList interface {
// 	GetQueryBuilderList(filter string, limit int, page int, status int) (*gorm.DB, error)
// }

// type BaseQueryBuilderListImpl struct {
// 	db *gorm.DB
// }

// func NewBaseQueryBuilderList(db *gorm.DB) *BaseQueryBuilderListImpl {
// 	return &BaseQueryBuilderListImpl{
// 		db: db,
// 	}
// }

// func (baseQueryBuilder *BaseQueryBuilderListImpl) GetQueryBuilderList(filter string, limit int, page int, status int) (*gorm.DB, error) {
// 	query := baseQueryBuilder.db

// 	// Jika filter tidak kosong, tambahkan kondisi OR untuk mencari di dua kolom
// 	if filter != "" {
// 		searchPattern := "%" + filter + "%"
// 		query = query.Where("no_invoice ILIKE ? OR name_client ILIKE ?", searchPattern, searchPattern)
// 	}

// 	if status == 0 {
// 		query = query.Where("status = ?", status)
// 	}else{
// 		query = query.Where("status = ?", status)
// 	}

// 	// Set limit dan pagination
// 	if limit == 0 {
// 		limit = 15
// 	}
// 	query = query.Limit(limit)

// 	if page == 0 {
// 		page = 1
// 	}
// 	query = query.Offset((page - 1) * limit)

// 	return query, nil
// }


package querybuilder

import (
	"gorm.io/gorm"
)

type BaseQueryBuilderList interface {
	GetQueryBuilderList(query *gorm.DB, filter string, limit int, page int) (*gorm.DB, error)
}

type BaseQueryBuilderListImpl struct {
	db *gorm.DB
}

func NewBaseQueryBuilderList(db *gorm.DB) *BaseQueryBuilderListImpl {
	return &BaseQueryBuilderListImpl{
		db: db,
	}
}

func (baseQueryBuilder *BaseQueryBuilderListImpl) GetQueryBuilderList(query *gorm.DB, filter string, limit int, page int) (*gorm.DB, error) {
	// Set limit dan pagination
	if limit == 0 {
		limit = 15
	}
	query = query.Limit(limit)

	if page == 0 {
		page = 1
	}
	query = query.Offset((page - 1) * limit)

	return query, nil
}
