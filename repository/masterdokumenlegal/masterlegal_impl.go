package masterdokumenlegalrepo

import (
	"context"
	"errors"
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

	err = query.Order("id ASC").Find(&masterlegal).Error
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


func (repo *MasterDokumenLegalRepoImpl) GetDokumenUmkmStatus(umkmId uuid.UUID, filters string, limit int, page int) ([]domain.DokumenStatusResponse, int, int, int, *int, *int, error) {
    var results []domain.DokumenStatusResponse
    var totalcount int64

    if limit <= 0 {
        limit = 15
    }
    
    // Pastikan umkmId tidak kosong
    if umkmId == uuid.Nil {
        return nil, 0, 0, 0, nil, nil, errors.New("invalid UMKM ID")
    }

    // Gunakan query builder dengan pagination
    query, err := repo.masterlegalQuerybuilder.GetBuilderDokumenUmkmStatus(umkmId, filters, limit, page)
    if err != nil {
        return nil, 0, 0, 0, nil, nil, err
    }

    // Eksekusi query dan scan hasilnya
    err = query.Find(&results).Error
    if err != nil {
        return nil, 0, 0, 0, nil, nil, err
    }

    // Hitung total count
    countQuery, err := repo.masterlegalQuerybuilder.GetBuilderDokumenUmkmStatus(umkmId, filters, 0, 0)
    if err != nil {
        return nil, 0, 0, 0, nil, nil, err
    }

    err = countQuery.Model(&domain.DokumenStatusResponse{}).Count(&totalcount).Error
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

    return results, int(totalcount), currentPage, totalPages, nextPage, prevPage, nil
}


func (r *MasterDokumenLegalRepoImpl) GetAllMasterDokumenLegal(ctx context.Context) ([]domain.MasterDokumenLegal, error) {
    var masterDokumenLegalList []domain.MasterDokumenLegal

    // Menggunakan Gorm untuk mendapatkan semua data master_dokumen_legal
    err := r.db.WithContext(ctx).
        Find(&masterDokumenLegalList).Error

    if err != nil {
        return nil, err
    }

    return masterDokumenLegalList, nil
}


// //all
// func (repo *MasterDokumenLegalRepoImpl) GetDokumenUmkmStatusAll(userId int, filters string, limit int, page int) ([]domain.UmkmDocumentsResponse, int, int, int, *int, *int, error) {
// 	var results []domain.DokumenStatusResponseALL
// 	var totalCount int64

// 	if limit <= 0 {
// 		limit = 100
// 	}

// 	query, err := repo.masterlegalQuerybuilder.GetBuilderDokumenUmkmStatusAll(userId, filters, limit, page)
// 	if err != nil {
// 		return nil, 0, 0, 0, nil, nil, err
// 	}

// 	if err = query.Find(&results).Error; err != nil {
// 		return nil, 0, 0, 0, nil, nil, err
// 	}

// 	// Hitung total dokumen
// 	countQuery, err := repo.masterlegalQuerybuilder.GetBuilderDokumenUmkmStatusAll(userId, filters, 0, 0)
// 	if err != nil {
// 		return nil, 0, 0, 0, nil, nil, err
// 	}

// 	if err = countQuery.Count(&totalCount).Error; err != nil {
// 		return nil, 0, 0, 0, nil, nil, err
// 	}

// 	totalPages := int((totalCount + int64(limit) - 1) / int64(limit))
// 	if page > totalPages {
// 		return nil, int(totalCount), page, totalPages, nil, nil, nil
// 	}

// 	currentPage := page
// 	var nextPage *int
// 	if currentPage < totalPages {
// 		np := currentPage + 1
// 		nextPage = &np
// 	}

// 	var prevPage *int
// 	if currentPage > 1 {
// 		pp := currentPage - 1
// 		prevPage = &pp
// 	}

// 	// Mengelompokkan hasil berdasarkan umkm_id
// 	umkmMap := make(map[uuid.UUID][]domain.DokumenStatusResponseALL)
// 	for _, result := range results {
// 		umkmMap[result.UmkmId] = append(umkmMap[result.UmkmId], result)
// 	}

// 	// Membangun respons akhir
// 	var umkmDocumentsResponses []domain.UmkmDocumentsResponse
// 	for umkmID, docs := range umkmMap {
// 		umkmDocumentsResponses = append(umkmDocumentsResponses, domain.UmkmDocumentsResponse{
// 			UmkmID:  umkmID,
// 			Dokumen: docs,
// 		})
// 	}

// 	return umkmDocumentsResponses, int(totalCount), currentPage, totalPages, nextPage, prevPage, nil
// }


func (repo *MasterDokumenLegalRepoImpl) GetDokumenUmkmStatusAll(userId int, filters string, limit int, page int) ([]domain.DokumenStatusResponseALL, int, int, int, *int, *int, error) {
	var results []domain.DokumenStatusResponseALL
	var totalCount int64

	if limit <= 0 {
		limit = 100
	}

	// Fetch documents with status
	query, err := repo.masterlegalQuerybuilder.GetBuilderDokumenUmkmStatusAll(userId, filters, limit, page)
	if err != nil {
		return nil, 0, 0, 0, nil, nil, err
	}

	if err = query.Find(&results).Error; err != nil {
		return nil, 0, 0, 0, nil, nil, err
	}

	// Get total count of documents
	countQuery, err := repo.masterlegalQuerybuilder.GetBuilderDokumenUmkmStatusAll(userId, filters, 0, 0)
	if err != nil {
		return nil, 0, 0, 0, nil, nil, err
	}

	if err = countQuery.Count(&totalCount).Error; err != nil {
		return nil, 0, 0, 0, nil, nil, err
	}

	totalPages := int((totalCount + int64(limit) - 1) / int64(limit))
	if page > totalPages {
		return nil, int(totalCount), page, totalPages, nil, nil, nil
	}

	currentPage := page
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

	return results, int(totalCount), currentPage, totalPages, nextPage, prevPage, nil
}
