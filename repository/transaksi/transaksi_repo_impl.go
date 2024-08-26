package transaksirepo

import (
	"errors"
	"umkm/model/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransaksirepositoryImpl struct {
	db *gorm.DB
}

func NewTransaksiRepositoryImpl(db *gorm.DB) *TransaksirepositoryImpl {
	return &TransaksirepositoryImpl{db: db}
}

// regoster
func (repo *TransaksirepositoryImpl) CreateRequetsTransaksi(transaksi domain.Transaksi)(domain.Transaksi, error) {
	if err := repo.db.Create(&transaksi).Error; err != nil {
		return domain.Transaksi{}, err
	}
	return transaksi, nil
}

//get kategori by id
func(repo *TransaksirepositoryImpl) GetRequestTransaksi(idTransaksi int)(domain.Transaksi, error){
	var TrasnsaksiUmkm domain.Transaksi
	err := repo.db.Find(&TrasnsaksiUmkm, "id = ?", idTransaksi).Error
	if err != nil{
		return domain.Transaksi{}, errors.New("Transaksi tidak ditemukan")
	}
	return TrasnsaksiUmkm, nil
}

func (repo *TransaksirepositoryImpl) GetFilterTransaksi(umkmID uuid.UUID) ([]domain.Transaksi, error) {
	var transaksi []domain.Transaksi
	err := repo.db.Where("umkm_id = ?", umkmID).Find(&transaksi).Error
	if err != nil {
		return nil, err
	}
	return transaksi, nil
}

func (repo *TransaksirepositoryImpl) GetFilterTransaksiWebTahun(umkmId uuid.UUID, year int) ([]domain.Transaksi, error) {
	var transaksi []domain.Transaksi
	err := repo.db.Where("umkm_id = ? AND EXTRACT(YEAR FROM tanggal) = ?", umkmId, year).Find(&transaksi).Error
	if err != nil {
		return nil, err
	}
	return transaksi, nil
}

func (repo *TransaksirepositoryImpl) GetFilterTransaksiWebBulan(umkmId uuid.UUID, mounth int) ([]domain.Transaksi, error) {
	var transaksi []domain.Transaksi
	err := repo.db.Where("umkm_id = ? AND EXTRACT(MONTH FROM tanggal) = ?", umkmId, mounth).Find(&transaksi).Error
	if err != nil {
		return nil, err
	}
	return transaksi, nil
}



