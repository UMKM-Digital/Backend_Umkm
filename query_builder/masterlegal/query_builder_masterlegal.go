package query_builder_masterlegal

import (
	querybuilder "umkm/query_builder"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MasteLegalQueryBuilder interface {
	querybuilder.BaseQueryBuilderList
	GetBuilderMasterLegal(filters string, limit int, page int) (*gorm.DB, error)
	GetBuilderDokumenUmkmStatus(umkmId uuid.UUID, filters string, limit int, page int) (*gorm.DB, error)
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


func (builder *MasteLegalQueryBuilderImpl) GetBuilderDokumenUmkmStatus(umkmId uuid.UUID, filters string, limit int, page int) (*gorm.DB, error) {
	query := builder.db.Table("master_dokumen_legal").
		Select("master_dokumen_legal.nama, CASE WHEN umkm_dokumen_legal.id IS NOT NULL THEN 1 ELSE 0 END AS status").
		Joins("LEFT JOIN umkm_dokumen_legal ON master_dokumen_legal.id = umkm_dokumen_legal.dokumen_id AND umkm_dokumen_legal.umkm_id = ?", umkmId)

	// Implementasi filter jika ada
	if filters != "" {
		searchPattern := "%" + filters + "%"
		query = query.Where("master_dokumen_legal.nama ILIKE ?", searchPattern)
	}
	
	// Tambahkan pagination (limit dan offset)
	offset := (page - 1) * limit
	query = query.Offset(offset).Limit(limit)
	
	// Gunakan GetQueryBuilderList dengan menambahkan `builder` sebagai argumen pertama
	query, err := builder.GetQueryBuilderList( query, limit, page)
	if err != nil {
		return nil, err
	}
	
	return query, nil
}
