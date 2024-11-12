package daerahrepo

import (
	domain "umkm/model/domain/master"

	"gorm.io/gorm"
)

type DaerahRepoImpl struct {
	db *gorm.DB
}

func NewDaerah(db *gorm.DB) *DaerahRepoImpl{
	return &DaerahRepoImpl{db: db}
}

func (repo *DaerahRepoImpl) GetProvinsi() ([]domain.Provinsi, error) {
	var provinsi []domain.Provinsi
	err := repo.db.Order("id ASC").Find(&provinsi).Error
	if err != nil {
		return nil, err
	}
	return provinsi, nil
}

func (repo *DaerahRepoImpl) GetKabupaten(id string) ([]domain.Kabupaten, error){
	var Kabupaten []domain.Kabupaten
	err := repo.db.Order("id ASC").Where("id_prov  = ?", id).Find(&Kabupaten).Error
	if err != nil{
		return nil, err
	}
	return Kabupaten, nil
}

func (repo *DaerahRepoImpl) GetKecamatan(id string) ([]domain.Kecamatan, error){
	var kecamatan []domain.Kecamatan
	err := repo.db.Order("id ASC").Where("id_kab  = ?", id).Find(&kecamatan).Error
	if err != nil{
		return nil, err
	}
	return kecamatan, nil
}

func (repo *DaerahRepoImpl) GetKelurahan(id string) ([]domain.Keluarahan, error){
	var kecamatan []domain.Keluarahan
	err := repo.db.Order("id ASC").Where("id_kec  = ?", id).Find(&kecamatan).Error
	if err != nil{
		return nil, err
	}
	return kecamatan, nil
}