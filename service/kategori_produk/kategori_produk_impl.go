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

func (service *KategoriProdukServiceImpl) GetKategoriProdukList(umkmID uuid.UUID, filters string, limit int, page int) ([]entity.KategoriProdukEntity, int, int, int, *int, *int, error) {
    getKategoriProdukList, totalcount, currentPage, totalPages, nextPage, prevPage, err := service.kategoriprodukrepository.GetKategoriProduk(umkmID, filters, limit, page)
    if err != nil {
		return nil, 0, 0, 0, nil, nil, err
    }
    
    KatgeoriEntitie := entity.ToKategoriProdukEntities(getKategoriProdukList)

	return KatgeoriEntitie, totalcount, currentPage, totalPages, nextPage, prevPage, nil
}


func (service *KategoriProdukServiceImpl) GetKategoriProdukId(id int) (entity.KategoriProdukEntity, error){
    GetKategoriProduk, errGetKategoriProduk := service.kategoriprodukrepository.GetKategoriProdukId(id)

	if errGetKategoriProduk != nil {
		return entity.KategoriProdukEntity{}, errGetKategoriProduk
	}

	return entity.ToKategoriProdukEntity(GetKategoriProduk), nil
}

func (service *KategoriProdukServiceImpl) UpdateKategoriProduk(request web.UpdateCategoriProduk, pathId int) (map[string]interface{}, error){
    getkategoriprodukId, err := service.kategoriprodukrepository.GetKategoriProdukId(pathId)
	if err != nil {
		return nil, err
	}

	if request.Name == "" {
		request.Name = getkategoriprodukId.Nama
	}

	KategoriumkmRequest := domain.KategoriProduk{
		Nama: request.Name,
	}

	updateKategoriProduk, errUpdate := service.kategoriprodukrepository.UpdateKategoriId(pathId, KategoriumkmRequest)
	if errUpdate != nil {
		return nil, errUpdate
	}

	
	response := map[string]interface{}{"name": updateKategoriProduk.Nama}
	return response, nil
}

func (service *KategoriProdukServiceImpl) DeleteKategoriProdukId(idproduk int) error{
	return service.kategoriprodukrepository.DeleteKategoriProdukId(idproduk)
}