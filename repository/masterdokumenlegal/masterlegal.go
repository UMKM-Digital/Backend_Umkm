package masterdokumenlegalrepo

import (
	"context"
	"umkm/model/domain"

	"github.com/google/uuid"
)

type MasterDokumenLegal interface {
	Created(dokumen domain.MasterDokumenLegal) (domain.MasterDokumenLegal, error)
	GetmasterlegalUmkm(filters string, limit int, page int) ([]domain.MasterDokumenLegal, int, int, int, *int, *int, error)
	DeleteMasterLegalId(id int) error
	GetMasterLegalId(id int)(domain.MasterDokumenLegal, error)
	UpdateMasterLegalId(id int, dokumen domain.MasterDokumenLegal)(domain.MasterDokumenLegal, error)
	GetDokumenUmkmStatus(umkmId uuid.UUID, filters string, limit int, page int) ([]domain.DokumenStatusResponse, int, int, int, *int, *int, error)
	GetAllMasterDokumenLegal(ctx context.Context) ([]domain.MasterDokumenLegal, error) 
	GetDokumenUmkmStatusAll(userId int, filters string, limit int, page int) ([]domain.UmkmDocumentsResponse, int, int, int, *int, *int, error) 
}