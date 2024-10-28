package omsetrepo

import (
	"umkm/model/domain"

	"github.com/google/uuid"
)

type OmsetRepo interface {
	CreateRequest(produk domain.Omset) (domain.Omset, error)
	ListOmsetRequest(umkm_id uuid.UUID, tahun string)([]domain.Omset, error)
	GetOmsetId(id int)(domain.Omset, error)
	UpdateOmsetId(id int, omset domain.Omset) (domain.Omset, error) 
	OmsetTahunan(umkm_id uuid.UUID, tahun string) (float64, error) 
	OmsetBulanan(umkm_id uuid.UUID, tahun string) (map[string]float64, error) 
}