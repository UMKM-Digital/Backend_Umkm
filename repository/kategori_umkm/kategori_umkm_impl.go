package repokategoriumkm

import (
	"errors"
	"umkm/model/domain"
	query_builder_kategori_umkm "umkm/query_builder/kategoriumkm"

	"gorm.io/gorm"
)

type KategoriRepoUmkmImpl struct {
	db *gorm.DB
	kategoriQuerybuilder query_builder_kategori_umkm.KategoriUmkmQueryBuilder
}

func NewKategoriUmkmRepositoryImpl(db *gorm.DB, kategoriQuerybuilder query_builder_kategori_umkm.KategoriUmkmQueryBuilder) *KategoriRepoUmkmImpl{
	return &KategoriRepoUmkmImpl{
		db:db,
		kategoriQuerybuilder: kategoriQuerybuilder,
	}
}

 func (repo *KategoriRepoUmkmImpl) CreateRequest(categoryumkm domain.Kategori_Umkm) (domain.Kategori_Umkm, error){
	err := repo.db.Create(&categoryumkm).Error
	if err != nil{
		return domain.Kategori_Umkm{}, err
	}

	return categoryumkm, nil
 }

 //membaca seluruh list kategori
		func (repo *KategoriRepoUmkmImpl) GetKategoriUmkm(filters string, limit int, page int) ([]domain.Kategori_Umkm, int, int, int, *int, *int, error) {
			var kategori []domain.Kategori_Umkm
			var totalcount int64

			if limit <= 0 {
				limit = 15
			}	

			query, err := repo.kategoriQuerybuilder.GetBuilder(filters,limit, page)
			if err != nil{
				return nil, 0, 0, 0, nil, nil, err
			}

			err = query.Order("id ASC").Find(&kategori).Error
			if err != nil {
				return nil, 0, 0, 0, nil, nil, err
			}

			kategoriQuerybuilder, err := repo.kategoriQuerybuilder.GetBuilder(filters, 0, 0)
			if err != nil {
				return nil, 0, 0, 0, nil, nil, err
			}
			
			err = kategoriQuerybuilder.Model(&domain.Kategori_Umkm{}).Count(&totalcount).Error
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

//get kategori by id
func(repo *KategoriRepoUmkmImpl) GetKategoriUmkmId(idKategori int) (domain.Kategori_Umkm, error){
	var KategoriUmkmData domain.Kategori_Umkm

	err := repo.db.Find(&KategoriUmkmData, "id = ?", idKategori).Error

	if err != nil {
		return domain.Kategori_Umkm{},errors.New("kategori tidak ditemukan")
	}

	return KategoriUmkmData, nil
}

//update kategori
func (repo *KategoriRepoUmkmImpl) UpdateKategoriId(idKategori int, kategori domain.Kategori_Umkm) (domain.Kategori_Umkm, error) {
    if err := repo.db.Model(&domain.Kategori_Umkm{}).Where("id = ?", idKategori).Updates(kategori).Error; err != nil {
        return domain.Kategori_Umkm{}, errors.New("failed to update kategori umkm")
    }
    return kategori, nil
}

//delete
func (repo *KategoriRepoUmkmImpl) DeleteKategoriUmkmId(id int) error {
    if err := repo.db.Delete(&domain.Kategori_Umkm{}, id).Error; err != nil {
        return err
    }
    return nil
}

func (repo *KategoriRepoUmkmImpl) GetKategoriUmkmBySektor(id int)([]domain.Kategori_Umkm, error){
	var kategoriumkm []domain.Kategori_Umkm
	err := repo.db.Where("id_sektor_usaha  = ?", id).Find(&kategoriumkm).Error
	if err != nil{
		return nil, err
	}
	return kategoriumkm, nil
}