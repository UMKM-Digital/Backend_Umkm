package produkrepo

import (
	"context"
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

func (repo *ProdukRepoImpl) GetProduk(ProdukId uuid.UUID, filters string, limit int, page int, kategori_produk_id string, sort string) ([]domain.Produk, int, int, int, *int, *int, error) {
	var produk []domain.Produk
	var totalcount int64

	if limit <= 0 {
        limit = 15
    }
	// Mendapatkan query dengan limit dan pagination
	query, err := repo.produkQueryBuilder.GetBuilderProduk(filters, limit, page, kategori_produk_id, sort)
	if err != nil {
		return nil, 0, 0, 0, nil, nil, err
	}

	// Mendapatkan data produk
	err = query.Where("umkm_id = ?", ProdukId).Order("created_at ASC").Find(&produk).Error
	if err != nil {
		return nil, 0, 0, 0, nil, nil, err
	}

	ProdukQueryBuilder, err := repo.produkQueryBuilder.GetBuilderProduk(filters, 0, 0, kategori_produk_id, sort)
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


	func (r *ProdukRepoImpl) GetProductsByUmkmIds(ctx context.Context, umkmIDs []uuid.UUID) ([]domain.Produk, error) {
		var products []domain.Produk
	
		// Menggunakan GORM untuk query dengan kondisi `IN`
		if err := r.db.WithContext(ctx).Where("umkm_id IN ?", umkmIDs).Find(&products).Error; err != nil {
			return nil, err
		}
	
		return products, nil
	}

	func (repo *ProdukRepoImpl) GetProdukList(limit int, page int, filters string, kategoriproduk string, sort string) ([]domain.Produk, int, int, int, *int, *int, error) {
		var produk []domain.Produk
		var totalcount int64
	
		if limit <= 0 {
			limit = 15
		}
	
		// Mendapatkan query dengan limit, pagination, filter, kategori produk, dan sort
		query, err := repo.produkQueryBuilder.GetBuilderProdukListWeb(limit, page, filters, kategoriproduk, sort)
		if err != nil {
			return nil, 0, 0, 0, nil, nil, err
		}
	
		// Eksekusi query untuk mendapatkan daftar produk
		err = query.Preload("Umkm").Find(&produk).Error
		if err != nil {
			return nil, 0, 0, 0, nil, nil, err
		}
	
		// Mendapatkan total count untuk pagination
		ProdukQueryBuilder, err := repo.produkQueryBuilder.GetBuilderProdukListWeb(0, 0, filters, kategoriproduk, sort)
		if err != nil {
			return nil, 0, 0, 0, nil, nil, err
		}
	
		err = ProdukQueryBuilder.Model(&domain.Produk{}).Count(&totalcount).Error
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
	

	func(repo *ProdukRepoImpl) FindWebId(id uuid.UUID) (domain.Produk, error){
		var produk domain.Produk
	if err := repo.db.Preload("Umkm").Find(&produk, "id = ?", id).Error; err != nil {
		return produk, err
	}
	return produk, nil
	}

	func(repo *ProdukRepoImpl) DeleteProdukUmkmId(id uuid.UUID) error{
		return repo.db.Where("umkm_id = ?", id).Delete(&domain.Produk{}).Error
	}


	func (repo *ProdukRepoImpl) GetProdukByUmkmId(id uuid.UUID) ([]domain.Produk, error) {
		var produkList []domain.Produk
		if err := repo.db.Where("umkm_id = ?", id).Find(&produkList).Error; err != nil {
			return produkList, err
		}
		return produkList, nil
	}
	

	func (repo *ProdukRepoImpl) GetProdukByUmkmLogin(umkmIds []uuid.UUID) ([]domain.Produk, error) {
		var produkList []domain.Produk
		// Ambil produk yang umkm_id-nya ada dalam daftar umkmIds
		if err := repo.db.Order("created_at ASC").Where("umkm_id IN (?)", umkmIds).Find(&produkList).Error; err != nil {
			return produkList, err
		}
		return produkList, nil
	}
	
	func(repo *ProdukRepoImpl) GetProdukBaru(umkmId uuid.UUID)([]domain.Produk, error){
		var produkList []domain.Produk
		// Ambil produk yang umkm_id-nya ada dalam daftar umkmIds
		if err := repo.db.Order("created_at DESC"). Limit(4).
		Where("umkm_id IN (?)", umkmId).
		Find(&produkList).Error; err != nil {
			return produkList, err
		}
		return produkList, nil
	}

	//list produk semua
	func(repo *ProdukRepoImpl) GetTopProduk(idUmkm uuid.UUID) ([]domain.Produk, error) {
		var produklis []domain.Produk
		query := repo.db.Preload("Umkm").Order("id ASC")
	
		if idUmkm != uuid.Nil {
			// Filter berdasarkan id_umkm jika parameter diberikan
			query = query.Where("umkm_id = ?", idUmkm)
		}
	
		err := query.Find(&produklis).Error
		if err != nil {
			return nil, err
		}
	
		return produklis, nil
	}
	
	//update active proudk

	func(repo *ProdukRepoImpl) UpdateTopProduk(idproduk uuid.UUID, active int) error {
		result := repo.db.Model(&domain.Produk{}).Where("id = ?", idproduk).Update("active", active)
	
		// Periksa apakah ada baris yang diperbarui
		if result.Error != nil {
			return errors.New("failed to update produk active status")
		}
		if result.RowsAffected == 0 {
			return errors.New("no rows affected, produk may not exist")
		}
		
		return nil
	}
	

	func (repo *ProdukRepoImpl) GetProdukActive(active int)([]domain.Produk, error) {
		var produk []domain.Produk
		err := repo.db.Preload("Umkm").Where("active = ?", active).Find(&produk).Error
		if err != nil {
			return nil, err
		}
		return produk, nil
	}
	