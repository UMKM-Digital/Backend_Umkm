package transaksiservice

import (
	"errors"
	"fmt"
	"umkm/helper"
	"umkm/model/domain"
	"umkm/model/web"
	transaksirepo "umkm/repository/transaksi"

	"github.com/shopspring/decimal"
)

type TranssaksiServiceImpl struct {
	transaksirepository transaksirepo.TransaksiRepo
}

func NewTransaksiservice(transaksirepository transaksirepo.TransaksiRepo) *TranssaksiServiceImpl {
	return &TranssaksiServiceImpl{
		transaksirepository: transaksirepository,
	}
}

// register
// service/transaksiservice/auth_service_impl.go
func (service *TranssaksiServiceImpl) CreateTransaksi(transaksi web.CreateTransaksi) (map[string]interface{}, error) {
	date, err := helper.ParseDate(transaksi.Tanggal)
	IDKategoriProduk, err := helper.RawMessageToJSONB(transaksi.IDKategoriProduk)
	if err != nil {
		return nil, errors.New("invalid type for idKategoriProduk")
	}

	invoiceNumber, err := helper.GenerateInvoiceNumber()
	if err != nil {
		return nil, errors.New("failed to generate invoice number")
	}

	qrcode, err := helper.GenerateValidationTicket()
	if err != nil {
		return nil, errors.New("failed to generate QR code")
	}

	if qrcode == "" {
		return nil, errors.New("QR code generation returned empty value")
	}

	// Logging QR code untuk memastikan
	fmt.Printf("Generated QR Code: %s\n", qrcode)
	fmt.Printf("Received Total Jml: %f\n", transaksi.TotalJml) // Log total_jml

	totalJmlDecimal := decimal.NewFromFloat(transaksi.TotalJml)

	newTransaksi := domain.Transaksi{
		UmkmId:           transaksi.UmkmId,
		NoInvoice:        invoiceNumber,
		Tanggal:          date,
		Nameclient:       transaksi.NamaClient,
		IdKategoriProduk: IDKategoriProduk,
		TotalJml:         totalJmlDecimal,
		Keterangan:       transaksi.Keteranagan,
		Status:           transaksi.Status,
		TiketValidasi:    qrcode, // Pastikan ini terisi
	}

	saveTransaksi, errSaveTransaksi := service.transaksirepository.CreateRequetsTransaksi(newTransaksi)
	if errSaveTransaksi != nil {
		return nil, errSaveTransaksi
	}

	return map[string]interface{}{
		"Id Umkm":            saveTransaksi.UmkmId,
		"No Invoice":         saveTransaksi.NoInvoice,
		"Tanggal":            saveTransaksi.Tanggal,
		"Name Client":        saveTransaksi.Nameclient,
		"Id Kategori Produk": saveTransaksi.IdKategoriProduk,
		"Total":              saveTransaksi.TotalJml,
		"Keterangan":         saveTransaksi.Keterangan,
	}, nil
}
