package query_builder_umkm

import (
	"fmt"
	"strings"
	querybuilder "umkm/query_builder"

	"gorm.io/gorm"
)

type UmkmQueryBuilder interface {
	querybuilder.BaseQueryBuilderList
	GetBuilder(filters string, limit int, page int) (*gorm.DB, error)
    GetBuilderWebList(filters string, limit int, page int, KategoriUmkm string) (*gorm.DB, error)
    GetBuilderDetailList( limit int, page int) (*gorm.DB, error) 
}

type UmkmQueryBuilderImpl struct {
	querybuilder.BaseQueryBuilderList
	db *gorm.DB
}

func NewUmkmQueryBuilder(db *gorm.DB) *UmkmQueryBuilderImpl {
	return &UmkmQueryBuilderImpl{
		BaseQueryBuilderList: querybuilder.NewBaseQueryBuilderList(db),
		db: db,
	}
}

func (umkmQueryBuilder *UmkmQueryBuilderImpl) GetBuilder(filters string, limit int, page int) (*gorm.DB, error) {
    query := umkmQueryBuilder.db

    // Tambahkan filter sebelum pagination
    if filters != "" {
        searchPattern := "%" + filters + "%"
        query = query.Where("name ILIKE ?", searchPattern)
    }

    // Set default limit jika limit == 0
    if limit <= 0 {
        limit = 15
    }

    // Hitung offset
    offset := (page - 1) * limit
    query = query.Offset(offset).Limit(limit)

    // Apply pagination
    query = query.Preload("HakAkses")

    return query, nil
}


func (umkmQueryBuilder *UmkmQueryBuilderImpl) GetBuilderWebList(filters string, limit int, page int, KategoriUmkm string) (*gorm.DB, error) {
    query := umkmQueryBuilder.db

    // Tambahkan filter sebelum pagination
    if filters != "" {
        searchPattern := "%" + filters + "%"
        query = query.Where("name ILIKE ?", searchPattern)
    }

    if KategoriUmkm != "" {
		kategoriList := strings.Split(KategoriUmkm, ",")
		var queryConditions []string
		var queryParams []interface{}
		for _, kategori := range kategoriList {
			kategori = strings.TrimSpace(kategori) // Remove extra spaces
			queryConditions = append(queryConditions, "kategori_umkm_id @> ?")
			queryParams = append(queryParams, fmt.Sprintf(`{"nama": ["%s"]}`, kategori))
		}
		query = query.Where(strings.Join(queryConditions, " OR "), queryParams...)
	}

    // Set default limit jika limit == 0
    if limit <= 0 {
        limit = 15
    }

    // Hitung offset
    offset := (page - 1) * limit
    query = query.Offset(offset).Limit(limit)

    // Apply pagination
    query = query.Preload("HakAkses")

    return query, nil
}

func (umkmQueryBuilder *UmkmQueryBuilderImpl) GetBuilderDetailList( limit int, page int) (*gorm.DB, error) {
    query := umkmQueryBuilder.db

    // Set default limit jika limit == 0
    if limit <= 0 {
        limit = 15
    }

    // Hitung offset
    offset := (page - 1) * limit
    query = query.Offset(offset).Limit(limit)

    // Apply pagination
    query = query.Preload("Produk")

    return query, nil
}
