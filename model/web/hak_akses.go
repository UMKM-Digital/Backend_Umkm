package web

import "errors"

// Definisikan enum untuk status
const (
    StatusPending    = "menunggu"
    StatusAccepted   = "disetujui"
    StatusRejected     = "ditolak"
)


type HakAksesUpdate struct {
	 UmkmIds []string `json:"umkm_ids" validate:"required"`
    Status   string    `validate:"required" json:"status"`
    Pesan   string    `validate:"required" json:"pesan"`
    Pembina   string    `validate:"required" json:"pembina"`
}

func IsValidStatus(status string) bool {
    switch status {
    case StatusPending, StatusAccepted, StatusRejected:
        return true
    }
    return false
}

// Fungsi yang bisa digunakan untuk validasi struct HakAksesUpdate
func (h *HakAksesUpdate) ValidateStatus() error {
    if !IsValidStatus(h.Status) {
        return errors.New("invalid status value")
    }
    return nil
}
