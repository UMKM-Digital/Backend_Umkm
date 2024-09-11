package masterdokumenlegalservice

import (
	"errors"
	"umkm/helper"
	"umkm/model/domain"
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

func (service MasterLegalServiceImpl) CreatedRequest(masterlegal web.CreateMasterDokumenLegal) (map[string]interface{}, error){
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