package query_builder_produk

import (
	querybuilder "umkm/query_builder"

	"gorm.io/gorm"
	"fmt"
	"strings"
)

type ProdukQueryBuilder interface {
	querybuilder.BaseQueryBuilderList
	GetBuilderProduk(filters string, limit int, page int, kategori_produk_id string) (*gorm.DB, error) 
}

type ProdukQueryBuilderImpl struct {
	querybuilder.BaseQueryBuilderList
	db *gorm.DB
}

func NewProdukQueryBuilder(db *gorm.DB) *ProdukQueryBuilderImpl {
	return &ProdukQueryBuilderImpl{
		BaseQueryBuilderList: querybuilder.NewBaseQueryBuilderList(db),
		db: db,
	}
}

func (produkQueryBuilder *ProdukQueryBuilderImpl) GetBuilderProduk(filters string, limit int, page int, kategori_produk_id string) (*gorm.DB, error) {
	query := produkQueryBuilder.db

	// Implementasi filter di sini
	if filters != "" {
		searchPattern := "%" + filters 
		query = query.Where("nama ILIKE ?", searchPattern)
	}

	if kategori_produk_id != "" {
		kategoriList := strings.Split(kategori_produk_id, ",")
		var queryConditions []string
		var queryParams []interface{}
		for _, kategori := range kategoriList {
			kategori = strings.TrimSpace(kategori) // Remove extra spaces
			queryConditions = append(queryConditions, "kategori_produk_id @> ?")
			queryParams = append(queryParams, fmt.Sprintf(`{"nama": ["%s"]}`, kategori))
		}
		query = query.Where(strings.Join(queryConditions, " OR "), queryParams...)
	}

	query, err := produkQueryBuilder.GetQueryBuilderList(query, filters, limit, page)
	if err != nil {
		return nil, err
	}

	query = query.Preload("Umkm")

	return query, nil
}
