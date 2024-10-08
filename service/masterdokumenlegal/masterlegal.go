package masterdokumenlegalservice

import (
	"umkm/model/domain"
	"umkm/model/entity"
	"umkm/model/web"

	"github.com/google/uuid"
	// "github.com/google/uuid"
)

type MasterDokumenLegal interface {
	CreatedRequest(masterlegal web.CreateMasterDokumenLegal) (map[string]interface{}, error)
	GetMasterLegalList(filters string, limit int, page int) ([]entity.MasterlegalEntity, int, int, int, *int, *int, error)
	DeleteMasterLegalId(id int) error
	GetMasterLegalid(id int)(entity.MasterlegalEntity, error)
	UpdateMasterLegal(request web.UpdateMasterDokumenLegal, id int) (map[string]interface{}, error)
	GetDokumenUmkmStatus(umkmId uuid.UUID, filters string, limit int, page int) ([]domain.DokumenStatusResponse, int, int, int, *int, *int, error) 
	GetDokumenUmkmStatusAll(userId int, filters string, limit int, page int) ([]domain.DokumenStatusResponseALL, int, int, int, *int, *int, error) 
}