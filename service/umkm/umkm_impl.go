package umkmservice

import (
	"encoding/json"
	"errors"
	"fmt"
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

	var Images domain.JSONB
    if len(umkm.Gambar) > 0 {
        var imgURLs []string
        if err := json.Unmarshal(umkm.Gambar, &imgURLs); err != nil {
            return nil, errors.New("invalid type for Images")
        }
        Images = domain.JSONB{"urls": imgURLs}
    } else {
        Images = domain.JSONB{"urls": []string{}} // Ensure Images is not null
    }

	// Log converted values for debugging
	fmt.Printf("KategoriUmkmId: %+v\n", kategoriUmkmId)
	fmt.Printf("InformasiJamBuka: %+v\n", informasiJamBuka)
	fmt.Printf("Maps: %+v\n", Maps)
	fmt.Printf("Images: %+v\n", Images)

	newUmkm := domain.UMKM{
		Name:                umkm.Name,
		NoNpwp:              umkm.NoNpwp,
		Images:              Images,
		KategoriUmkmId:      kategoriUmkmId,
		NamaPenanggungJawab: umkm.Nama_Penanggung_Jawab,
		InformasiJambuka:    informasiJamBuka,
		NoKontak:            umkm.No_Kontak,
		Lokasi:              umkm.Lokasi,
		Maps:                Maps,
	}

	saveUmkm, errSaveUmkm := service.umkmrepository.CreateRequest(newUmkm)
	if errSaveUmkm != nil {
		return nil, errSaveUmkm
	}

	return map[string]interface{}{
		"Name":                  saveUmkm.Name,
		"KategoriUmkmId":        saveUmkm.KategoriUmkmId,
		"Nama Penanggung Jawab": saveUmkm.NamaPenanggungJawab,
		"Informasi Jam":         saveUmkm.InformasiJambuka,
		"No Kontak":             saveUmkm.NoKontak,
		"Lokasi":                saveUmkm.Lokasi,
		"Images":                saveUmkm.Images,
	}, nil
}
