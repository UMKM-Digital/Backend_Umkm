package masterdokumenlegalrepo

import "umkm/model/domain"

type MasterDokumenLegal interface {
	Created(dokumen domain.MasterDokumenLegal) (domain.MasterDokumenLegal, error)
}