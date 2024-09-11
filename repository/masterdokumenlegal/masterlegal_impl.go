package masterdokumenlegalrepo

import (
	"umkm/model/domain"
	query_builder_masterlegal "umkm/query_builder/masterlegal"

	"gorm.io/gorm"
)

type MasterDokumenLegalRepoImpl struct {
	db *gorm.DB
	masterlegalQuerybuilder query_builder_masterlegal.MasteLegalQueryBuilder
}

func NewDokumenLegalRepoImpl(db *gorm.DB, masterlegalQuerybuilder query_builder_masterlegal.MasteLegalQueryBuilder) *MasterDokumenLegalRepoImpl {
	return &MasterDokumenLegalRepoImpl{
		db:                 db,
		masterlegalQuerybuilder: masterlegalQuerybuilder,
	}
}

func (repo *MasterDokumenLegalRepoImpl) Created(dokumen domain.MasterDokumenLegal) (domain.MasterDokumenLegal, error) {
	err := repo.db.Create(&dokumen).Error
	if err != nil {
		return domain.MasterDokumenLegal{}, err
	}

	return dokumen, nil
}

func (repo *MasterDokumenLegalRepoImpl) GetmasterlegalUmkm(filters string, limit int, page int) ([]domain.MasterDokumenLegal, int, error) {
	var masterlegal []domain.MasterDokumenLegal
	var totalcount int64

	query, err := repo.masterlegalQuerybuilder.GetBuilderMasterLegal(filters,limit,page)
	if err != nil{
		return nil, 0, err
	}

	err = query.Find(&masterlegal).Error
	if err != nil {
		return nil, 0, err
	}

	masterlegalQuerybuilder, err := repo.masterlegalQuerybuilder.GetBuilderMasterLegal(filters, 0, 0)
	if err != nil {
		return nil, 0, err
	}
	
	err = masterlegalQuerybuilder.Model(&domain.MasterDokumenLegal{}).Count(&totalcount).Error
	if err != nil {
		return nil, 0, err
	}

	return masterlegal, int(totalcount), nil
}

func (repo *MasterDokumenLegalRepoImpl) DeleteMasterLegalId(id int) error {
    if err := repo.db.Delete(&domain.MasterDokumenLegal{}, id).Error; err != nil {
        return err
    }
    return nil
}

func (repo *MasterDokumenLegalRepoImpl) GetMasterLegalId(id int)(domain.MasterDokumenLegal, error){
	var masterdokumenlegal domain.MasterDokumenLegal
	if err := repo.db.First(&masterdokumenlegal, "id = ?", id).Error; err != nil {
		return masterdokumenlegal, err
	}
	return masterdokumenlegal, nil
}