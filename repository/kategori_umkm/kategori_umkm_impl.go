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
	func (repo *KategoriRepoUmkmImpl) GetKategoriUmkm(filters string, limit int, page int) ([]domain.Kategori_Umkm, int, error) {
		var kategori []domain.Kategori_Umkm
		var totalcount int64

		query, err := repo.kategoriQuerybuilder.GetBuilder(filters,limit, page)
		if err != nil{
			return nil, 0, err
		}

		err = query.Find(&kategori).Error
		if err != nil {
			return nil, 0, err
		}

		kategoriQuerybuilder, err := repo.kategoriQuerybuilder.GetBuilder(filters, 0, 0)
		if err != nil {
			return nil, 0, err
		}
		
		err = kategoriQuerybuilder.Model(&domain.Kategori_Umkm{}).Count(&totalcount).Error
		if err != nil {
			return nil, 0, err
		}

		return kategori, int(totalcount), nil
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
        return domain.Kategori_Umkm{}, errors.New("failed to update profile")
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




