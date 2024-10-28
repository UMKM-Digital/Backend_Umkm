package omsetrepo

import (
	"errors"
	"umkm/model/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OmmsetRepoImpl struct {
	db                 *gorm.DB
	
}

func NewomsetRepositoryImpl(db *gorm.DB) *OmmsetRepoImpl {
	return &OmmsetRepoImpl{
		db:                 db,
	}
}

func (repo *OmmsetRepoImpl) CreateRequest(omset domain.Omset) (domain.Omset, error) {
	err := repo.db.Create(&omset).Error
	if err != nil {
		return domain.Omset{}, err
	}

	return omset, nil
}

//list
func(repo *OmmsetRepoImpl) ListOmsetRequest(umkm_id uuid.UUID, tahun string)([]domain.Omset, error){
	var omset []domain.Omset

	err := repo.db.Where("umkm_id = ? AND SUBSTRING(bulan FROM 1 FOR 4) = ?", umkm_id, tahun).Find(&omset).Error

	if err != nil {
		return []domain.Omset{}, err
	}

	return omset, nil
}

func(repo *OmmsetRepoImpl) GetOmsetId(id int)(domain.Omset, error){
	var OmsetUmkm domain.Omset

	err := repo.db.First(&OmsetUmkm, "id = ?", id).Error

	if err != nil {
		return domain.Omset{}, errors.New("omset tidak ditemukan")
	}

	return OmsetUmkm, nil
}


//update

func (repo *OmmsetRepoImpl) UpdateOmsetId(id int, omset domain.Omset) (domain.Omset, error) {
    if err := repo.db.Debug().Model(&domain.Omset{}).Where("id = ?", id).Updates(omset).Error; err != nil {
        return domain.Omset{}, errors.New("failed to update omset")
    }
    return omset, nil
}