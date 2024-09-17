package masterdokumenlegalrepo

import (
	"errors"
	"log"
	"umkm/model/domain"
	query_builder_masterlegal "umkm/query_builder/masterlegal"

	"github.com/google/uuid"
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

func (repo *MasterDokumenLegalRepoImpl) GetmasterlegalUmkm(filters string, limit int, page int) ([]domain.MasterDokumenLegal, int, int, int, *int, *int, error) {
	var masterlegal []domain.MasterDokumenLegal
	var totalcount int64

	if limit <= 0 {
		limit = 15
	}

	query, err := repo.masterlegalQuerybuilder.GetBuilderMasterLegal(filters,limit,page)
	if err != nil{
		 return nil, 0, 0, 0, nil, nil, err
	}

	err = query.Find(&masterlegal).Error
	if err != nil {
		 return nil, 0, 0, 0, nil, nil, err
	}

	masterlegalQuerybuilder, err := repo.masterlegalQuerybuilder.GetBuilderMasterLegal(filters, 0, 0)
	if err != nil {
		 return nil, 0, 0, 0, nil, nil, err
	}
	
	err = masterlegalQuerybuilder.Model(&domain.MasterDokumenLegal{}).Count(&totalcount).Error
	if err != nil {
		 return nil, 0, 0, 0, nil, nil, err
	}

	 // Hitung total pages
	 totalPages := 1
	 if limit > 0 {
		 totalPages = int((totalcount + int64(limit) - 1) / int64(limit))
	 }
 
	 // Jika page > totalPages, return kosong
	 if page > totalPages {
		 return nil, int(totalcount), page, totalPages, nil, nil, nil
	 }
 
	 currentPage := page
 
	 // Tentukan nextPage dan prevPage
	 var nextPage *int
	 if currentPage < totalPages {
		 np := currentPage + 1
		 nextPage = &np
	 }
 
	 var prevPage *int
	 if currentPage > 1 {
		 pp := currentPage - 1
		 prevPage = &pp
	 }

	return masterlegal, int(totalcount), currentPage, totalPages, nextPage, prevPage, nil
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

func (repo *MasterDokumenLegalRepoImpl) UpdateMasterLegalId(id int, dokumen domain.MasterDokumenLegal)(domain.MasterDokumenLegal, error){
	if err := repo.db.Model(&domain.MasterDokumenLegal{}).Where("id = ?", id).Updates(dokumen).Error; err != nil{
		return domain.MasterDokumenLegal{}, errors.New("failed to update master legal")
	}
	return dokumen, nil
}

type DokumenStatusResponse struct {
    Nama   string `json:"nama"`
    Status int    `json:"status"` // 0 = Not Uploaded, 1 = Uploaded
}


func (repo *MasterDokumenLegalRepoImpl) GetDokumenUmkmStatus(umkmId uuid.UUID) ([]domain.DokumenStatusResponse, error) {
    log.Printf("UMKM ID: %s", umkmId.String()) // Tambahkan log ini
    
    var results []domain.DokumenStatusResponse
    
    // Pastikan umkmId tidak kosong
    if umkmId == uuid.Nil {
        return nil, errors.New("invalid UMKM ID")
    }
    
    err := repo.db.Table("master_dokumen_legal").
        Select("master_dokumen_legal.nama, CASE WHEN umkm_dokumen_legal.id IS NOT NULL THEN 1 ELSE 0 END AS status").
        Joins("LEFT JOIN umkm_dokumen_legal ON master_dokumen_legal.id = umkm_dokumen_legal.dokumen_id AND umkm_dokumen_legal.umkm_id = ?", umkmId).
        Scan(&results).Error
    
    if err != nil {
        return nil, err
    }
    
    return results, nil
}
