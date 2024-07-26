package repokategoriumkm

import (
	"errors"
	"umkm/model/domain"

	"gorm.io/gorm"
)

type KategoriRepoUmkmImpl struct {
	db *gorm.DB
}

func NewKategoriUmkmRepositoryImpl(db *gorm.DB) *KategoriRepoUmkmImpl{
	return &KategoriRepoUmkmImpl{db:db}
}

 func (repo *KategoriRepoUmkmImpl) CreateRequest(categoryumkm domain.Kategori_Umkm) (domain.Kategori_Umkm, error){
	err := repo.db.Create(&categoryumkm).Error
	if err != nil{
		return domain.Kategori_Umkm{}, err
	}

	return categoryumkm, nil
 }

 //membaca seluruh list kategori
 func (repo *KategoriRepoUmkmImpl) GetKategoriUmkm() ([]domain.Kategori_Umkm, error) {
	var kategori []domain.Kategori_Umkm
	err := repo.db.Find(&kategori).Error
	if err != nil {
		return []domain.Kategori_Umkm{}, err
	}
	return kategori, nil
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