	package web

	import (
		"encoding/json"

		"github.com/google/uuid"
	)

	type CreateUmkmDokumenLegal struct {
		UmkmId         uuid.UUID `validate:"required" json:"umkm_id"`
		DokumenId 	   int `validate:"required" json:"dokumen_id"`
		DokumenUpload           json.RawMessage `json:"dok_upload"`
	}
