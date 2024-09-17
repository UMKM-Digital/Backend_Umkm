package kategoriprodukrepo

// import (
// 	"errors"
// 	"umkm/model/domain"
// 	query_builder_kategori_produk "umkm/query_builder/kategoriproduk"

// 	"github.com/google/uuid"
// 	"gorm.io/gorm"
// )

// type KategoriProdukRepoImpl struct {
// 	db *gorm.DB
// 	KategoriProdukQueryBuilder query_builder_kategori_produk.KategoriProdukQueryBuilder
// }

// func NewKategoriProdukRepo(db *gorm.DB, KategoriProdukQueryBuilder query_builder_kategori_produk.KategoriProdukQueryBuilder) *KategoriProdukRepoImpl {
// 	return &KategoriProdukRepoImpl{
// 		db: db,
// 		KategoriProdukQueryBuilder: KategoriProdukQueryBuilder,
// 	}
// }

// func (repo *KategoriProdukRepoImpl) CreateKategoriProduk(kategoriproduk domain.KategoriProduk) (domain.KategoriProduk, error) {
// 	err := repo.db.Create(&kategoriproduk).Error
// 	if err != nil {
// 		return domain.KategoriProduk{}, err
// 	}
// 	return kategoriproduk, nil
// }

// func (repo *KategoriProdukRepoImpl) GetKategoriProduk(umkmID uuid.UUID, filters string, limit int, page int) ([]domain.KategoriProduk, int, error) {
//     var kategori []domain.KategoriProduk
//     var totalcount int64

//     // Gunakan satu query builder untuk semua operasi
//     query, err := repo.KategoriProdukQueryBuilder.GetBuilder(filters, limit, page)
//     if err != nil {
//         return nil, 0, err
//     }

//     // Terapkan filter untuk mendapatkan data
//     err = query.Where("umkm_id = ?", umkmID).Find(&kategori).Error
//     if err != nil {
//         return nil, 0, err
//     }

//     // Terapkan filter yang sama untuk menghitung total count
//     err = query.Model(&domain.KategoriProduk{}).Where("umkm_id = ?", umkmID).Count(&totalcount).Error
//     if err != nil {
//         return nil, 0, err
//     }

//     return kategori, int(totalcount), nil
// }

// func (repo *KategoriProdukRepoImpl) GetKategoriProdukId(idproduk int) (domain.KategoriProduk, error){
//     var KategoriUmkmData domain.KategoriProduk

//     err := repo.db.Find(&KategoriUmkmData, "id = ?", idproduk).Error

// 	if err != nil {
// 		return domain.KategoriProduk{},errors.New("kategori tidak ditemukan")
// 	}

// 	return KategoriUmkmData, nil
// }

// func (repo *KategoriProdukRepoImpl) UpdateKategoriId(idProduk int, kategori domain.KategoriProduk) (domain.KategoriProduk, error) {
// 	if err := repo.db.Model(&domain.KategoriProduk {}).Where("id = ?", idProduk).Updates(kategori).Error; err != nil {
//         return domain.KategoriProduk{}, errors.New("failed to update kategori produk")
//     }
//     return kategori, nil
// }

// func (repo *KategoriProdukRepoImpl) DeleteKategoriProdukId(idproduk int) error {
//     if err := repo.db.Delete(&domain.KategoriProduk{}, idproduk).Error; err != nil{
//         return err
//     }
//     return nil
// }