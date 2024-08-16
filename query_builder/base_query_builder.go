package querybuilder

import (
	"fmt"
	"slices"
	"time"
	"umkm/model/domain"

	"gorm.io/gorm"
)

type BaseQueryBuilder interface {
	GetQueryBuilder(filters map[string]string, allowedFilters []string) (*gorm.DB, error)
}

type BaseQueryBuilderImpl struct {
	db *gorm.DB
}

func NewBaseQueryBuilder(db *gorm.DB) *BaseQueryBuilderImpl {
	return &BaseQueryBuilderImpl{
		db: db,
	}
}

func (baseQuerybuilder *BaseQueryBuilderImpl) GetQueryBuilder(filters map[string]string, allowedFilters []string) (*gorm.DB, error) {
    query := baseQuerybuilder.db.Model(&domain.Transaksi{}) // Sesuaikan model jika perlu

    // Filter berdasarkan tanggal jika ada
    if dateStr, ok := filters["tanggal"]; ok {
        if !slices.Contains(allowedFilters, "tanggal") {
            return nil, fmt.Errorf("filter 'tanggal' tidak diizinkan")
        }

        date, err := time.Parse("2006-01-02", dateStr)
        if err != nil {
            return nil, fmt.Errorf("format tanggal tidak valid: %v", err)
        }
        // Hanya filter berdasarkan tanggal dengan operator yang sesuai
        query = query.Where("DATE(tanggal) = ?", date.Format("2006-01-02"))
    }

    // Filter untuk field lain (teks) dengan ILIKE
    for field, value := range filters {
        if !slices.Contains(allowedFilters, field) {
            continue
        }

        // Pastikan hanya menggunakan ILIKE untuk field teks
        if field == "tanggal" { 
            continue
        }

        pat := "%" + value + "%"
        query = query.Where(fmt.Sprintf("%s ILIKE ?", field), pat)
    }

    return query, nil
}
