package transaksirepo

import (
	"errors"
	"umkm/model/domain"

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

