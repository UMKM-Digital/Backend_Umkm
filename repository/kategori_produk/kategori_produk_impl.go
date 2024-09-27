package kategoriprodukrepo

import (
	"errors"
	"umkm/model/domain"
	query_builder_kategori_produk "umkm/query_builder/kategoriproduk"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type KategoriProdukRepoImpl struct {
	db                         *gorm.DB
	KategoriProdukQueryBuilder query_builder_kategori_produk.KategoriProdukQueryBuilder
}

func NewKategoriProdukRepo(db *gorm.DB, KategoriProdukQueryBuilder query_builder_kategori_produk.KategoriProdukQueryBuilder) *KategoriProdukRepoImpl {
	return &KategoriProdukRepoImpl{
		db:                         db,
		KategoriProdukQueryBuilder: KategoriProdukQueryBuilder,
	}
}

func (repo *KategoriProdukRepoImpl) CreateKategoriProduk(kategoriproduk domain.KategoriProduk) (domain.KategoriProduk, error) {
	err := repo.db.Create(&kategoriproduk).Error
	if err != nil {
		return domain.KategoriProduk{}, err
	}
	return kategoriproduk, nil
}

func (repo *KategoriProdukRepoImpl) GetKategoriProduk( filters string, limit int, page int) ([]domain.KategoriProduk, int,int, int, *int, *int, error) {
	var kategori []domain.KategoriProduk
	var totalcount int64

	if limit <= 0 {
		limit = 15
	}

	// Gunakan satu query builder untuk semua operasi
	query, err := repo.KategoriProdukQueryBuilder.GetBuilder(filters, limit, page)
	if err != nil {
		 return nil, 0, 0, 0, nil, nil, err
	}

	// Terapkan filter untuk mendapatkan data
	err = query.Order("id ASC").Find(&kategori).Error
	if err != nil {
		 return nil, 0, 0, 0, nil, nil, err
	}

	// Terapkan filter yang sama untuk menghitung total count
	err = query.Model(&domain.KategoriProduk{}).Count(&totalcount).Error
	if err != nil {
		 return nil, 0, 0, 0, nil, nil, err
	}

	  // Hitung total pages
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

	return kategori, int(totalcount), currentPage, totalPages, nextPage, prevPage, nil
}

func (repo *KategoriProdukRepoImpl) GetKategoriProdukId(idproduk int) (domain.KategoriProduk, error) {
	var KategoriUmkmData domain.KategoriProduk

	err := repo.db.Find(&KategoriUmkmData, "id = ?", idproduk).Error

	if err != nil {
		return domain.KategoriProduk{}, errors.New("kategori tidak ditemukan")
	}

	return KategoriUmkmData, nil
}

func (repo *KategoriProdukRepoImpl) UpdateKategoriId(idProduk int, kategori domain.KategoriProduk) (domain.KategoriProduk, error) {
	if err := repo.db.Model(&domain.KategoriProduk{}).Where("id = ?", idProduk).Updates(kategori).Error; err != nil {
		return domain.KategoriProduk{}, errors.New("failed to update kategori produk")
	}
	return kategori, nil
}

func (repo *KategoriProdukRepoImpl) DeleteKategoriProdukId(idproduk int) error {
	if err := repo.db.Delete(&domain.KategoriProduk{}, idproduk).Error; err != nil {
		return err
	}
	return nil
}


func(repo *KategoriProdukRepoImpl) DeleteKategoriUmkmId(id uuid.UUID) error{
	return repo.db.Where("umkm_id = ?", id).Delete(&domain.KategoriProduk{}).Error
}