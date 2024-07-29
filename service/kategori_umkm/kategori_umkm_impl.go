package kategoriumkmservice

import (
	"umkm/helper"
	"umkm/model/domain"
	"umkm/model/entity"
	"umkm/model/web"
	repokategoriumkm "umkm/repository/kategori_umkm"
)

type KategoriUmkmServiceImpl struct {
	kategorirepository repokategoriumkm.CreateCategoryUmkm
}

func NewKategoriUmkmService(kategorirepository repokategoriumkm.CreateCategoryUmkm) *KategoriUmkmServiceImpl {
	return &KategoriUmkmServiceImpl{
		kategorirepository: kategorirepository,
	}
}
//post katgeori
func (service *KategoriUmkmServiceImpl) CreateKategori(kategori web.CreateCategoriUmkm) (map[string]interface{}, error) {
	newKategori := domain.Kategori_Umkm{
		Name: kategori.Name,
	}

	saveKategoriUmkm, errSaveKategoriUmkm := service.kategorirepository.CreateRequest(newKategori)
	if errSaveKategoriUmkm != nil {
		return nil, errSaveKategoriUmkm
	}

	return helper.ResponseToJson{"Name": saveKategoriUmkm.Name}, nil
}

//baca seluruh kategori
func (service *KategoriUmkmServiceImpl) GetKategoriUmkmList() ([]entity.KategoriEntity, error) {
    getKategoriUmkmList, errGetKategoriUmkmList := service.kategorirepository.GetKategoriUmkm()

    if errGetKategoriUmkmList != nil {
        return []entity.KategoriEntity{}, errGetKategoriUmkmList
    }

    return entity.ToKategoriEntities(getKategoriUmkmList), nil
}

//get by id
func (service *KategoriUmkmServiceImpl) GetKategoriUmkmId(id int) (entity.KategoriEntity, error) {
	GetKategoriUmkm, errGetKategGetKategoriUmkm := service.kategorirepository.GetKategoriUmkmId(id)

	if errGetKategGetKategoriUmkm != nil {
		return entity.KategoriEntity{}, errGetKategGetKategoriUmkm
	}

	return entity.ToKategoriEntity(GetKategoriUmkm), nil
}

//update
func (service *KategoriUmkmServiceImpl) UpdateKategori(request web.UpdateCategoriUmkm, pathId int) (map[string]interface{}, error) {
	// Ambil data kategori berdasarkan ID
	getKategoriById, err := service.kategorirepository.GetKategoriUmkmId(pathId)
	if err != nil {
		return nil, err
	}

	// Gunakan nama kategori yang ada jika tidak ada perubahan
	if request.Name == "" {
		request.Name = getKategoriById.Name
	}

	// Buat objek Kategori_Umkm baru untuk pembaruan
	KategoriumkmRequest := domain.Kategori_Umkm{
		IdKategori: pathId,
		Name:      request.Name,
	}

	// Update kategori UMKM
	updateKategoriUmkm, errUpdate := service.kategorirepository.UpdateKategoriId(pathId, KategoriumkmRequest)
	if errUpdate != nil {
		return nil, errUpdate
	}

	// Membentuk respons yang akan dikembalikan
	response := map[string]interface{}{"name": updateKategoriUmkm.Name}
	return response, nil
}

//delete
func (service *KategoriUmkmServiceImpl) DeleteKategoriUmkmId(id int) error {
    return service.kategorirepository.DeleteKategoriUmkmId(id)
}

