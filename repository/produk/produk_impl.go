package produkrepo

import (
	// "errors"
	"umkm/model/domain"
	query_builder_produk "umkm/query_builder/produk"

	"github.com/google/uuid"
	"gorm.io/gorm"
	
)

type ProdukRepoImpl struct {
	db *gorm.DB
	produkQueryBuilder query_builder_produk.ProdukQueryBuilder
}

func NewProdukRepositoryImpl(db *gorm.DB, produkQueryBuilder query_builder_produk.ProdukQueryBuilder) *ProdukRepoImpl {
	return &ProdukRepoImpl{
		db: db,
		produkQueryBuilder: produkQueryBuilder,
	}
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

func (repo *ProdukRepoImpl) GetProduk(ProdukId uuid.UUID, filters string, limit int, page int, kategori_produk_id string) ([]domain.Produk, error) {
	var produk []domain.Produk

	// Memanggil produkQueryBuilder.GetBuilderProduk dengan parameter yang sesuai
	query, err := repo.produkQueryBuilder.GetBuilderProduk(filters, limit, page, kategori_produk_id)
	if err != nil {
		return nil, err
	}

	 // Debug query sebelum dijalankan
	// Menerapkan filter umkm_id
	err = query.Where("umkm_id = ?", ProdukId).Find(&produk).Error
	if err != nil {
		return nil, err
	}
	return produk, nil
}
