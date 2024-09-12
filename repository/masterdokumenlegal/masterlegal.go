package masterdokumenlegalrepo

import "umkm/model/domain"

type MasterDokumenLegal interface {
	Created(dokumen domain.MasterDokumenLegal) (domain.MasterDokumenLegal, error)
	GetmasterlegalUmkm(filters string, limit int, page int) ([]domain.MasterDokumenLegal, int, error)
	DeleteMasterLegalId(id int) error
	GetMasterLegalId(id int)(domain.MasterDokumenLegal, error)
	UpdateMasterLegalId(id int, dokumen domain.MasterDokumenLegal)(domain.MasterDokumenLegal, error)
}