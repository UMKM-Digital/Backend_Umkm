package repokategoriumkm

import (
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

 func (repo *KategoriRepoUmkmImpl) GetKategoriUmkm() ([]domain.Kategori_Umkm, error) {
	var kategori []domain.Kategori_Umkm
	err := repo.db.Find(&kategori).Error
	if err != nil {
		return []domain.Kategori_Umkm{}, err
	}
	return kategori, nil
}