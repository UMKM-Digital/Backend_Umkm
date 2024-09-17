package query_builder_masterlegal

import (
	querybuilder "umkm/query_builder"
	"gorm.io/gorm"
)

type MasteLegalQueryBuilder interface {
	querybuilder.BaseQueryBuilderList
	GetBuilderMasterLegal(filters string, limit int, page int) (*gorm.DB, error)
}

type MasteLegalQueryBuilderImpl struct {
	querybuilder.BaseQueryBuilderList
	db *gorm.DB
}

func NewMasteLegalQueryBuilder(db *gorm.DB) *MasteLegalQueryBuilderImpl {
	return &MasteLegalQueryBuilderImpl{
		BaseQueryBuilderList: querybuilder.NewBaseQueryBuilderList(db),
		db:                   db,
	}
}

func (MasteLegalQueryBuilder *MasteLegalQueryBuilderImpl) GetBuilderMasterLegal(filters string, limit int, page int) (*gorm.DB, error) {
	query := MasteLegalQueryBuilder.db

	// Implementasi filter di sini
	if filters != "" {
		searchPattern := "%" + filters + "%"
		query = query.Where("nama ILIKE ?", searchPattern)
	}

	offset := (page - 1) * limit
    query = query.Offset(offset).Limit(limit)

	query, err := MasteLegalQueryBuilder.GetQueryBuilderList(query,  limit, page)
	if err != nil {
		return nil, err
	}

	return query, nil
}
