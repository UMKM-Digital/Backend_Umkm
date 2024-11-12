package domain


type Provinsi struct {
	Id int       `gorm:"column:id;primaryKey;autoIncrement"`
	KodeWilayah  int       `gorm:"column:kode_wilayah"`
	NamaWilayah string  `gorm:"column:nama_wilayah"`
}

func (Provinsi) TableName() string {
	return "master.provinsi"
}

//Kabupaten
type Kabupaten struct {
	Id int       `gorm:"column:id;primaryKey;autoIncrement"`
	IdProvi  string       `gorm:"column:id_prov"`
	KodeKabupaten string  `gorm:"column:kode_kab"`
	NamaKabupaten string  `gorm:"column:nama"`
}

func (Kabupaten) TableName() string {
	return "master.kabupaten"
}

//kecamatan
type Kecamatan struct {
	Id int       `gorm:"column:id;primaryKey;autoIncrement"`
	IdProv  string       `gorm:"column:id_prov"`
	IdKab string `gorm:"column:id_kab"`
	KodeKec string `gorm:"column:kode_kec"`
	Nama string  `gorm:"column:nama"`
}

func (Kecamatan) TableName() string {
	return "master.kecamatan"
}


//keluarahan
type Keluarahan struct {
	Id int       `gorm:"column:id;primaryKey;autoIncrement"`
	IdProv  string       `gorm:"column:id_prov"`
	IdKab string `gorm:"column:id_kab"`
	IdKec string `gorm:"column:id_kec"`
	KodeKel string `gorm:"column:kode_kel"`
	Nama string  `gorm:"column:nama"`
}

func (Keluarahan) TableName() string {
	return "master.kelurahan"
}