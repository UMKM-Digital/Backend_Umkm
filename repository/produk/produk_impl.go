package produkrepo

import (
	// "errors"
	"umkm/model/domain"

	"github.com/google/uuid"
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

//
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



// func (repo *ProdukRepoImpl) ProdukById(id uuid.UUID) (domain.Produk, error){
// 	var produk domain.Produk
	
// 	err := repo.db.Find(&produk, "id = ?", id).Error


// 	if err != nil {
// 		return domain.Produk{},errors.New("produk tidak ditemukan")
// 	}

// 	return produk, nil
// }

func (repo *ProdukRepoImpl) GetProduk(ProdukId uuid.UUID)([]domain.Produk, error){
	var produk []domain.Produk
	err := repo.db.Where("umkm_id = ?", ProdukId).Find(&produk).Error
	if err != nil{
		return nil, err
	}
	return produk, nil
}