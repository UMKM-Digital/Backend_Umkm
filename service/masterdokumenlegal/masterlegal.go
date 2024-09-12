package masterdokumenlegalservice

import (
	"umkm/model/entity"
	"umkm/model/web"
)

type MasterDokumenLegal interface {
	CreatedRequest(masterlegal web.CreateMasterDokumenLegal) (map[string]interface{}, error)
	GetMasterLegalList(filters string, limit int, page int) (map[string]interface{}, error)
	DeleteMasterLegalId(id int) error
	GetMasterLegalid(id int)(entity.MasterlegalEntity, error)
	UpdateMasterLegal(request web.UpdateMasterDokumenLegal, id int) (map[string]interface{}, error)
}