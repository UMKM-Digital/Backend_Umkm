package domain

import "time"

type Users struct {
    IdUser     int       `gorm:"column:id;primaryKey;autoIncrement"`
    Username   string    `gorm:"column:username"`
    Email      string    `gorm:"column:email"`
    Password   string    `gorm:"column:password"`
	Role       string    `gorm:"column:role"`
	No_Phone   string    `gorm:"column:no_phone"`
	Picture    string    `gorm:"column:picture"`
    GoogleId string         `gorm:"column:google_id"`
    NoKk string         `gorm:"column:no_kk"`
    Nik string         `gorm:"column:nik"`
    Nib string         `gorm:"column:nib"`
    TanggalLahir time.Time         `gorm:"column:tanggal_lahir"`
    JenisKelamin string         `gorm:"column:jenis_kelamin"`
    StatusMenikah string         `gorm:"column:status_menikah"`
    Rt string         `gorm:"column:rt"`
    Rw string         `gorm:"column:rw"`
    KodePos string         `gorm:"column:kode_pos"`
    Fullname string         `gorm:"column:full_name"`
    Alamat string         `gorm:"column:alamat"`
    Provinsi string         `gorm:"column:provinsi"`
    Kabupaten string         `gorm:"column:kabupaten"`
    Kelurahan string         `gorm:"column:kelurahan"`
    Kecamatan string         `gorm:"column:kecamatan"`
    Ktp JSONB  `gorm:"column:ktp"`
    KartuKeluarga JSONB `gorm:"column:kartu_keluarga"`
   CreatedAt            time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt            time.Time `gorm:"column:updated_at;autoUpdateTime"`
    HakAkses   []HakAkses `gorm:"foreignKey:user_id"`
    Berita   []Berita `gorm:"foreignKey:author"`
}

func (Users) TableName() string {
    return "users"
}
