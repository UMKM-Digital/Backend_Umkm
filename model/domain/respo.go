package domain

import (
	"time"

	"github.com/google/uuid"
)

type DokumenStatusResponse struct {
	Id            int    `json:"id"`
	Nama          string `json:"nama"`
	Status        int    `json:"status"` // 0 = Not Uploaded, 1 = Uploaded
	TanggalUpload time.Time `json:"tanggal_upload"`
}




type DokumenStatusResponseALL struct {
	UmkmId        uuid.UUID `json:"umkm_id"`
	Id            int       `json:"id"`
	Nama          string    `json:"nama"`
	Status        int 		`json:"status"`
	TanggalUpload time.Time `json:"tanggal_upload"`
}

type UmkmDocumentsResponse struct {
	UmkmID  uuid.UUID `json:"umkm_id"`
	Dokumen []DokumenStatusResponseALL `json:"dokumen"`
}
