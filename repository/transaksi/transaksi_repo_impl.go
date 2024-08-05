package transaksirepo

import (
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


