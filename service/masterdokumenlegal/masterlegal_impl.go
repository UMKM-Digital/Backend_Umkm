package masterdokumenlegalservice

import (

	"errors"
	"umkm/helper"
	"umkm/model/domain"
	"umkm/model/entity"
	"umkm/model/web"
	masterdokumenlegalrepo "umkm/repository/masterdokumenlegal"

	"github.com/google/uuid"
)

type MasterLegalServiceImpl struct {
	masterlegal masterdokumenlegalrepo.MasterDokumenLegal
}

func NewMasterLegalService(masterlegal masterdokumenlegalrepo.MasterDokumenLegal ) *MasterLegalServiceImpl {
	return &MasterLegalServiceImpl{
		masterlegal: masterlegal,
	}
}

func (service *MasterLegalServiceImpl) CreatedRequest(masterlegal web.CreateMasterDokumenLegal) (map[string]interface{}, error) {
	if masterlegal.Name == "" {
		return nil, errors.New("name cannot be empty")
	}
	newMasterLegal := domain.MasterDokumenLegal{
		Name:    masterlegal.Name,
		Iswajib: masterlegal.Is_Wajib,
	}

	saveMasterLegal, errSaveMasterLegal := service.masterlegal.Created(newMasterLegal)
	if errSaveMasterLegal != nil {
		return nil, errSaveMasterLegal
	}

	return helper.ResponseToJson{"id": saveMasterLegal.IdMasterDokumenLegal, "Name": saveMasterLegal.Name, "is_wajib": saveMasterLegal.Iswajib}, nil
}

func (service *MasterLegalServiceImpl) GetMasterLegalList(filters string, limit int, page int) ([]entity.MasterlegalEntity, int, int, int, *int, *int, error) {
	getMasterLegalList, totalcount, currentPage, totalPages, nextPage, prevPage, errGetMasterLegalList := service.masterlegal.GetmasterlegalUmkm(filters, limit, page)

	if errGetMasterLegalList != nil {
		return nil, 0, 0, 0, nil, nil, errGetMasterLegalList
	}

	masterLegalEntitis := entity.TomasterlegalEntities(getMasterLegalList)
	return masterLegalEntitis, totalcount, currentPage, totalPages,nextPage, prevPage, nil
}

func (service *MasterLegalServiceImpl) DeleteMasterLegalId(id int) error {
	return service.masterlegal.DeleteMasterLegalId(id)
}

func (service *MasterLegalServiceImpl) GetMasterLegalid(id int) (entity.MasterlegalEntity, error) {
	GetMasterLegal, errGetMasterLegal := service.masterlegal.GetMasterLegalId(id)

	if errGetMasterLegal != nil {
		return entity.MasterlegalEntity{}, errGetMasterLegal
	}

	return entity.ToMasterlegalEntity(GetMasterLegal), nil
}

func (service *MasterLegalServiceImpl) UpdateMasterLegal(request web.UpdateMasterDokumenLegal, id int) (map[string]interface{}, error) {
	getMasterLegal, errMasterLegal := service.masterlegal.GetMasterLegalId(id)
	if errMasterLegal != nil {
		return nil, errMasterLegal
	}

	if request.Name == "" {
		request.Name = getMasterLegal.Name
	}

	if request.Is_Wajib == nil {
		request.Is_Wajib = &getMasterLegal.Iswajib
	}

	masterlegalRequest := domain.MasterDokumenLegal{
		Name:       request.Name,
		Iswajib: *request.Is_Wajib,
	}

	updateMasterLeglal, errUpdate := service.masterlegal.UpdateMasterLegalId(id, masterlegalRequest)
	if errUpdate != nil {
		return nil, errUpdate
	}

	response := map[string]interface{}{"nama": updateMasterLeglal.Name, "is_wajib": updateMasterLeglal.Iswajib}
	return response, nil
}




func (service *MasterLegalServiceImpl) GetDokumenUmkmStatus(umkmId uuid.UUID, filters string, limit int, page int) ([]domain.DokumenStatusResponse, int, int, int, *int, *int, error) {
    // Panggil repository untuk mendapatkan data dokumen UMKM dengan pagination
    dokumenStatusList, totalcount, currentPage, totalPages, nextPage, prevPage, err := service.masterlegal.GetDokumenUmkmStatus(umkmId, filters, limit, page)

    if err != nil {
        return nil, 0, 0, 0, nil, nil, err
    }

    // Mengembalikan hasil dengan informasi pagination
    return dokumenStatusList, totalcount, currentPage, totalPages, nextPage, prevPage, nil
}
