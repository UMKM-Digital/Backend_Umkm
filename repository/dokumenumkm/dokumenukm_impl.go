package dokumenumkmrepo

import (
	"context"
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

	func(repo *UmkmDokumenImpl) DeleteDokumenUmkmId(id uuid.UUID) error{
		return repo.db.Where("umkm_id = ?", id).Delete(&domain.UmkmDokumen{}).Error
	}

	func (repo *UmkmDokumenImpl) GetDokumnByUmkmId(id uuid.UUID) ([]domain.UmkmDokumen, error) {
		var umkmDokumenList []domain.UmkmDokumen
		if err := repo.db.Where("umkm_id = ?", id).Find(&umkmDokumenList).Error; err != nil {
			return umkmDokumenList, err
		}
		return umkmDokumenList, nil
	}



	func (r *UmkmDokumenImpl) GetUmkmDokumenByUmkmIds(ctx context.Context, umkmIDs []uuid.UUID) ([]domain.UmkmDokumen, error) {
		var umkmDokumenList []domain.UmkmDokumen
	
		// Menggunakan `Where` untuk mendapatkan umkm_dokumen berdasarkan umkm_id
		err := r.db.WithContext(ctx).
			Where("umkm_id IN ?", umkmIDs).
			Find(&umkmDokumenList).Error
	
		if err != nil {
			return nil, err
		}
	
		return umkmDokumenList, nil
	}