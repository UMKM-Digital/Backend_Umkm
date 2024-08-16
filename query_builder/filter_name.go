package querybuilder

import (
	"slices"
	"umkm/model/domain"
	"gorm.io/gorm"
	"fmt"
)

type BaseQueryBuilderName interface {
	GetQueryBuilderName(filters map[string]string, allowedFilters []string) (*gorm.DB, error) 
}

type BaseQueryBuilderNameImpl struct {
	db *gorm.DB
}

func NewBaseQueryBuilderName(db *gorm.DB) *BaseQueryBuilderNameImpl {
	return &BaseQueryBuilderNameImpl{
		db: db,
	}
}

func (baseQuerybuilderName *BaseQueryBuilderNameImpl) GetQueryBuilderName(filters map[string]string, allowedFilters []string) (*gorm.DB, error) {
    query := baseQuerybuilderName.db.Model(&domain.UMKM{})

    for field, value := range filters {
        if !slices.Contains(allowedFilters, field) {
            continue
        }

        // Penanganan khusus untuk filter name
        if field == "name" {
            pat := "%" + value + "%"
            query = query.Where("name ILIKE ?", pat)
            continue
        }

        // Penanganan untuk field lainnya
        pat := "%" + value + "%"
        query = query.Where(fmt.Sprintf("%s ILIKE ?", field), pat)
    }

    return query, nil
}
