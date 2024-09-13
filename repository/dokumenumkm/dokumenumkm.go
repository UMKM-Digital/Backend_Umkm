package dokumenumkmrepo

import (
	"umkm/model/domain"

)

type DokumenUmkmrRepo interface {
	CreateRequest(dokumen domain.UmkmDokumen) (domain.UmkmDokumen, error)
}
