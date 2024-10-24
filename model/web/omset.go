package web

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Omset struct {
	Bulan      string          `json:"bulan"`
	JumlahOmset decimal.Decimal `json:"jumlah_omset"`
	UmkmId     uuid.UUID       `json:"umkm_id"`
}

