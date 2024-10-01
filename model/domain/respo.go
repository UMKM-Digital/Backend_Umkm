package domain

import "time"

type DokumenStatusResponse struct {
	Id            int    `json:"id"`
	Nama          string `json:"nama"`
	Status        int    `json:"status"` // 0 = Not Uploaded, 1 = Uploaded
	TanggalUpload time.Time `json:"tanggal_upload"`
}