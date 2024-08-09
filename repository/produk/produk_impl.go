package produkrepo

import (
	"umkm/model/domain"

	"gorm.io/gorm"
)

type ProdukRepoImpl struct {
	db *gorm.DB
}

func NewProdukRepositoryImpl(db *gorm.DB) *ProdukRepoImpl {
	return &ProdukRepoImpl{db: db}
}

func (repo *ProdukRepoImpl) CreateRequest(produk domain.Produk)(domain.Produk, error) {
	err := repo.db.Create(&produk).Error
	if err != nil {
		return domain.Produk{}, err
	}

	return produk, nil
}