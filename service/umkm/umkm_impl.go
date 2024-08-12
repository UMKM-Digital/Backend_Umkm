package umkmservice

import (
	"context"
	"encoding/json"
	"errors"

	// "fmt"
	// "os/user"
	"umkm/helper"
	"umkm/model/domain"
	"umkm/model/entity"
	"umkm/model/web"
	hakaksesrepo "umkm/repository/hakakses" // Tambahkan import untuk HakAkses repository
	umkmrepo "umkm/repository/umkm"

	"github.com/google/uuid"
)

type UmkmServiceImpl struct {
	umkmrepository    umkmrepo.CreateUmkm
	hakaksesrepository hakaksesrepo.CreateHakakses // Tambahkan field untuk HakAkses repository
}

func NewUmkmService(umkmrepository umkmrepo.CreateUmkm, hakaksesrepository hakaksesrepo.CreateHakakses) *UmkmServiceImpl {
	return &UmkmServiceImpl{
		umkmrepository:    umkmrepository,
		hakaksesrepository: hakaksesrepository,
	}
}

// func (service *UmkmServiceImpl) CreateUmkm(umkm web.UmkmRequest) (map[string]interface{}, error) {
// 	kategoriUmkmId, err := helper.RawMessageToJSONB(umkm.Kategori_Umkm_Id)
// 	if err != nil {
// 		return nil, errors.New("invalid type for Kategori_Umkm_Id")
// 	}

// 	informasiJamBuka, err := helper.RawMessageToJSONB(umkm.Informasi_JamBuka)
// 	if err != nil {
// 		return nil, errors.New("invalid type for Informasi_JamBuka")
// 	}

// 	Maps, err := helper.RawMessageToJSONB(umkm.Maps)
// 	if err != nil {
// 		return nil, errors.New("invalid type for Maps")
// 	}

// 	var Images domain.JSONB
// 	if len(umkm.Gambar) > 0 {
// 		var imgURLs []string
// 		if err := json.Unmarshal(umkm.Gambar, &imgURLs); err != nil {
// 			return nil, errors.New("invalid type for Images")
// 		}
// 		Images = domain.JSONB{"urls": imgURLs}
// 	} else {
// 		Images = domain.JSONB{"urls": []string{}} // Ensure Images is not null
// 	}

// 	// Log converted values for debugging
// 	fmt.Printf("KategoriUmkmId: %+v\n", kategoriUmkmId)
// 	fmt.Printf("InformasiJamBuka: %+v\n", informasiJamBuka)
// 	fmt.Printf("Maps: %+v\n", Maps)
// 	fmt.Printf("Images: %+v\n", Images)

// 	newUmkm := domain.UMKM{
// 		Name:                umkm.Name,
// 		NoNpwp:              umkm.NoNpwp,
// 		Images:              Images,
// 		KategoriUmkmId:      kategoriUmkmId,
// 		NamaPenanggungJawab: umkm.Nama_Penanggung_Jawab,
// 		InformasiJambuka:    informasiJamBuka,
// 		NoKontak:            umkm.No_Kontak,
// 		Lokasi:              umkm.Lokasi,
// 		Maps:                Maps,
// 	}

// 	// Save UMKM
// 	saveUmkm, errSaveUmkm := service.umkmrepository.CreateRequest(newUmkm)
// 	if errSaveUmkm != nil {
// 		return nil, errSaveUmkm
// 	}

// 	// Create HakAkses
// 	hakAkses := domain.HakAkses{
// 		UserId: umkm.UserId, // Ensure UserId is part of the request
// 		UmkmId: saveUmkm.IdUmkm,
// 		Status: 1, // Example status
// 	}
// 	if err := service.hakaksesrepository.CreateHakAkses(&hakAkses); err != nil {
// 		return nil, err
// 	}

// 	return map[string]interface{}{
// 		"Name":                  saveUmkm.Name,
// 		"KategoriUmkmId":        saveUmkm.KategoriUmkmId,
// 		"Nama Penanggung Jawab": saveUmkm.NamaPenanggungJawab,
// 		"Informasi Jam":         saveUmkm.InformasiJambuka,
// 		"No Kontak":             saveUmkm.NoKontak,
// 		"Lokasi":                saveUmkm.Lokasi,
// 		"Images":                saveUmkm.Images,
// 	}, nil
// }


func (service *UmkmServiceImpl) CreateUmkm(umkm web.UmkmRequest, userID int) (map[string]interface{}, error) {
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
        Images = domain.JSONB{"urls": []string{}}
    }

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

    hakAkses := domain.HakAkses{
        UserId: userID,
        UmkmId: saveUmkm.IdUmkm,
        Status: 0,
    }
    if err := service.hakaksesrepository.CreateHakAkses(&hakAkses); err != nil {
        return nil, err
    }

    return map[string]interface{}{
        "name":                  saveUmkm.Name,
        "kategori_umkm_id":        saveUmkm.KategoriUmkmId,
        "nama_penanggung_jawab": saveUmkm.NamaPenanggungJawab,
        "informasi_jam":         saveUmkm.InformasiJambuka,
        "no_kontak":             saveUmkm.NoKontak,
        "lokasi":                saveUmkm.Lokasi,
        "images":                saveUmkm.Images,
		"user_id":                 userID,
		"status":            hakAkses.Status,
    }, nil
}

func (s *UmkmServiceImpl) GetUmkmListByUserId(ctx context.Context, userId int) ([]entity.UmkmEntity, error) {
	// Fetch HakAkses for the given user ID
	hakAksesList, err := s.hakaksesrepository.GetHakAksesByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	// Collect UMKM IDs from HakAkses
	var umkmIDs []uuid.UUID
	for _, hakAkses := range hakAksesList {
		umkmIDs = append(umkmIDs, hakAkses.UmkmId)
	}

	// Fetch UMKM entities based on the collected IDs
	umkmList, err := s.umkmrepository.GetUmkmListByIds(ctx, umkmIDs)
	if err != nil {
		return nil, err
	}

	// Convert UMKM entities to UmkmEntity format
	return entity.ToUmkmEntities(umkmList), nil
}