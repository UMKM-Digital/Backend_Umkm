package hakaksesservice

import (
	"umkm/model/web"
)

type HakAkses interface {
	UpdateBulkHakAkses(request web.HakAksesUpdate) ([]map[string]interface{}, error)
}