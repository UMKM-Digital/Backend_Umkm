package sektorusahaservice

import (
	domain "umkm/model/domain/master"
		entity "umkm/model/entity/master"
	web "umkm/model/web/master"
	sektorusaharepo "umkm/repository/sektorusaha"

)

type SektorUsahaServiceImpl struct {
	sektorusaharepo sektorusaharepo.SektorUsaha
}

func NewSektorUsahaService(sektorusaharepo sektorusaharepo.SektorUsaha) *SektorUsahaServiceImpl {
	return &SektorUsahaServiceImpl{
		sektorusaharepo: sektorusaharepo,
	}
}

func (service *SektorUsahaServiceImpl) CreateSektorUsaha(sektorusaha web.CreateSektorUsaha) (map[string]interface{}, error) {
	newSektorUsaha := domain.SektorUsaha{
		Nama: sektorusaha.Name,
	}

	saveSektorUsaha, errSaveSektorUsaha := service.sektorusaharepo.CreateSektorUsaha(newSektorUsaha)
	if errSaveSektorUsaha != nil {
		return nil, errSaveSektorUsaha
	}

	return map[string]interface{}{
		"nama":    saveSektorUsaha.Nama,
	}, nil
}

func (service *SektorUsahaServiceImpl) 	GetSektorUsaha() ([]entity.SektorUsahaEntity, error) {
	getKategoriProdukList, err := service.sektorusaharepo.GetSektorUsaha()
	if err != nil {
		return nil, err
	}

	KatgeoriEntitie := entity.ToKategoriEntities(getKategoriProdukList)

	return KatgeoriEntitie,  nil
}
//bentukusaha
func (service *SektorUsahaServiceImpl) 	GetBentukUsaha() ([]entity.BentukUsahaEntity, error) {
	getBentukUsahaList, err := service.sektorusaharepo.GetBentukUsaha()
	if err != nil {
		return nil, err
	}

	bentukEntitie := entity.ToBentukUsahaEntities(getBentukUsahaList)

	return bentukEntitie,  nil
}
//statustempatusaha
func (service *SektorUsahaServiceImpl) 	GetStatusTempatUsaha() ([]entity.StatusTempatUsahaEntity, error) {
	getKategoriProdukList, err := service.sektorusaharepo.GetStatusTempatUsaha()
	if err != nil {
		return nil, err
	}

	KatgeoriEntitie := entity.ToStatusTempatUsahaEntities(getKategoriProdukList)

	return KatgeoriEntitie,  nil
}
