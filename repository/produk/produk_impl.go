package produkrepo

import (
	"errors"
	"umkm/model/domain"
	query_builder_produk "umkm/query_builder/produk"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProdukRepoImpl struct {
	db                 *gorm.DB
	produkQueryBuilder query_builder_produk.ProdukQueryBuilder
}

func NewProdukRepositoryImpl(db *gorm.DB, produkQueryBuilder query_builder_produk.ProdukQueryBuilder) *ProdukRepoImpl {
	return &ProdukRepoImpl{
		db:                 db,
		produkQueryBuilder: produkQueryBuilder,
	}
}

func (repo *ProdukRepoImpl) CreateRequest(produk domain.Produk) (domain.Produk, error) {
	err := repo.db.Create(&produk).Error
	if err != nil {
		return domain.Produk{}, err
	}

	return produk, nil
}

func (repo *ProdukRepoImpl) DeleteProdukId(id uuid.UUID) error {
	if err := repo.db.Delete(&domain.Produk{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (repo *ProdukRepoImpl) FindById(id uuid.UUID) (domain.Produk, error) {
	var produk domain.Produk
	if err := repo.db.First(&produk, "id = ?", id).Error; err != nil {
		return produk, err
	}
	return produk, nil
}

func (repo *ProdukRepoImpl) GetProduk(ProdukId uuid.UUID, filters string, limit int, page int, kategori_produk_id string) ([]domain.Produk, int, int, int, *int, *int, error) {
	var produk []domain.Produk
	var totalcount int64

	if limit <= 0 {
        limit = 15
    }
	// Mendapatkan query dengan limit dan pagination
	query, err := repo.produkQueryBuilder.GetBuilderProduk(filters, limit, page, kategori_produk_id)
	if err != nil {
		return nil, 0, 0, 0, nil, nil, err
	}

	// Mendapatkan data produk
	err = query.Where("umkm_id = ?", ProdukId).Find(&produk).Error
	if err != nil {
		return nil, 0, 0, 0, nil, nil, err
	}

	ProdukQueryBuilder, err := repo.produkQueryBuilder.GetBuilderProduk(filters, 0, 0, kategori_produk_id)
	if err != nil {
		return nil, 0, 0, 0, nil, nil, err
	}

	err = ProdukQueryBuilder.Model(&domain.Produk{}).Where("umkm_id = ?", ProdukId).Count(&totalcount).Error
	if err != nil {
		return nil, 0, 0, 0, nil, nil, err
	}
	totalPages := 1
    if limit > 0 {
        totalPages = int((totalcount + int64(limit) - 1) / int64(limit))
    }

    // Jika page > totalPages, return kosong
    if page > totalPages {
        return nil, int(totalcount), page, totalPages, nil, nil, nil
    }

    currentPage := page

    // Tentukan nextPage dan prevPage
    var nextPage *int
    if currentPage < totalPages {
        np := currentPage + 1
        nextPage = &np
    }

    var prevPage *int
    if currentPage > 1 {
        pp := currentPage - 1
        prevPage = &pp
    }

	return produk, int(totalcount), currentPage, totalPages, nextPage, prevPage, nil
}

func (repo *ProdukRepoImpl) UpdatedProduk(ProdukId uuid.UUID, produk domain.Produk) (domain.Produk, error) {
	if err := repo.db.Model(&domain.Produk{}).Where("id = ?", ProdukId).Updates(produk).Error; err != nil {
		return domain.Produk{}, errors.New("gagal memperbarui produk")
	}
	return produk, nil
}
