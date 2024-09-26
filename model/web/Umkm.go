package web

import "encoding/json"

type UmkmRequest struct {
	Name                  string          `json:"name"`
	NoNpwp                string          `json:"no_npwp"`
	Kategori_Umkm_Id      json.RawMessage `json:"kategori_umkm_id"`
	Nama_Penanggung_Jawab string          `json:"nama_penanggung_jawab"`
	Informasi_JamBuka     json.RawMessage `json:"informasi_jambuka"`
	No_Kontak             string          `json:"no_kontak"`
	Lokasi                string          `json:"lokasi"`
	Deskripsi                string          `json:"deskripsi"`
	Maps                  json.RawMessage `json:"maps"`
	Gambar                json.RawMessage `json:"gambar"`  // Field for images
	UserId                int             `json:"user_id"` // Add UserId field
}

type Updateumkm struct{
	Name                  string          `json:"name"`
	NoNpwp                string          `json:"no_npwp"`
	Kategori_Umkm_Id      json.RawMessage `json:"kategori_umkm_id"`
	Nama_Penanggung_Jawab string          `json:"nama_penanggung_jawab"`
	Informasi_JamBuka     json.RawMessage `json:"informasi_jambuka"`
	No_Kontak             string          `json:"no_kontak"`
	Lokasi                string          `json:"lokasi"`
	Deskripsi                string          `json:"deskripsi"`
	Maps                  json.RawMessage `json:"maps"`
	Gambar                json.RawMessage `json:"gambar"`
}