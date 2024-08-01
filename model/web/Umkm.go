package web

import "encoding/json"

type UmkmRequest struct {
	Name                   string          `validate:"required" json:"name"`
	NoNpwp                 string          `validate:"required" json:"no_npwp"`
	Gambar                 json.RawMessage `validate:"required" json:"gambar"`
	Kategori_Umkm_Id       json.RawMessage `validate:"required" json:"kategori_umkm_id"`
	Nama_Penanggung_Jawab  string          `validate:"required" json:"nama_penanggung_jawab"`
	Informasi_JamBuka      json.RawMessage `validate:"required" json:"informasi_jambuka"`
	No_Kontak              string          `validate:"required" json:"no_kontak"`
	Lokasi                 string          `validate:"required" json:"lokasi"`
	Maps                   json.RawMessage `validate:"required" json:"maps"`
}
