package querybuilder

import (
	"fmt"
	"slices"

	"gorm.io/gorm"
)

type BaseQueryBuilderTransaksi interface {
	GetQueryBuilder(filters map[string]string, limit string, page int) (*gorm.DB, error)
}

type BaseQueryBuilderMain struct {
	db *gorm.DB
}

func NewBaseQueryBuilderMain(db *gorm.DB) *BaseQueryBuilderMain {
	return &BaseQueryBuilderMain{db: db}
}

// Implementasi GetQueryBuilder dengan signature yang diharapkan
func (baseQueryBuilder *BaseQueryBuilderMain) GetQueryBuilder(filters map[string]string, limit string, page int, allowedFilters []string) (*gorm.DB, error) {
	query := baseQueryBuilder.db

	for filter, value := range filters {
		if !slices.Contains(allowedFilters, filter) {
			return nil, fmt.Errorf("%s filter is not allowed", filter)
		}
		pat := "%" + value + "%"
		query = query.Where(fmt.Sprintf("%s ilike ?", filter), pat)
	}

	// Set default limit dan tangani kasus khusus
	defaultLimit := 5
	switch limit {
	case "10":
		defaultLimit = 10
	case "25":
		defaultLimit = 25
	case "all":
		defaultLimit = -1 // Tanpa batas
	default:
		defaultLimit = 5
	}

	if defaultLimit > 0 {
		query = query.Limit(defaultLimit)
	}

	if page == 0 {
		page = 1
	}
	if defaultLimit > 0 {
		query = query.Offset((page - 1) * defaultLimit)
	}

	return query, nil
}
