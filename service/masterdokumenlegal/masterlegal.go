package masterdokumenlegalservice

import "umkm/model/web"

type MasterDokumenLegal interface {
	CreatedRequest(masterlegal web.CreateMasterDokumenLegal) (map[string]interface{}, error)
}