package web

type CreateMasterDokumenLegal struct{
	Name string `validate:"required" json:"nama"`
	Is_Wajib int `validate:"required" json:"is_wajib"`
}
type UpdateMasterDokumenLegal struct{
	Name string `validate:"required" json:"nama"`
	Is_Wajib *int `validate:"required" json:"is_wajib"`
}