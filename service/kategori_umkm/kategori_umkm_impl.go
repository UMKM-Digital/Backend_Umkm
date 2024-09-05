package kategoriumkmservice

import (
	"errors"
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

// post katgeori
func (service *KategoriUmkmServiceImpl) CreateKategori(kategori web.CreateCategoriUmkm) (map[string]interface{}, error) {
	if kategori.Name == "" {
		return nil, errors.New("kategori name cannot be empty")
	}

	newKategori := domain.Kategori_Umkm{
		Name: kategori.Name,
	}

	saveKategoriUmkm, errSaveKategoriUmkm := service.kategorirepository.CreateRequest(newKategori)
	if errSaveKategoriUmkm != nil {
		return nil, errSaveKategoriUmkm
	}

	return helper.ResponseToJson{"id": saveKategoriUmkm.IdKategori, "Name": saveKategoriUmkm.Name}, nil
}

// baca seluruh kategori
func (service *KategoriUmkmServiceImpl) GetKategoriUmkmList(filters string, limit int, page int) (map[string]interface{}, error) {
	getKategoriUmkmList, totalcount, errGetKategoriUmkmList := service.kategorirepository.GetKategoriUmkm(filters, limit, page)

	if errGetKategoriUmkmList != nil {
		return nil, errGetKategoriUmkmList
	}

	kategoriUmkmEntitis := entity.ToKategoriEntities(getKategoriUmkmList)
	result := map[string]interface{}{
		"total_records": totalcount,
		"kategori_umkm":   kategoriUmkmEntitis,
	}

	return result, nil
}

// get by id
func (service *KategoriUmkmServiceImpl) GetKategoriUmkmId(id int) (entity.KategoriEntity, error) {
	GetKategoriUmkm, errGetKategGetKategoriUmkm := service.kategorirepository.GetKategoriUmkmId(id)

	if errGetKategGetKategoriUmkm != nil {
		return entity.KategoriEntity{}, errGetKategGetKategoriUmkm
	}

	return entity.ToKategoriEntity(GetKategoriUmkm), nil
}

// update
func (service *KategoriUmkmServiceImpl) UpdateKategori(request web.UpdateCategoriUmkm, pathId int) (map[string]interface{}, error) {
	getKategoriById, err := service.kategorirepository.GetKategoriUmkmId(pathId)
	if err != nil {
		return nil, err
	}

	if request.Name == "" {
		request.Name = getKategoriById.Name
	}

	
	KategoriumkmRequest := domain.Kategori_Umkm{
		Name:       request.Name,
	}

	
	updateKategoriUmkm, errUpdate := service.kategorirepository.UpdateKategoriId(pathId, KategoriumkmRequest)
	if errUpdate != nil {
		return nil, errUpdate
	}

	
	response := map[string]interface{}{"name": updateKategoriUmkm.Name}
	return response, nil
}

// delete
func (service *KategoriUmkmServiceImpl) DeleteKategoriUmkmId(id int) error {
	return service.kategorirepository.DeleteKategoriUmkmId(id)
}
