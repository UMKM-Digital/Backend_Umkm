package umkmrepo

import (
	"context"
	"errors"
	"umkm/model/domain"
	query_builder_umkm "umkm/query_builder/umkm"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RepoUmkmImpl struct {
	db               *gorm.DB
	umkmQueryBuilder query_builder_umkm.UmkmQueryBuilder
}

func NewUmkmRepositoryImpl(db *gorm.DB, umkmQueryBuilder query_builder_umkm.UmkmQueryBuilder) *RepoUmkmImpl {
	return &RepoUmkmImpl{
		db:               db,
		umkmQueryBuilder: umkmQueryBuilder,
	}
}

func (repo *RepoUmkmImpl) CreateRequest(umkm domain.UMKM) (domain.UMKM, error) {
	err := repo.db.Create(&umkm).Error
	if err != nil {
		return domain.UMKM{}, err
	}

	return umkm, nil
}

func (repo *RepoUmkmImpl) GetUmkmListByIds(ctx context.Context, umkmIDs []uuid.UUID, filters string, limit int, page int) ([]domain.UMKM, int, int, int, *int, *int, error) {
	var umkm []domain.UMKM
	var totalcount int64

	// Set default limit jika limit == 0
	if limit <= 0 {
		limit = 15
	}

	// Dapatkan query dengan filter dan pagination
	query, err := repo.umkmQueryBuilder.GetBuilder(filters, limit, page)
	if err != nil {
		return nil, 0, 0, 0, nil, nil, err
	}

	// Filter berdasarkan umkmIDs
	err = query.Where("id IN (?)", umkmIDs).Find(&umkm).Error
	if err != nil {
		return nil, 0, 0, 0, nil, nil, err
	}

	if len(umkm) == 0 {
		umkm = []domain.UMKM{}
	}

	// Hitung total records dari hasil pencarian, tanpa pagination
	totalQuery, err := repo.umkmQueryBuilder.GetBuilder(filters, 0, 0) // Tanpa pagination
	if err != nil {
		return nil, 0, 0, 0, nil, nil, err
	}

	// Hitung jumlah total records
	err = totalQuery.Model(&domain.UMKM{}).Where("id IN (?)", umkmIDs).Count(&totalcount).Error
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

	return umkm, int(totalcount), currentPage, totalPages, nextPage, prevPage, nil
}

func (repo *RepoUmkmImpl) GetUmkmFilterName(ctx context.Context, umkmIDs []uuid.UUID) ([]domain.UMKM, error) {
	var umkm []domain.UMKM
	err := repo.db.Where("id IN (?)", umkmIDs).Find(&umkm).Error
	if err != nil {
		return []domain.UMKM{}, err
	}
	return umkm, nil
}

func (repo *RepoUmkmImpl) GetUmkmListWeb(ctx context.Context, umkmIds []uuid.UUID) ([]domain.UMKM, error) {
	var umkm []domain.UMKM

	err := repo.db.Where("id IN (?)", umkmIds).Find(&umkm).Error
	if err != nil {
		return []domain.UMKM{}, err
	}
	return umkm, nil
}

func (repo *RepoUmkmImpl) GetUmkmID(id uuid.UUID) (domain.UMKM, error) {
	var umkm domain.UMKM
	if err := repo.db.First(&umkm, "id = ?", id).Error; err != nil {
		return umkm, err
	}
	return umkm, nil
}

func (repo *RepoUmkmImpl) UpdateUmkmId(id uuid.UUID, umkm domain.UMKM) (domain.UMKM, error) {
	if err := repo.db.Model(&domain.UMKM{}).Where("id = ?", id).Updates(umkm).Error; err != nil {
		return domain.UMKM{}, errors.New("gagal memperbarui umkm")
	}
	return umkm, nil
}

func (repo *RepoUmkmImpl) GetUmkmList(filters string, limit int, page int, kategori_umkm string, sortOrder string) ([]domain.UMKM, int, int, int, *int, *int, error) {
	var umkm []domain.UMKM
	var totalcount int64

	if limit <= 0 {
		limit = 15
	}

	if page <= 0 {
		page = 1 // Default ke halaman pertama jika page <= 0
	}

	query, err := repo.umkmQueryBuilder.GetBuilderWebList(filters, limit, page, kategori_umkm, sortOrder)
	if err != nil {
		return nil, 0, 0, 0, nil, nil, err
	}

	// Menggunakan Preload untuk mengambil produk terkait
	err = query.Preload("Produk").Find(&umkm).Error
	if err != nil {
		return nil, 0, 0, 0, nil, nil, err
	}

	ProdukQueryBuilder, err := repo.umkmQueryBuilder.GetBuilderWebList(filters, 0, 0, kategori_umkm, sortOrder)
	if err != nil {
		return nil, 0, 0, 0, nil, nil, err
	}
	err = ProdukQueryBuilder.Model(&domain.UMKM{}).Count(&totalcount).Error
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

	return umkm, int(totalcount), currentPage, totalPages, nextPage, prevPage, nil
}

// func(RepoUmkmImpl *RepoUmkmImpl) GetUmkmListDetailId(id uuid.UUID) ([]domain.UMKM, error){
//     var umkm []domain.UMKM
// 	err := RepoUmkmImpl.db.Preload("Produk").Find(&umkm, "id = ?", id).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return umkm, nil
// }

func (repo *RepoUmkmImpl) GetUmkmListDetailPaginated(id uuid.UUID, limit int, page int) ([]domain.UMKM, int, int, int, *int, *int, error) {
	var umkm []domain.UMKM
	var totalCount int64

	// Menghitung total data sebelum melakukan query dengan pagination
	err := repo.db.Model(&domain.UMKM{}).Where("id = ?", id).Error
	if err != nil {
		return nil, 0, 0, 0, nil, nil, err
	}

	// Menghitung offset berdasarkan halaman saat ini
	offset := (page - 1) * limit

	// Menggunakan query builder untuk mendapatkan detail UMKM
	err = repo.db.Preload("Produk").
		Where("id = ?", id).
		Find(&umkm).Error // Memuat UMKM, tetapi tidak melakukan limit di sini
	if err != nil {
		return nil, 0, 0, 0, nil, nil, err
	}

	// Mengambil produk dengan pagination
	var produkList []domain.Produk
	err = repo.db.Model(&domain.Produk{}).
		Where("umkm_id = ?", id).
		Order("created_at ASC"). // Ganti `umkm_id` dengan nama kolom yang benar jika berbeda
		Limit(limit).
		Offset(offset).
		Count(&totalCount).
		Find(&produkList).Error

	if err != nil {
		return nil, 0, 0, 0, nil, nil, err
	}

	// Menghitung total halaman
	totalPages := int((totalCount + int64(limit) - 1) / int64(limit))

	// Menghitung halaman berikutnya dan sebelumnya
	var nextPage *int
	if page < totalPages {
		nextPageVal := page + 1
		nextPage = &nextPageVal
	}

	var prevPage *int
	if page > 1 {
		prevPageVal := page - 1
		prevPage = &prevPageVal
	}

	// Menambahkan produk yang telah dipaginasikan ke dalam UMKM
	for i := range umkm {
		umkm[i].Produk = produkList // Mengupdate produk untuk UMKM
	}

	return umkm, int(totalCount), page, totalPages, nextPage, prevPage, nil
}

func (repo *RepoUmkmImpl) DeleteUmkmId(id uuid.UUID) error {
	if err := repo.db.Delete(&domain.UMKM{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (repo *RepoUmkmImpl) FindById(umkmId uuid.UUID) (domain.UMKM, error) {
	var umkm domain.UMKM
	err := repo.db.Where("id = ?", umkmId).First(&umkm).Error
	if err != nil {
		return umkm, err
	}
	return umkm, nil
}

func (repo *RepoUmkmImpl) ListUmkmActiceBack() ([]domain.UMKM, error) {
	var umkm []domain.UMKM
	err := repo.db.Order("created_at ASC").Find(&umkm).Error
	if err != nil {
		return nil, err
	}
	return umkm, nil
}

func (repo *RepoUmkmImpl) UpdateActiveId(idUmkm uuid.UUID, active int) error {
	if err := repo.db.Model(&domain.UMKM{}).Where("id = ?", idUmkm).Update("active", active).Error; err != nil {
		return errors.New("failed to update umkm active status")
	}
	return nil
}

// nampilin umkm active
func (repo *RepoUmkmImpl) GetUmkmActive(active int) ([]domain.UMKM, error) {
	var umkm []domain.UMKM
	err := repo.db.Raw("SELECT * FROM Umkm WHERE active = ?", active).Scan(&umkm).Error
	if err != nil {
		return nil, err
	}
	return umkm, nil
}
