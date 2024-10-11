package daerahservice

import (
	entity "umkm/model/entity/master"
	daerahrepo "umkm/repository/daerah"
)

type DaerahServiceImpl struct {
	daerahrepo daerahrepo.Daerah
}

func NewDaerahService(daerahrepo daerahrepo.Daerah) *DaerahServiceImpl {
	return &DaerahServiceImpl{
		daerahrepo: daerahrepo,
	}
}

func (service *DaerahServiceImpl) GetDaerah() ([]entity.DaerahEntity, error) {
	GetDaerahList, err := service.daerahrepo.GetProvinsi()
	if err != nil {
		return nil, err
	}
	return entity.ToDaerahEntities(GetDaerahList), nil
}

func (service *DaerahServiceImpl) GetKabupaten(id string) ([]entity.KabupatenEntity, error){
	GetKabupatenList, err := service.daerahrepo.GetKabupaten(id)
	if err != nil {
		return nil, err
	}
	return entity.ToDaerahKabupatenEntities(GetKabupatenList), nil
}

func (service *DaerahServiceImpl) GetKecamatan(id string) ([]entity.KecamatanEntity,error){
	GetKecamatanList, err := service.daerahrepo.GetKecamatan(id)
	if err != nil {
		return nil, err
	}
	return entity.ToDaerahKecamatanEntities(GetKecamatanList), nil
}

func (service *DaerahServiceImpl) GetKelurahan(id string) ([]entity.KeluarahanEntity, error){
	GetKelurahanList, err := service.daerahrepo.GetKelurahan(id)
	if err != nil {
		return nil, err
	}
	return entity.ToDaerahKeluarahanEntities(GetKelurahanList), nil
}