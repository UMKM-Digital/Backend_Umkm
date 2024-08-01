package umkmservice

import (
	"errors"
	"umkm/helper"
	"umkm/model/domain"
	"umkm/model/web"
	umkmrepo "umkm/repository/umkm"
)

type UmkmServiceImpl struct {
	umkmrepository umkmrepo.CreateUmkm
}

func NewUmkmService(umkmrepository umkmrepo.CreateUmkm) *UmkmServiceImpl {
	return &UmkmServiceImpl{
		umkmrepository: umkmrepository,
	}
}

func (service *UmkmServiceImpl) CreateUmkm(umkm web.UmkmRequest) (map[string]interface{}, error) {
	// Convert json.RawMessage to domain.JSONB
	gambar, err := helper.RawMessageToJSONB(umkm.Gambar)
	if err != nil {
		return nil, errors.New("invalid type for Gambar")
	}

	kategoriUmkmId, err := helper.RawMessageToJSONB(umkm.Kategori_Umkm_Id)
	if err != nil {
		return nil, errors.New("invalid type for Kategori_Umkm_Id")
	}

	informasiJamBuka, err := helper.RawMessageToJSONB(umkm.Informasi_JamBuka)
	if err != nil {
		return nil, errors.New("invalid type for Informasi_JamBuka")
	}

	Maps, err := helper.RawMessageToJSONB(umkm.Maps)
	if err != nil {
		return nil, errors.New("invalid type for Maps")
	}

	// Create new UMKM instance with converted values
	newUmkm := domain.UMKM{
		Name:                   umkm.Name,
		NoNpwp:                 umkm.NoNpwp,
		Gambar:                 gambar,
		KategoriUmkmId:         kategoriUmkmId,
		NamaPenanggungJawab:    umkm.Nama_Penanggung_Jawab,
		InformasiJambuka:       informasiJamBuka,
		NoKontak:               umkm.No_Kontak,
		Lokasi:                 umkm.Lokasi,
		Maps:                   Maps,
	}

	saveUmkm, errSaveUmkm := service.umkmrepository.CreateRequest(newUmkm)
	if errSaveUmkm != nil {
		return nil, errSaveUmkm
	}

	return map[string]interface{}{"Name": saveUmkm.Name, "Gambar":saveUmkm.Gambar, "KategoriUmkmId": saveUmkm.KategoriUmkmId, "Nama Penangungg Jawab":saveUmkm.NamaPenanggungJawab,
	"informasi jam": saveUmkm.InformasiJambuka, "no kontak": saveUmkm.NoKontak, "lokasi": saveUmkm.Lokasi }, nil
}
