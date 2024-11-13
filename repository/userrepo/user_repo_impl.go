package userrepo

import (
	"errors"
	"math"

	"umkm/model/domain"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthrepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepositoryImpl(db *gorm.DB) *AuthrepositoryImpl{
	return &AuthrepositoryImpl{db:db}
}

//register
func (repo *AuthrepositoryImpl) RegisterRequest(user domain.Users)(domain.Users, error){
	err := repo.db.Create(&user).Error
	if err != nil {
		return domain.Users{}, err
	}

	return user, nil
}

//login
func (repo *AuthrepositoryImpl) FindUserByUsername(username string) (*domain.Users, error){
	user := new(domain.Users)

	if err	:= repo.db.Where("username = ?", username).Take(&user).Error; err != nil{
		return user, err
	}

	return user, nil
}

//otp
func (repo *AuthrepositoryImpl) FindUserByPhone(phone string) (*domain.Users, error) {
	user := new(domain.Users)

	if err := repo.db.Where("no_phone = ?", phone).Take(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

//profile
func (repo *AuthrepositoryImpl) GetByID(idUser int) (domain.Users, error) {
	var user domain.Users
	if err := repo.db.Where("id = ?", idUser).Take(&user).Error; err != nil {
		return domain.Users{}, errors.New("user not found")
	}
	return user, nil
}


//send otp register
func (repo *AuthrepositoryImpl) FindUserByPhoneRegister(phone string) (*domain.Users, error) {
    user := new(domain.Users)
    result := repo.db.Where("no_phone = ?", phone).Take(user)
    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return nil, nil // Nomor telepon tidak ditemukan
        }
        return nil, result.Error // Terjadi kesalahan lain
    }
    return user, nil
}


// repository/userrepo/auth_repository.go
func (repo *AuthrepositoryImpl) UpdateId(idUser int, user domain.Users) (domain.Users, error) {
    if err := repo.db.Model(&domain.Users{}).Where("id = ?", idUser).Updates(user).Error; err != nil {
        return domain.Users{}, errors.New("failed to update profile")
    }
    return user, nil
}

//verivy otp


//cek in paswword
func(repo *AuthrepositoryImpl) CekInPassword(userId int) (*domain.Users, error) {
    user := new(domain.Users)

    if err := repo.db.Where("id = ?", userId).Take(&user).Error; err != nil {
        return nil, err
    }

    return user, nil
}

//ubah password
func (repo *AuthrepositoryImpl) UpdatePassword(userId int, newPassword string) error {
    // Hash password baru
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    // Update password di database
    if err := repo.db.Model(&domain.Users{}).Where("id = ?", userId).Update("password", hashedPassword).Error; err != nil {
        return err
    }

    return nil
}


//login goole
func (repo *AuthrepositoryImpl) FindOrCreateUserByGoogleID(googleID string, email string, username string, picture string) (*domain.Users, error) {
    var user domain.Users	
    if err := repo.db.Where("google_id = ?", googleID).First(&user).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            // Jika user tidak ditemukan, buat user baru
            newUser := domain.Users{
                GoogleId: googleID,
                Email:    email,
				Fullname: username,
                Picture: picture,
                Role: "umkm",
            }
            if err := repo.db.Create(&newUser).Error; err != nil {
                return nil, err
            }
            return &newUser, nil
        }
        return nil, err
    }
    return &user, nil
}

//resert Password
func(repo *AuthrepositoryImpl) ChangePassword(email string) (*domain.Users, error){
    user := new(domain.Users)

	if err	:= repo.db.Where("email = ?", email).Take(&user).Error; err != nil{
		return user, err
	}

	return user, nil
}

func(repo *AuthrepositoryImpl) ListUser()([]domain.Users, error){
    var user []domain.Users
	err := repo.db.Order("id ASC").Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *AuthrepositoryImpl) CountUserByGenderWithPercentage() (map[string]map[string]float64, error) {
    var countLakiLaki, countPerempuan, totalUsers int64

    // Menghitung total pengguna unik dengan role 'umkm' yang memiliki UMKM dengan status 1
    err := repo.db.Table("users").
        Joins("JOIN hak_akses ON users.id = hak_akses.user_id").
        Where("users.role = ? AND hak_akses.status = ?", "umkm", "disetujui").
        Group("users.id").
        Count(&totalUsers).Error
    if err != nil {
        return nil, err
    }

    // Menghitung jumlah laki-laki unik dengan role 'umkm' yang memiliki UMKM dengan status 1
    err = repo.db.Table("users").
        Joins("JOIN hak_akses ON users.id = hak_akses.user_id").
        Where("users.role = ? AND users.jenis_kelamin = ? AND hak_akses.status = ?", "umkm", "Laki-laki", "disetujui").
        Group("users.id").
        Count(&countLakiLaki).Error
    if err != nil {
        return nil, err
    }

    // Menghitung jumlah perempuan unik dengan role 'umkm' yang memiliki UMKM dengan status 1
    err = repo.db.Table("users").
        Joins("JOIN hak_akses ON users.id = hak_akses.user_id").
        Where("users.role = ? AND users.jenis_kelamin = ? AND hak_akses.status = ?", "umkm", "Perempuan", "disetujui").
        Group("users.id").
        Count(&countPerempuan).Error
    if err != nil {
        return nil, err
    }

    // Menghitung persentase
    var percentageLakiLaki, percentagePerempuan float64
    if totalUsers > 0 {
        percentageLakiLaki = math.Round((float64(countLakiLaki) / float64(totalUsers)) * 100 * 10) / 10
        percentagePerempuan = math.Round((float64(countPerempuan) / float64(totalUsers)) * 100 * 10) / 10
    } else {
        // Jika tidak ada pengguna, set persentase ke 0
        percentageLakiLaki = 0
        percentagePerempuan = 0
    }
    // Mengembalikan hasil dalam bentuk map dengan persentase
    result := map[string]map[string]float64{
        "l": {
            "total":      float64(countLakiLaki),
            "persentase": percentageLakiLaki,
        },
        "p": {
            "total":      float64(countPerempuan),
            "persentase": percentagePerempuan,
        },
    }

    return result, nil
}

func (repo *AuthrepositoryImpl) CountUserByStudy() (map[string]int64, error) {
    var SD, SMP, SMA, SLTA, D3, D4, S2, S3 int64

    // Menghitung jumlah pengguna dengan pendidikan terakhir SD
    err := repo.db.Table("users").
        Joins("JOIN hak_akses ON users.id = hak_akses.user_id").
        Where("users.role = ? AND users.pendidikan_terakhir = ? AND hak_akses.status = ?", "umkm", "SD", "disetujui").
        Group("users.id"). // Mengelompokkan berdasarkan user.id
        Count(&SD).Error
    if err != nil {
        return nil, err
    }

    // Menghitung jumlah pengguna dengan pendidikan terakhir SMP
    err = repo.db.Table("users").
        Joins("JOIN hak_akses ON users.id = hak_akses.user_id").
        Where("users.role = ? AND users.pendidikan_terakhir = ? AND hak_akses.status = ?", "umkm", "SMP", "disetujui").
        Group("users.id"). // Mengelompokkan berdasarkan user.id
        Count(&SMP).Error
    if err != nil {
        return nil, err
    }

    // Menghitung jumlah pengguna dengan pendidikan terakhir SMA
    err = repo.db.Table("users").
        Joins("JOIN hak_akses ON users.id = hak_akses.user_id").
        Where("users.role = ? AND users.pendidikan_terakhir = ? AND hak_akses.status = ?", "umkm", "SMA/K", "disetujui").
        Group("users.id"). // Mengelompokkan berdasarkan user.id
        Count(&SMA).Error
    if err != nil {
        return nil, err
    }

    // Menghitung jumlah pengguna dengan pendidikan terakhir SLTA
    err = repo.db.Table("users").
        Joins("JOIN hak_akses ON users.id = hak_akses.user_id").
        Where("users.role = ? AND users.pendidikan_terakhir = ? AND hak_akses.status = ?", "umkm", "SLTA", "disetujui").
        Group("users.id"). // Mengelompokkan berdasarkan user.id
        Count(&SLTA).Error
    if err != nil {
        return nil, err
    }

    // Menghitung jumlah pengguna dengan pendidikan terakhir D3
    err = repo.db.Table("users").
        Joins("JOIN hak_akses ON users.id = hak_akses.user_id").
        Where("users.role = ? AND users.pendidikan_terakhir = ? AND hak_akses.status = ?", "umkm", "D3", "disetujui").
        Group("users.id"). // Mengelompokkan berdasarkan user.id
        Count(&D3).Error
    if err != nil {
        return nil, err
    }

    // Menghitung jumlah pengguna dengan pendidikan terakhir D4/S1
    err = repo.db.Table("users").
        Joins("JOIN hak_akses ON users.id = hak_akses.user_id").
        Where("users.role = ? AND users.pendidikan_terakhir = ? AND hak_akses.status = ?", "umkm", "D4 / S-I", "disetujui").
        Group("users.id"). // Mengelompokkan berdasarkan user.id
        Count(&D4).Error
    if err != nil {
        return nil, err
    }

    // Menghitung jumlah pengguna dengan pendidikan terakhir S2
    err = repo.db.Table("users").
        Joins("JOIN hak_akses ON users.id = hak_akses.user_id").
        Where("users.role = ? AND users.pendidikan_terakhir = ? AND hak_akses.status = ?", "umkm", "S-II", "disetujui").
        Group("users.id"). // Mengelompokkan berdasarkan user.id
        Count(&S2).Error
    if err != nil {
        return nil, err
    }

    // Menghitung jumlah pengguna dengan pendidikan terakhir S3
    err = repo.db.Table("users").
        Joins("JOIN hak_akses ON users.id = hak_akses.user_id").
        Where("users.role = ? AND users.pendidikan_terakhir = ? AND hak_akses.status = ?", "umkm", "S-III", "disetujui").
        Group("users.id"). // Mengelompokkan berdasarkan user.id
        Count(&S3).Error
    if err != nil {
        return nil, err
    }

    // Mengembalikan hasil dalam bentuk map
    result := map[string]int64{
        "sd":   SD,
        "smp":  SMP,
        "sma":  SMA,
        "slta": SLTA,
        "d3":   D3,
        "d4":   D4,
        "s2":   S2,
        "s3":   S3,
    }

    return result, nil
}

func (repo *AuthrepositoryImpl) CountUserByAge() (map[string]int64, error) {
    var age11_20, age21_30, age31_40, age41_50, age51_60, age61_70, age71_80, age81_90 int64

    // Menghitung jumlah pengguna dengan role 'umkm' dan rentang umur 11-20
    err := repo.db.Table("users").
        Joins("JOIN hak_akses ON users.id = hak_akses.user_id").
        Where("role = ? AND date_part('year', age(tanggal_lahir)) BETWEEN ? AND ? AND hak_akses.status = ?", "umkm", 11, 20, "disetujui").
        Distinct("users.id").Count(&age11_20).Error
    if err != nil {
        return nil, err
    }

    // Menghitung jumlah pengguna dengan rentang umur 21-30
    err = repo.db.Table("users").
        Joins("JOIN hak_akses ON users.id = hak_akses.user_id").
        Where("role = ? AND date_part('year', age(tanggal_lahir)) BETWEEN ? AND ? AND hak_akses.status = ?", "umkm", 21, 30, "disetujui").
        Distinct("users.id").Count(&age21_30).Error
    if err != nil {
        return nil, err
    }

    // Menghitung jumlah pengguna dengan rentang umur 31-40
    err = repo.db.Table("users").
        Joins("JOIN hak_akses ON users.id = hak_akses.user_id").
        Where("role = ? AND date_part('year', age(tanggal_lahir)) BETWEEN ? AND ? AND hak_akses.status = ?", "umkm", 31, 40, "disetujui").
        Distinct("users.id").Count(&age31_40).Error
    if err != nil {
        return nil, err
    }

    // Menghitung jumlah pengguna dengan rentang umur 41-50
    err = repo.db.Table("users").
        Joins("JOIN hak_akses ON users.id = hak_akses.user_id").
        Where("role = ? AND date_part('year', age(tanggal_lahir)) BETWEEN ? AND ? AND hak_akses.status = ?", "umkm", 41, 50, "disetujui").
        Distinct("users.id").Count(&age41_50).Error
    if err != nil {
        return nil, err
    }

    // Menghitung jumlah pengguna dengan rentang umur 51-60
    err = repo.db.Table("users").
        Joins("JOIN hak_akses ON users.id = hak_akses.user_id").
        Where("role = ? AND date_part('year', age(tanggal_lahir)) BETWEEN ? AND ? AND hak_akses.status = ?", "umkm", 51, 60, "disetujui").
        Distinct("users.id").Count(&age51_60).Error
    if err != nil {
        return nil, err
    }

    // Menghitung jumlah pengguna dengan rentang umur 61-70
    err = repo.db.Table("users").
        Joins("JOIN hak_akses ON users.id = hak_akses.user_id").
        Where("role = ? AND date_part('year', age(tanggal_lahir)) BETWEEN ? AND ? AND hak_akses.status = ?", "umkm", 61, 70, "disetujui").
        Distinct("users.id").Count(&age61_70).Error
    if err != nil {
        return nil, err
    }

    // Menghitung jumlah pengguna dengan rentang umur 71-80
    err = repo.db.Table("users").
        Joins("JOIN hak_akses ON users.id = hak_akses.user_id").
        Where("role = ? AND date_part('year', age(tanggal_lahir)) BETWEEN ? AND ? AND hak_akses.status = ?", "umkm", 71, 80, "disetujui").
        Distinct("users.id").Count(&age71_80).Error
    if err != nil {
        return nil, err
    }

    // Menghitung jumlah pengguna dengan rentang umur 81-90
    err = repo.db.Table("users").
        Joins("JOIN hak_akses ON users.id = hak_akses.user_id").
        Where("role = ? AND date_part('year', age(tanggal_lahir)) BETWEEN ? AND ? AND hak_akses.status = ?", "umkm", 81, 90, "disetujui").
        Distinct("users.id").Count(&age81_90).Error
    if err != nil {
        return nil, err
    }

    // Mengembalikan hasil dalam bentuk map dengan jumlah pengguna berdasarkan rentang umur
    result := map[string]int64{
        "age11_20": age11_20,
        "age21_30": age21_30,
        "age31_40": age31_40,
        "age41_50": age41_50,
        "age51_60": age51_60,
        "age61_70": age61_70,
        "age71_80": age71_80,
        "age81_90": age81_90,
    }

    return result, nil
}
