package querybuilder

import (
	"fmt"
	"slices"

	"gorm.io/gorm"
)

type BaseQueryBuilder interface {
	GetQueryBuilder(filters map[string]string,  allowedFilters []string) (*gorm.DB, error)
}

type BaseQueryBuilderImpl struct{
	db *gorm.DB
}

func NewBaseQueryBuilder(db *gorm.DB) *BaseQueryBuilderImpl{
	return &BaseQueryBuilderImpl{
		db: db,
	}
}

func (baseQuerybuilder *BaseQueryBuilderImpl) GetQueryBuilder(filters map[string]string,  allowedFilters []string) (*gorm.DB, error){
	query := baseQuerybuilder.db
	 
	for filters, value := range filters{
		if !slices.Contains(allowedFilters, filters){
			return nil, fmt.Errorf("transaksi tidak ada")
		}
		pat := "%" + value + "%"
		query = query.Where(fmt.Sprintf("%s ilike ?", filters), pat)
	}

	return query, nil
}