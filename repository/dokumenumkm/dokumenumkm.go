package dokumenumkmrepo

import (
	"umkm/model/domain"

	"github.com/google/uuid"
)

type DokumenUmkmrRepo interface {
	CreateRequest(dokumen domain.UmkmDokumen) (domain.UmkmDokumen, error)
	GetId(id int, umkmid uuid.UUID) (domain.UmkmDokumen, error)
	UpdateDokumen(id int, umkmid uuid.UUID, dokumenumkm domain.UmkmDokumen)(domain.UmkmDokumen, error)
	DeleteDokumenUmkmId(id uuid.UUID) error
}
