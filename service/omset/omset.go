package omsetservice

import (
	"umkm/model/entity"
	"umkm/model/web"

	"github.com/google/uuid"
)

type OmsetService interface {
	CreateOmsetService(omset web.Omset) (map[string]interface{}, error)
	ListOmsetService(umkm_id uuid.UUID, tahun string)([]entity.OmsetEntity, error)
	GetOmsetServiceId(id int) (entity.OmsetEntity, error)
	UpdateOmset(request web.UpdateOmset, pathId int) (map[string]interface{}, error) 
}