package dokumenumkmrepo

import (
	"errors"
	"umkm/model/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UmkmDokumenImpl struct {
	db *gorm.DB
}

func NewDokumenRepositoryImpl(db *gorm.DB) *UmkmDokumenImpl {
	return &UmkmDokumenImpl{
		db: db,
	}
}

func (repo *UmkmDokumenImpl) CreateRequest(dokumenumkm domain.UmkmDokumen) (domain.UmkmDokumen, error) {
	err := repo.db.Create(&dokumenumkm).Error
	if err != nil {
		return domain.UmkmDokumen{}, err
	}

	return dokumenumkm, nil
}

func(repo *UmkmDokumenImpl) GetId(id int, umkmid uuid.UUID) (domain.UmkmDokumen, error){
	var dokumenumkm domain.UmkmDokumen

	err := repo.db.First(&dokumenumkm,"dokumen_id = ? AND umkm_id = ?",id, umkmid).Error

	if err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			return domain.UmkmDokumen{}, errors.New("dokumen umkm not found")
		}
		return domain.UmkmDokumen{}, err
	}

	return dokumenumkm, nil
}

func(repo *UmkmDokumenImpl) UpdateDokumen(id int, umkmid uuid.UUID, dokumenumkm domain.UmkmDokumen)(domain.UmkmDokumen, error){
	if err := repo.db.Model(&domain.UmkmDokumen{}).Where("dokumen_id = ? AND umkm_id = ?",id,umkmid).Updates(dokumenumkm).Error;err != nil{
		return domain.UmkmDokumen{}, err
	}
	return dokumenumkm, nil
}