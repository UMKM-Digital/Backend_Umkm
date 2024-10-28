package omsetservice

import (
	"umkm/helper"
	"umkm/model/domain"
	"umkm/model/entity"
	"umkm/model/web"
	omsetrepo "umkm/repository/omset"

	"github.com/google/uuid"
)

type OmzetServiceImpl struct {
	omsetrepository omsetrepo.OmsetRepo
}

func NewOmsetService(omsetrepo omsetrepo.OmsetRepo) *OmzetServiceImpl{
	return &OmzetServiceImpl{
		omsetrepository: omsetrepo,
	}
}


func(service *OmzetServiceImpl) CreateOmsetService( omset web.Omset)(map[string]interface{}, error){
	newOmset := domain.Omset{
		UmkmId: omset.UmkmId,
		Nominal: omset.JumlahOmset,
		Bulan: omset.Bulan,
	}

	saveKategoriUmkm, errSaveKategoriUmkm := service.omsetrepository.CreateRequest(newOmset)
	if errSaveKategoriUmkm != nil {
		return nil, errSaveKategoriUmkm
	}

	return helper.ResponseToJson{"umkm_id": saveKategoriUmkm.UmkmId, "nominal": saveKategoriUmkm.Nominal, "bulan": saveKategoriUmkm.Bulan}, nil
}

func(service *OmzetServiceImpl) ListOmsetService(umkm_id uuid.UUID, tahun string)([]entity.OmsetEntity, error){
	getKategoriUmkmList, err := service.omsetrepository.ListOmsetRequest(umkm_id, tahun)

	if err != nil {
        return nil, err
    }

	return entity.ToOmsetListEntities(getKategoriUmkmList), nil
}

func(service *OmzetServiceImpl) GetOmsetServiceId(id int) (entity.OmsetEntity, error){
	GetKategoriUmkm, errGetKategGetKategoriUmkm := service.omsetrepository.GetOmsetId(id)

	if errGetKategGetKategoriUmkm != nil {
		return entity.OmsetEntity{}, errGetKategGetKategoriUmkm
	}

	return entity.ToOmsetEntityList(GetKategoriUmkm), nil
}


//update
func (service *OmzetServiceImpl) UpdateOmset(request web.UpdateOmset, pathId int) (map[string]interface{}, error) {
	getKategoriById, err := service.omsetrepository.GetOmsetId(pathId)
	if err != nil {
		return nil, err
	}

	// Gunakan nilai request jika tersedia, jika tidak gunakan nilai yang ada di database
	nominal := request.Nominal
	if nominal.IsZero() {
		nominal = getKategoriById.Nominal
	}

	bulan := request.Bulan
	if bulan == "" {
		bulan = getKategoriById.Bulan
	}

	// Siapkan data yang akan diperbarui
	KategoriumkmRequest := domain.Omset{
		Nominal: nominal,
		Bulan:   bulan,
	}

	// Lakukan update melalui repository
	updateKategoriUmkm, errUpdate := service.omsetrepository.UpdateOmsetId(pathId, KategoriumkmRequest)
	if errUpdate != nil {
		return nil, errUpdate
	}

	response := map[string]interface{}{
		"Nominal": updateKategoriUmkm.Nominal,
		"Bulan":   updateKategoriUmkm.Bulan,
	}

	return response, nil
}