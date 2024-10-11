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
	IdKabupaten  string       `gorm:"column:kode_kabupaten"`
	KodeWilayah  string       `gorm:"column:kode_wilayah"`
	NamaWilayah string  `gorm:"column:nama_wilayah"`
}

func (Kecamatan) TableName() string {
	return "master.kode_kec"
}


//keluarahan
type Keluarahan struct {
	Id int       `gorm:"column:id;primaryKey;autoIncrement"`
	KodeKecamatan  string       `gorm:"column:kode_kec"`
	KodeKewilayah  string       `gorm:"column:kode_wilayah"`
	NamaWilayah string  `gorm:"column:nama_wilayah"`
}

func (Keluarahan) TableName() string {
	return "master.kode_kel"
}