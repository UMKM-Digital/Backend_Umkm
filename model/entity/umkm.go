package entity

import (
	"umkm/model/domain"
	"sort"
	"github.com/google/uuid"
)

// Struct untuk UmkmEntityList (versi list singkat tanpa detail total produk)
type UmkmEntityList struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// Fungsi untuk mengonversi domain.UMKM ke UmkmEntityList (list sederhana)
func ToUmkmEntityList(umkm domain.UMKM) UmkmEntityList {
	return UmkmEntityList{
		Id:   umkm.IdUmkm,
		Name: umkm.Name,
	}
}

// Fungsi untuk mengonversi daftar domain.UMKM ke daftar UmkmEntityList (versi sederhana)
func ToUmkmListEntities(umkmList []domain.UMKM) []UmkmEntityList {
	var umkmListEntities []UmkmEntityList
	for _, umkm := range umkmList {
		umkmListEntities = append(umkmListEntities, ToUmkmEntityList(umkm))
	}
	return umkmListEntities
}


type  UmkmEntity struct{
	Id uuid.UUID `json:"id"`
	Name string `json:"name"`
	NoNpwp string `json:"no_npwp"`
	Gambar domain.JSONB `json:"gambar"`
	KategoriUmkmId domain.JSONB `json:"kategori_umkm_id"`
	NamaPenanggungJawab string `json:"nama_penanggung_jawab"`
	InformasiJamBuka domain.JSONB `json:"informasi_jambuka"`
	Deskripsi string `json:"deskripsi"`
	NoKontak string `json:"no_kontak"`
	Lokasi string `json:"lokasi"`
	Maps domain.JSONB `json:"maps"`
}

func ToUmkmEntity(umkm domain.UMKM) UmkmEntity {
	return UmkmEntity{
		Id:   umkm.IdUmkm,
		Name: umkm.Name,
		NoNpwp: umkm.NoNpwp,
		Gambar: umkm.Images,
		KategoriUmkmId: umkm.KategoriUmkmId,
		NamaPenanggungJawab: umkm.NamaPenanggungJawab,
		InformasiJamBuka: umkm.InformasiJambuka,
		NoKontak: umkm.NoKontak,
		Lokasi: umkm.Lokasi,
		Maps: umkm.Maps,
	}
}

// Fungsi untuk mengonversi daftar domain.UMKM ke daftar UmkmEntityList (versi sederhana)
func ToUmkmEntities(umkmList []domain.UMKM) []UmkmEntity {
	var umkmListEntities []UmkmEntity
	for _, umkm := range umkmList {
		umkmListEntities = append(umkmListEntities, ToUmkmEntity(umkm))
	}
	return umkmListEntities
}


// ProdukEntityWebList menyimpan informasi produk yang akan ditampilkan di web
type ProdukEntityWebList struct {
	Id     uuid.UUID   `json:"id"`   
	Gambar domain.JSONB `json:"gambar"` // Menyimpan gambar produk
	Nama string `json:"nama_produk"`
}

// UmkmEntityWebList menyimpan informasi UMKM yang akan ditampilkan di web
type UmkmEntityWebList struct {
	Id           uuid.UUID             `json:"id"` 
	Name         string                `json:"name"`
	Gambar       domain.JSONB          `json:"gambar"` // Menggunakan domain.JSONB untuk menyimpan gambar UMKM
	Lokasi       string                `json:"lokasi"`
	KategoriUmkm domain.JSONB 			`json:"kategori"`
	InformasiJmbuka domain.JSONB        `json:"informasi_jambuka"`
	Deskripsi string 					`json:"deskripsi"`
	GambarProduk []ProdukEntityWebList  `json:"gambar_produk"` // Menyimpan daftar produk
}

// ToProdukEntityWebList mengonversi domain.Produk ke ProdukEntityWebList
func ToProdukEntityWebList(produk domain.Produk) ProdukEntityWebList {
	return ProdukEntityWebList{
		Id:     produk.IdUmkm,     // Pastikan ini adalah ID produk yang benar
		Gambar: produk.Gamabr, // Perbaiki dari Gamabr ke Gambar
		Nama: produk.Nama,
	}
}

// ToProdukEntitiesWebList mengonversi daftar produk menjadi daftar ProdukEntityWebList
func ToProdukEntitiesWebList(produkList []domain.Produk) []ProdukEntityWebList {
	// Urutkan produk berdasarkan created_at dari terbaru ke terlama
	sort.Slice(produkList, func(i, j int) bool {
		return produkList[i].CreatedAt.After(produkList[j].CreatedAt)
	})

	// Ambil maksimal 3 produk
	var produkListEntities []ProdukEntityWebList
	limit := 3
	for i, produk := range produkList {
		if i >= limit {
			break
		}
		produkListEntities = append(produkListEntities, ToProdukEntityWebList(produk))
	}
	return produkListEntities
}

// ToUmkmEntityWebList mengonversi UMKM menjadi UmkmEntityWebList
func ToUmkmEntityWebList(umkm domain.UMKM) UmkmEntityWebList {
	

	// Panggil fungsi untuk mendapatkan 3 produk terbaru
	produkList := ToProdukEntitiesWebList(umkm.Produk)

	return UmkmEntityWebList{
		Id:           umkm.IdUmkm,
		Name:         umkm.Name,
		Gambar:       umkm.Images, // Pastikan ini adalah field gambar UMKM yang benar
		Lokasi:       umkm.Lokasi,
		GambarProduk: produkList,
		InformasiJmbuka: umkm.InformasiJambuka,
		Deskripsi:  umkm.Deskripsi,
		KategoriUmkm: umkm.KategoriUmkmId,
	}
}

// ToUmkmEntitiesWebList mengonversi daftar UMKM menjadi daftar UmkmEntityWebList
func ToUmkmEntitiesWebList(umkmList []domain.UMKM) []UmkmEntityWebList {
	var umkmListEntities []UmkmEntityWebList
	for _, umkm := range umkmList {
		umkmListEntities = append(umkmListEntities, ToUmkmEntityWebList(umkm)) // Perbaiki penamaan fungsi
	}
	return umkmListEntities
}


//detailumkm

type UmkmDetailEntity struct {
	Id                uuid.UUID                     `json:"id"`
	Gambar            domain.JSONB                  `json:"gambar_umkm"`
	Name              string                         `json:"name_umkm"`
	KatageoriUmkm     domain.JSONB                  `json:"kategori_umkmk"`
	NoKontak          string                         `json:"no_kontak"`
	Deskripsi         string                         `json:"deskripsi_umkm"`
	InformasiJambuka  domain.JSONB                  `json:"jambuka_umkm"`
	Lokasi            string                         `json:"lokasi"`
	Maps              domain.JSONB                  `json:"maps"`
	Produk            []ProdukEntityDetailList       `json:"all_products"` // Produk yang terpaginasikan
}


type ProdukEntityDetailList struct {
	Id     uuid.UUID   `json:"id"`   
	Gambar domain.JSONB `json:"gambar_porduk"` // Menyimpan gambar produk
	Nama string `json:"nama_produk"`
	Harga int `json:"harga"`
}

func ToProdukEntityDetailList(produk domain.Produk) ProdukEntityDetailList {
	return ProdukEntityDetailList{
		Id:     produk.IdUmkm,    
		Gambar: produk.Gamabr, 
		Nama: produk.Nama,
		Harga: produk.Harga,
	}
}

func ToProdukEntitiesDetailList(produkList []domain.Produk) []ProdukEntityDetailList {
	// Urutkan produk berdasarkan created_at dari terlama ke terbaru
	sort.Slice(produkList, func(i, j int) bool {
		return produkList[i].CreatedAt.Before(produkList[j].CreatedAt)
	})

	// Konversi seluruh produk ke ProdukEntityDetailList tanpa limit
	var produkListEntities []ProdukEntityDetailList
	for _, produk := range produkList {
		produkListEntities = append(produkListEntities, ToProdukEntityDetailList(produk))
	}
	return produkListEntities
}
//
func ToProdukEntitiesDetailListNew(produkList []domain.Produk) []ProdukEntityDetailList {
	// Urutkan produk berdasarkan created_at dari terbaru ke terlama
	sort.Slice(produkList, func(i, j int) bool {
		return produkList[i].CreatedAt.After(produkList[j].CreatedAt)
	})

	// Ambil maksimal 4 produk
	limit := 4
	if len(produkList) < limit {
		limit = len(produkList)
	}

	// Konversi hanya 4 produk teratas ke ProdukEntityDetailList
	var produkListEntities []ProdukEntityDetailList
	for i := 0; i < limit; i++ {
		produkListEntities = append(produkListEntities, ToProdukEntityDetailList(produkList[i]))
	}
	return produkListEntities
}


// ToUmkmEntityWebList mengonversi UMKM menjadi UmkmEntityWebList
func ToUmkmEntityDetailList(umkm domain.UMKM) UmkmDetailEntity {
	

	// Panggil fungsi untuk mendapatkan 3 produk terbaru
	produkList := ToProdukEntitiesDetailList(umkm.Produk)
	

	return UmkmDetailEntity{
		Id:           umkm.IdUmkm,
		Name:         umkm.Name,
		Gambar:       umkm.Images, // Pastikan ini adalah field gambar UMKM yang benar
		Lokasi:       umkm.Lokasi,
		KatageoriUmkm: umkm.KategoriUmkmId,
		Deskripsi: umkm.Deskripsi,
		InformasiJambuka: umkm.InformasiJambuka,
		NoKontak: umkm.NoKontak,
		Maps: umkm.Maps,
		Produk: produkList,
	}
}

// ToUmkmEntitiesWebList mengonversi daftar UMKM menjadi daftar UmkmEntityWebList
func ToUmkmEntitiesDetailList(umkmList []domain.UMKM) []UmkmDetailEntity {
	var umkmListEntities []UmkmDetailEntity
	for _, umkm := range umkmList {
		umkmListEntities = append(umkmListEntities, ToUmkmEntityDetailList(umkm)) // Perbaiki penamaan fungsi
	}
	return umkmListEntities
}


//list buat rashit
type UmkmActive struct{
	Id uuid.UUID `json:"id"`
	Nama string `json:"nama"`
	Gambar domain.JSONB `json:"gambar"`
	Deskripsi string 	`json:"deskripsi"`
	Active int 	`json:"active"`
}

func ToUmkmEntityActive(umkm domain.UMKM) UmkmActive {
	return UmkmActive{
		Id:           umkm.IdUmkm,
		Nama:         umkm.Name,
		Gambar:       umkm.Images, // Pastikan ini adalah field gambar UMKM yang bena
		Deskripsi: umkm.Deskripsi,
		Active: umkm.Active,
	}
}

func ToUmkmListEntitiesActive(umkmList []domain.UMKM) []UmkmActive {
	var umkmListEntitiesActive []UmkmActive
	for _, umkm := range umkmList {
		umkmListEntitiesActive = append(umkmListEntitiesActive, ToUmkmEntityActive(umkm))
	}
	return umkmListEntitiesActive
}