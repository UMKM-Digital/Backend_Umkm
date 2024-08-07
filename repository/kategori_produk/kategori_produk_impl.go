package kategoriprodukrepo

import (
	"umkm/model/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type KategoriProdukRepoImpl struct {
	db *gorm.DB
}

func NewKategoriProdukRepo(db *gorm.DB) *KategoriProdukRepoImpl {
	return &KategoriProdukRepoImpl{db: db}
}

func (repo *KategoriProdukRepoImpl) CreateKategoriProduk(kategoriproduk domain.KategoriProduk) (domain.KategoriProduk, error) {
	err := repo.db.Create(&kategoriproduk).Error
	if err != nil {
		return domain.KategoriProduk{}, err
	}
	return kategoriproduk, nil
}

// func (repo *KategoriProdukRepoImpl) GetKategoriUmkm(umkmID uuid.UUID) ([]domain.KategoriProduk, error) {
//     var kategori []domain.KategoriProduk

//     // Query dengan filter berdasarkan umkm_id
//     err := repo.db.Where("umkm_id = ?", umkmID).Find(&kategori).Error
//     if err != nil {
//         return nil, err
//     }

//     return kategori, nil
// }

func (repo *KategoriProdukRepoImpl) GetKategoriUmkm(umkmID uuid.UUID) ([]domain.KategoriProduk, error) {
    var kategori []domain.KategoriProduk
    err := repo.db.Where("umkm_id = ?", umkmID).Find(&kategori).Error
    if err != nil {
        return nil, err
    }
    return kategori, nil
}




