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


