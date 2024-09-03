package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)



type Transaksi struct{
	IdTransaksi     int       `gorm:"column:id;primaryKey;autoIncrement"`
	UmkmId           uuid.UUID `gorm:"column:umkm_id;type:uuid"`
    NoInvoice       string    `gorm:"column:no_invoice"`
    Tanggal       time.Time    `gorm:"column:tanggal"`
    Nameclient       string    `gorm:"column:name_client"`
    IdKategoriProduk       JSONB    `gorm:"column:id_kategori_produk"`
    TotalJml        decimal.Decimal `gorm:"column:total_jml;type:numeric(15,2)"`
	Keterangan     string   			`gorm:"column:keterangan"`
	Status     	   	int    `gorm:"column:status"`
	NoHp     	   	string    `gorm:"column:no_hp"`
	AlasanPerubahan     	   	string    `gorm:"column:status"`
	TiketValidasi     	   	string    `gorm:"column:tiket_validasi"`
    Created_at time.Time `gorm:"column:created_at"`
    Updated_at time.Time `gorm:"column:updated_at"`
    Umkm    UMKM `gorm:"foreignKey:UmkmId"`
}

type Kategori_Produk struct {
    ID   []string `json:"id"`
    Nama []string `json:"nama"`
}


func (Transaksi) TableName() string {
    return "transaksi"
}