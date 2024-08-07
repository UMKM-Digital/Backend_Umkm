package kategoriprodukservice

import (
	"umkm/model/domain"
	"umkm/model/entity"
	"umkm/model/web"
	kategoriprodukrepo "umkm/repository/kategori_produk"

	"github.com/google/uuid"
)

type KategoriProdukServiceImpl struct {
    kategoriprodukrepository kategoriprodukrepo.KategoriProdukRepository
}

func NewKategoriProdukService(kategorirepository kategoriprodukrepo.KategoriProdukRepository) *KategoriProdukServiceImpl {
    return &KategoriProdukServiceImpl{
        kategoriprodukrepository: kategorirepository,
    }
}

func (service *KategoriProdukServiceImpl) CreateKategori(kategoriproduk web.CreateCategoriProduk) (map[string]interface{}, error) {
    newKategoriProduk := domain.KategoriProduk{
        Umkm: kategoriproduk.UmkmId, // This should align with uuid.UUID
        Nama: kategoriproduk.Name,
    }

    saveKategoriProduk, errSaveKategoriProduk := service.kategoriprodukrepository.CreateKategoriProduk(newKategoriProduk)
    if errSaveKategoriProduk != nil {
        return nil, errSaveKategoriProduk
    }

    return map[string]interface{}{
        "umkm_id": saveKategoriProduk.Umkm, // Ensure field names are correct
        "nama":   saveKategoriProduk.Nama,
    }, nil
}

func (service *KategoriProdukServiceImpl) GetKategoriProdukList(umkmID uuid.UUID) ([]entity.KategoriProdukEntity, error) {
    getKategoriProdukList, err := service.kategoriprodukrepository.GetKategoriUmkm(umkmID)
    if err != nil {
        return nil, err
    }
    return entity.ToKategoriProdukEntities(getKategoriProdukList), nil
}



