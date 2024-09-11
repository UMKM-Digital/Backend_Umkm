package masterdokumenlegalservice

import (
	"errors"
	"umkm/helper"
	"umkm/model/domain"
	"umkm/model/entity"
	"umkm/model/web"
	masterdokumenlegalrepo "umkm/repository/masterdokumenlegal"
)

type MasterLegalServiceImpl struct {
	masterlegal masterdokumenlegalrepo.MasterDokumenLegal
}

func NewMasterLegalService(masterlegal masterdokumenlegalrepo.MasterDokumenLegal) *MasterLegalServiceImpl {
	return &MasterLegalServiceImpl{
		masterlegal: masterlegal,
	}
}

func (service *MasterLegalServiceImpl) CreatedRequest(masterlegal web.CreateMasterDokumenLegal) (map[string]interface{}, error){
	if masterlegal.Name == "" {
		return nil, errors.New("name cannot be empty")
	}
	newMasterLegal := domain.MasterDokumenLegal{
		Name: masterlegal.Name,
		Iswajib: masterlegal.Is_Wajib,
	}

	saveMasterLegal, errSaveMasterLegal := service.masterlegal.Created(newMasterLegal)
	if errSaveMasterLegal != nil {
		return nil, errSaveMasterLegal
	}

	return helper.ResponseToJson{"id": saveMasterLegal.IdMasterDokumenLegal, "Name": saveMasterLegal.Name, "is_wajib": saveMasterLegal.Iswajib}, nil
}

func (service *MasterLegalServiceImpl) GetMasterLegalList(filters string, limit int, page int) (map[string]interface{}, error) {
	getMasterLegalList, totalcount, errGetMasterLegalList := service.masterlegal.GetmasterlegalUmkm(filters, limit, page)

	if errGetMasterLegalList != nil {
		return nil, errGetMasterLegalList
	}

	masterLegalEntitis := entity.TomasterlegalEntities(getMasterLegalList)
	result := map[string]interface{}{
		"total_records": totalcount,
		"kategori_umkm":   masterLegalEntitis,
	}

	return result, nil
}

func (service *MasterLegalServiceImpl) DeleteMasterLegalId(id int) error {
	return service.masterlegal.DeleteMasterLegalId(id)
}

func (service *MasterLegalServiceImpl) GetMasterLegalid(id int)(entity.MasterlegalEntity, error){
	GetMasterLegal, errGetMasterLegal := service.masterlegal.GetMasterLegalId(id)

	if errGetMasterLegal != nil {
		return entity.MasterlegalEntity{}, errGetMasterLegal
	}

	return entity.ToMasterlegalEntity(GetMasterLegal), nil
}