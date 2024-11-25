package omsetrepo

import (
	"errors"
	"fmt"
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

//grafik omset mobile

func (repo *OmmsetRepoImpl) OmsetTahunan(umkm_id uuid.UUID, tahun string) (float64, error) {
	var totalNominal float64

	// Menggunakan COALESCE untuk menghindari NULL pada hasil SUM
	err := repo.db.Table("omzets").
		Where("umkm_id = ? AND SUBSTRING(bulan FROM 1 FOR 4) = ?", umkm_id, tahun).
		Select("COALESCE(SUM(nominal), 0)"). // Mengatasi NULL dengan COALESCE
		Scan(&totalNominal).Error

	if err != nil {
		return 0, err
	}

	return totalNominal, nil
}



func (repo *OmmsetRepoImpl) OmsetBulanan(umkm_id uuid.UUID, tahun string) (map[string]float64, error) {
	// Inisialisasi map untuk menyimpan total omzet per bulan
	totalOmzetPerBulan := make(map[string]float64)

	// Inisialisasi map dengan bulan dari Januari hingga Desember untuk tahun yang diberikan
	for month := 1; month <= 12; month++ {
		// Menggunakan fmt.Sprintf untuk format bulan menjadi MM
		bulan := fmt.Sprintf("%s-%02d", tahun, month)
		totalOmzetPerBulan[bulan] = 0
	}

	// Menyimpan total omzet per bulan dari Januari hingga Desember
	var results []struct {
		Bulan       string
		TotalNominal float64
	}

	// Query untuk menghitung total nominal per bulan dalam tahun tertentu
	err := repo.db.Table("omzets").
		Select("CONCAT(SUBSTRING(bulan FROM 1 FOR 4), '-', SUBSTRING(bulan FROM 6 FOR 2)) AS bulan, SUM(nominal) AS total_nominal").
		Where("umkm_id = ? AND SUBSTRING(bulan FROM 1 FOR 4) = ?", umkm_id, tahun).
		Group("bulan").
		Order("bulan").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	// Update map dengan hasil query
	for _, result := range results {
		totalOmzetPerBulan[result.Bulan] = result.TotalNominal
	}

	return totalOmzetPerBulan, nil
}

func(repo *OmmsetRepoImpl) DeleteUserOmzet(id uuid.UUID) error{
	return repo.db.Where("umkm_id = ?", id).Delete(&domain.Omset{}).Error
}