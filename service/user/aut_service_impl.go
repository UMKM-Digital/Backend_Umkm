package userservice

import (
	// "encoding/json"
	// "encoding/json"
	"errors"
	"log"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"umkm/helper"
	"umkm/model/domain"
	"umkm/model/entity"

	"umkm/model/web"
	dokumenumkmrepo "umkm/repository/dokumenumkm"
	hakaksesrepo "umkm/repository/hakakses"
	omsetrepo "umkm/repository/omset"
	produkrepo "umkm/repository/produk"
	transaksirepo "umkm/repository/transaksi"
	umkmrepo "umkm/repository/umkm"
	"umkm/repository/userrepo"

	"fmt"
	"net/smtp"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthServiceImpl struct {
	authrepository userrepo.AuthUserRepo
    hakaksesrepository hakaksesrepo.CreateHakakses
    umkmrepository     umkmrepo.CreateUmkm
	tokenUseCase   helper.TokenUseCase
    produkRepository produkrepo.CreateProduk
    dokumenrepository dokumenumkmrepo.DokumenUmkmrRepo
    transaksiRepository transaksirepo.TransaksiRepo
    omsetrepository omsetrepo.OmsetRepo
	db             *gorm.DB // Tambahkan field ini
}

func Newauthservice(authrepository userrepo.AuthUserRepo, token helper.TokenUseCase, db *gorm.DB, hakaksesrepository hakaksesrepo.CreateHakakses,  umkmrepository  umkmrepo.CreateUmkm,  produkRepository produkrepo.CreateProduk,
    dokumenrepository dokumenumkmrepo.DokumenUmkmrRepo,transaksiRepository transaksirepo.TransaksiRepo,
    omsetrepository omsetrepo.OmsetRepo) *AuthServiceImpl {
	return &AuthServiceImpl{
		authrepository: authrepository,
        hakaksesrepository: hakaksesrepository,
        umkmrepository: umkmrepository,
        produkRepository: produkRepository,
        dokumenrepository: dokumenrepository,
        transaksiRepository: transaksiRepository,
        omsetrepository: omsetrepository,
		tokenUseCase:   token,
		db:             db,
	}
}

type KTPDocument struct {
    ID       int    `json:"id"`
    Document string `json:"document"`
}

type KTPData struct {
    URLs []KTPDocument `json:"urls"`
}


const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// generateRandomFileName generates a random file name with the format YYYYMMDD_randomString_HHMMSS.ext
func generateRandomFileName(ext string) string {
    // Dapatkan tanggal saat ini
    now := time.Now()
    // Format tahun, bulan, dan tanggal
    datePrefix := now.Format("20060102") // Format: YYYYMMDD

    // Buat string acak 10 karakter
    randomString := generateRandomString(16)

    // Format jam, menit, dan detik
    timeSuffix := now.Format("150405") // Format: HHMMSS

    // Gabungkan prefix tanggal, angka acak, dan suffix waktu
    return fmt.Sprintf("%s_%s_%s%s", datePrefix, randomString, timeSuffix, ext)
}

// generateRandomString generates a random string of specified length
func generateRandomString(length int) string {
    bytes := make([]byte, length)
    for i := range bytes {
        bytes[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(bytes)
}

// register
func (service *AuthServiceImpl) RegisterRequest(user web.RegisterRequest) (map[string]interface{}, error) {

    // Hash password menggunakan bcrypt
    passHash, errHash := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
    if errHash != nil {
        return nil, errHash
    }
    user.Password = string(passHash)

    // Parse TanggalLahir (Date of Birth) tanpa memanggil .String()
    tanggalLahirParsed, errDate := helper.ParseDateLahir(user.TanggalLahir)
    if errDate != nil {
        return nil, errDate
    }

       // Validasi panjang beberapa field sekaligus
       fieldsToValidate := []struct {
        FieldName      string
        FieldValue     string
        ExpectedLength int
    }{
        {"No NIK", user.No_Nik, 16},
        // {"No NIB", user.No_Nib, 13},
        {"No KK", user.No_KK, 16},
    }

    if err := helper.ValidateFieldsLength(fieldsToValidate); err != nil {
        return nil, err
    }

    // Membuat object User baru
    newUser := domain.Users{
        Fullname:           user.Fullname,
        Password:           user.Password,
        Email:              user.Email,
        Role:               "umkm",
        No_Phone:           user.No_Phone,
        Nik:                user.No_Nik,
        NoKk:               user.No_KK,
        // Nib:                user.No_Nib,
        TanggalLahir:       tanggalLahirParsed,
        JenisKelamin:       user.JenisKelamin,
        PendidikanTerakhir: user.PendidikanTerakhir,
        StatusMenikah:      user.StatusMenikah,
        Alamat:             user.Alamat,
        Provinsi:           user.Provinsi,
        Kabupaten:          user.Kabupaten,
        Kecamatan:          user.Kecamatan,
        Kelurahan:          user.Kelurahan,
		KodePos: user.KodePos,
        Rt:                 user.Rt,
        Rw:                 user.Rw,
    }

    // Menyimpan user ke database
    saveUser, errSaveUser := service.authrepository.RegisterRequest(newUser)
    if errSaveUser != nil {
        if strings.Contains(errSaveUser.Error(), "duplicate key value violates unique constraint") && strings.Contains(errSaveUser.Error(), "users_email_key") {
            return nil, fmt.Errorf("Email %s telah terdaftar sebelumnya pada sistem ini", user.Email)
        }
        return nil, errSaveUser
    }

    // Membuat claims untuk token JWT
    claims := helper.JwtCustomClaims{
        ID:      strconv.Itoa(saveUser.IdUser), // Menggunakan ID dari saveUser setelah disimpan ke DB
        Name:    saveUser.Username,
        Fullname: saveUser.Fullname,
        Email:   saveUser.Email,
        Phone:   saveUser.No_Phone,
        Role:    saveUser.Role,
        Picture: saveUser.Picture,
    }

    // Menghasilkan token JWT
    token, tokenErr := service.tokenUseCase.GenerateAccessToken(claims)
    if tokenErr != nil {
        return nil, tokenErr
    }

    // Menghitung waktu kedaluwarsa token
    expirationTime := time.Now().Add(1 * time.Hour).Format(time.RFC3339)

    // Mengembalikan token dan informasi user
    return map[string]interface{}{
        "token":      token,
        "expires_at": expirationTime, // Sertakan waktu kedaluwarsa yang sebenarnya
    }, nil
}


func (service *AuthServiceImpl) LoginRequest(username string, password string) (map[string]interface{}, error) {
	var user *domain.Users
	var err error

	
        if strings.Contains(username, "@") {
            user, err = service.authrepository.FindUserByUsername("", username, "")
            if err != nil {
                return nil, fmt.Errorf("Email tidak ditemukan") 
            } 
        }else if strings.HasPrefix(username, "08") || strings.HasPrefix(username, "62") {
            // Jika input adalah nomor telepon
            if len(username) < 10 {
                return nil, fmt.Errorf("Nomor telepon tidak valid")
            }
        
            var formattedPhone string
            if strings.HasPrefix(username, "08") {
                formattedPhone = "62" + username[1:]
            } else {
                formattedPhone = username
            }
        
            user, err = service.authrepository.FindUserByUsername("", "", formattedPhone)
            if err != nil {
                return nil, fmt.Errorf("Nomor telepon tidak ditemukan")
            }
        } else {
            user, err = service.authrepository.FindUserByUsername(username, "", "")
            if err != nil {
                return nil, fmt.Errorf("Username tidak ditemukan")
                
            }   
        }
        log.Print("user tidak ditemukan", user) 
        if checkPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); checkPassword != nil {
            return nil, errors.New("Password salah")
        }

	claims := helper.JwtCustomClaims{
		ID:      strconv.Itoa(user.IdUser),
        Fullname: user.Fullname,
		Name:    user.Username,
		Email:   user.Email,
		Phone:   user.No_Phone,
		Role:    user.Role,
		Picture: user.Picture,
        
	}

	isDataComplete := helper.IsUserDataComplete(*user)

    
	token, tokenErr := service.tokenUseCase.GenerateAccessToken(claims)
	if tokenErr != nil {
		return nil, tokenErr
	}

	// Hitung waktu kedaluwarsa token
	expirationTime := time.Now().Add(1 * time.Hour).Format(time.RFC3339)
	return map[string]interface{}{
		"token":      token,
		"expires_at": expirationTime, // Sertakan waktu kedaluwarsa yang sebenarnya
		"data": isDataComplete,
	}, nil
}

func (service *AuthServiceImpl) SendOtp(phone string) (map[string]interface{}, error) {
	// Temukan pengguna berdasarkan nomor telepon
	_, err := service.authrepository.FindUserByPhone(phone)
	if err != nil {
		return nil, errors.New("No Telepon tidak temukan!")
	}

	// Generate OTP

	// Tentukan waktu kadaluarsa OTP
	expirationTime := time.Now().Add(1 * time.Minute)

	// Kirim OTP melalui WhatsApp dan simpan ke database
	if err := helper.SendWhatsAppOTP(service.db, phone, expirationTime); err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"message":    "OTP sent successfully",
		"expires_at": expirationTime.Format(time.RFC3339),
	}, nil
}

// get profile
func (service *AuthServiceImpl) ViewMe(userId int) (entity.UserEntity, error) {
	user, err := service.authrepository.GetByID(userId)
	if err != nil {
		return entity.UserEntity{}, err
	}

	return entity.ToUserEntity(user), nil
}

// update profile
func (service *AuthServiceImpl) Update(Id int, req web.UpdateUserRequest, file *multipart.FileHeader, fileKTP *multipart.FileHeader, fileKK *multipart.FileHeader) (helper.ResponseToJson, error) {
    user, errUser := service.authrepository.GetByID(Id)

    // Parse TanggalLahir (Date of Birth)
    tanggalLahirParsed, errDate := helper.ParseDateLahir(req.TanggalLahir)
    if errDate != nil {
        return nil, errDate
    }
    fieldsToValidate := []struct {
        FieldName      string
        FieldValue     string
        ExpectedLength int
    }{
        {"No NIK", req.No_Nik, 16},
        // {"No NIB", req.No_Nib, 13},
        {"No KK", req.No_KK, 16},
    }

    if err := helper.ValidateFieldsLength(fieldsToValidate); err != nil {
        return nil, err
    }

    if errUser != nil {
        return nil, errUser
    }

    if req.Fullname != "" {
        user.Username = req.Fullname
    }
    if req.Email != "" {
        user.Email = req.Email
    }
    if req.No_Phone != "" {
        user.No_Phone = req.No_Phone
    }

    var Logo, KTPPath, KKPath  string

    if file != nil {
        // Check if the current picture is from Google (URL starting with https://lh3.googleusercontent.com/)
        if strings.HasPrefix(user.Picture, "https://lh3.googleusercontent.com/") {
            // Don't remove, just upload the new image
            src, err := file.Open()
            if err != nil {
                return nil, errors.New("failed to open the uploaded file")
            }
            defer src.Close()
    
            ext := filepath.Ext(file.Filename)
            randomFileName := generateRandomFileName(ext)
            Logo = filepath.Join("uploads/potoprofile", randomFileName)
    
            if err := helper.SaveFile(file, Logo); err != nil {
                return nil, errors.New("failed to save image")
            }
    
            // Normalize the file path
            Logo = filepath.ToSlash(Logo)
    
        } else if user.Picture == "" {
            // If no previous picture exists (user.Picture is empty), upload the new one
            src, err := file.Open()
            if err != nil {
                return nil, errors.New("failed to open the uploaded file")
            }
            defer src.Close()
    
            ext := filepath.Ext(file.Filename)
            randomFileName := generateRandomFileName(ext)
            Logo = filepath.Join("uploads/potoprofile", randomFileName)
    
            if err := helper.SaveFile(file, Logo); err != nil {
                return nil, errors.New("failed to save image")
            }
    
            // Normalize the file path
            Logo = filepath.ToSlash(Logo)
    
        } else if strings.HasPrefix(user.Picture, "uploads/potoprofile") {
            // If the current picture is in the local folder, remove the old image first
            err := os.Remove(user.Picture)
            if err != nil {
                return nil, errors.New("failed to remove old image")
            }
    
            // Then upload the new image
            src, err := file.Open()
            if err != nil {
                return nil, errors.New("failed to open the uploaded file")
            }
            defer src.Close()
    
            ext := filepath.Ext(file.Filename)
            randomFileName := generateRandomFileName(ext)
            Logo = filepath.Join("uploads/potoprofile", randomFileName)
    
            if err := helper.SaveFile(file, Logo); err != nil {
                return nil, errors.New("failed to save image")
            }
    
            // Normalize the file path
            Logo = filepath.ToSlash(Logo)
        }
    
    } else {
        // If no new file is uploaded, retain the existing picture (either URL or local file)
        Logo = user.Picture
    }

// Menangani file KTP
    // Handle KTP file upload
    if fileKTP != nil {
        // Hapus file KTP lama jika ada
    if user.Ktp != "" {
        err := os.Remove(user.Ktp)
        if err != nil {
            return nil, errors.New("failed to remove old KTP file")
        }
    }

    // Buka file KTP baru
    src, err := fileKTP.Open()
    if err != nil {
        return nil, errors.New("failed to open the uploaded KTP file")
    }
    defer src.Close()

    // Menghasilkan nama file acak untuk file KTP yang diunggah
    ext := filepath.Ext(fileKTP.Filename)
    randomFileName := generateRandomFileName("Yk" + ext)
    KTPPath = filepath.Join("uploads", "dokpribadi", randomFileName)

    // Menyimpan file KTP ke server
    if err := helper.SaveFile(fileKTP, KTPPath); err != nil {
        return nil, errors.New("failed to save KTP file")
    }

    // Mengonversi path untuk menggunakan forward slashes
    KTPPath = filepath.ToSlash(KTPPath)
    } else {
        // Retain existing KTP path if no new file
        KTPPath = user.Ktp
    }

    // Handle KK file upload
    if fileKK != nil {
       // Hapus file KK lama jika ada
    if user.KartuKeluarga != "" {
        err := os.Remove(user.KartuKeluarga)
        if err != nil {
            return nil, errors.New("failed to remove old KK file")
        }
    }

    // Buka file KK baru
    src, err := fileKK.Open()
    if err != nil {
        return nil, errors.New("failed to open the uploaded KK file")
    }
    defer src.Close()

    // Menghasilkan nama file acak untuk file KK yang diunggah
    ext := filepath.Ext(fileKK.Filename)
    randomFileName := generateRandomFileName("GI" + ext)
    KKPath = filepath.Join("uploads", "dokpribadi", randomFileName)

    // Menyimpan file KK ke server
    if err := helper.SaveFile(fileKK, KKPath); err != nil {
        return nil, errors.New("failed to save KK file")
    }

    // Mengonversi path untuk menggunakan forward slashes
    KKPath = filepath.ToSlash(KKPath)
    } else {
        // Retain existing KK path if no new file
        KKPath = user.KartuKeluarga
    }


    // Create user update data
    TestimonalRequest := domain.Users{
        IdUser:          Id,
        Fullname:        req.Fullname,
        Email:           req.Email,
        No_Phone:        req.No_Phone,
        Nik:             req.No_Nik,
        NoKk:            req.No_KK,
        // Nib:             req.No_Nib,
        TanggalLahir:    tanggalLahirParsed,
        JenisKelamin:    req.JenisKelamin,
        StatusMenikah:   req.StatusMenikah,
        Alamat:          req.Alamat,
        Provinsi:        req.Provinsi,
        Kabupaten:       req.Kabupaten,
        Kecamatan:       req.Kecamatan,
        Kelurahan:       req.Kelurahan,
        Rt:              req.Rt,
        Rw:              req.Rw,
        PendidikanTerakhir: req.PendidikanTerakhir,
        KodePos:         req.KodePos,
        Ktp:             KTPPath,
        KartuKeluarga:   KKPath,
        Picture: Logo,
    }

    result, errUpdate := service.authrepository.UpdateId(Id, TestimonalRequest)
    if errUpdate != nil {
        return nil, errUpdate
    }

    response := map[string]interface{}{
        "name":   result.Fullname,
        "email":  result.Email,
        "phone":  result.No_Phone,
    }
    return response, nil
}


//

// verify
func (service *AuthServiceImpl) VerifyOTP(phone_number string, otpCode string) (map[string]interface{}, error) {
	// Verifikasi OTP
	isValid, err := helper.VerifyOTP(service.db, phone_number, otpCode)
	if err != nil || !isValid {
		return nil, errors.New("OTP tidak sesuai")
	}

	// Temukan pengguna berdasarkan nomor telepon
	user, err := service.authrepository.FindUserByPhone(phone_number)
	if err != nil {
		return nil, errors.New("user not found")
	}

	isDataComplete := helper.IsUserDataComplete(*user)

	// Token
	claims := helper.JwtCustomClaims{
		ID:      strconv.Itoa(user.IdUser),
		Name:    user.Username,
        Fullname: user.Fullname,
		Email:   user.Email,
		Phone:   user.No_Phone,
		Picture: user.Picture,
		Role:    user.Role,
	}

	token, tokenErr := service.tokenUseCase.GenerateAccessToken(claims)
	if tokenErr != nil {
		return nil, tokenErr
	}

	// Hitung waktu kedaluwarsa token
	expirationTime := time.Now().Add(1 * time.Hour).Format(time.RFC3339)

	// OTP berhasil diverifikasi dan password benar
	return map[string]interface{}{
		"message": "Kode OTP sesuai",
		// "user":    user,
		"token":        token,
		"expired time": expirationTime,
		"data": isDataComplete,
	}, nil
}

func (service *AuthServiceImpl) SendOtpRegister(phone string) (map[string]interface{}, error) {
    if phone == "" {
        return nil, errors.New("No Telepon kosong")
    }

    user, err := service.authrepository.FindUserByPhoneRegister(phone)
    if err != nil {
        return nil, err
    }

    if user != nil {
        return nil, errors.New("No Telepon telah terdaftar") // Mengembalikan error jika nomor sudah terdaftar
    }

    expirationTime := time.Now().Add(1 * time.Minute)

    if err := helper.SendWhatsAppOTP(service.db, phone, expirationTime); err != nil {
        return nil, err
    }

    return map[string]interface{}{
        "message":    "OTP terkirim",
        "expires_at": expirationTime.Format(time.RFC3339),
    }, nil
}


// verify otp register
func (service *AuthServiceImpl) VerifyOTPRegister(otp_code string, phone_code string) (map[string]interface{}, error) {
	// Verifikasi OTP
	isValid, err := helper.VerifyOTP(service.db, otp_code, phone_code)
	if err != nil || !isValid {
		return nil, errors.New("OTP tidak sesuai")
	}

	return map[string]interface{}{
		"message": "Kode OTP sesuai",
	}, nil
}

func (service *AuthServiceImpl) CekInRequest(authID int, password string) (map[string]interface{}, error) {
    user, getUserErr := service.authrepository.CekInPassword(authID)
    if getUserErr != nil {
        return nil, errors.New("user not found")
    }

    if checkPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); checkPassword != nil {
        return nil, errors.New("incorrect password")
    }

    // Return the user details or any necessary information upon successful password check
    return map[string]interface{}{
        "id":    user.IdUser,
        "name":  user.Username,
        "email": user.Email,
    }, nil
}


//updatepasssword
func (service *AuthServiceImpl) ChangePassword(authID int, oldPassword string, newPassword string) error {
    // Dapatkan user berdasarkan authID
    user, getUserErr := service.authrepository.CekInPassword(authID)
    if getUserErr != nil {
        return errors.New("user not found")
    }

    // Verifikasi apakah password lama sesuai
    if checkPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword)); checkPassword != nil {
        return errors.New("incorrect old password")
    }

    // Ubah ke password baru
    if err := service.authrepository.UpdatePassword(authID, newPassword); err != nil {
        return err
    }

    return nil
}


func (service *AuthServiceImpl) HandleGoogleLoginOrRegister(googleID string, email string, username string, picture string) (map[string]interface{}, error) {
	// Mencari atau membuat pengguna berdasarkan ID Google
	user, err := service.authrepository.FindOrCreateUserByGoogleID(googleID, email, username, picture)
	if err != nil {
		return nil, err
	}

	// Jika pengguna baru dibuat, kita hash password kosong (opsional)
	if user.GoogleId != "" {
		// Hash password menggunakan bcrypt (opsional)
		passHash, errHash := bcrypt.GenerateFromPassword([]byte("temporaryPassword"), bcrypt.MinCost) // Password sementara
		if errHash != nil {
			return nil, errHash
		}
		user.Password = string(passHash) // Simpan password yang di-hash
	}

	isDataComplete := helper.IsUserDataComplete(*user)


	// Membuat claims untuk token JWT
	claims := helper.JwtCustomClaims{
		ID:      strconv.Itoa(user.IdUser),
		Name:    user.Fullname,
		Email:   user.Email,
		Phone:   user.No_Phone,
		Role:    user.Role,
		Picture: user.Picture,
	}

	// Menghasilkan token JWT
	token, tokenErr := service.tokenUseCase.GenerateAccessToken(claims)
	if tokenErr != nil {
		return nil, tokenErr
	}

	// Menghitung waktu kedaluwarsa token
	expirationTime := time.Now().Add(1 * time.Hour).Format(time.RFC3339)

	// Mengembalikan token dan informasi pengguna
	return map[string]interface{}{
		"token":      token,
		"expires_at": expirationTime, // Sertakan waktu kedaluwarsa yang sebenarnya
		"data": isDataComplete,
	}, nil
}


// SendPasswordResetLink mengirimkan link reset password ke email yang diberikan
func (service *AuthServiceImpl) SendPasswordResetLink(email string) error {
	// Cek apakah email ada di database
	user, err := service.authrepository.ChangePassword(email)
	if err != nil {
		return fmt.Errorf("email not found: %v", err)
	}

	// Buat token reset password (misalnya JWT atau random string)
	resetToken := "23214jadkhb" // Ganti dengan fungsi untuk membuat token yang aman

	// Buat link reset password
	resetLink := fmt.Sprintf("https://yourapp.com/reset-password?token=%s", resetToken)

	// Waktu expired link (misal 1 jam)
	expirationTime := time.Now().Add(1 * time.Hour)

	// Simpan token dan waktu expired ke database jika diperlukan (opsional)

	// Kirim email dengan SMTP
	err = service.sendEmail(user.Email, resetLink, expirationTime)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}

func (service *AuthServiceImpl) sendEmail(recipientEmail, resetLink string, expirationTime time.Time) error {
	from := "adlisantosanaufal@gmail.com"
	password := "bnpp toam ocrd yuid"

	to := []string{recipientEmail}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte(fmt.Sprintf(
		"Subject: Password Reset Request\n\nClick the following link to reset your password:\n%s\n\nThis link will expire at %s.",
		resetLink, expirationTime.Format("02 Jan 2006 15:04")))

	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}

func (service *AuthServiceImpl) GetListUser() ([]entity.UserEntityList, error) {
    GetUserList, err := service.authrepository.ListUser()
    if err != nil {
        return nil, err
    }
    return entity.ToUserEntitiesList(GetUserList), nil
}


//visualisai data
// Visualisasi data
func (service *AuthServiceImpl) CountUser() (map[string]interface{}, error) {
    // Memanggil fungsi dari repository untuk menghitung persentase gender
    genderResult, err := service.authrepository.CountUserByGenderWithPercentage()
    if err != nil {
        return nil, err
    }

    // Memanggil fungsi dari repository untuk menghitung jumlah berdasarkan pendidikan terakhir
    studyResult, err := service.authrepository.CountUserByStudy()
    if err != nil {
        return nil, err
    }

    // Memanggil fungsi dari repository untuk menghitung jumlah berdasarkan umur
    ageResult, err := service.authrepository.CountUserByAge()
    if err != nil {
        return nil, err
    }

    // Menggabungkan hasil gender, study, dan age dalam satu map
    result := map[string]interface{}{
        "gender": genderResult,
        "study":  studyResult,
        "age":    ageResult,
    }

    // Menambahkan struktur respons sesuai dengan yang Anda inginkan

    return result, nil
}


func(service *AuthServiceImpl) DeleteUser(iduser int) error{
    // Cek apakah user ada
    _, err := service.authrepository.GetByID(iduser)
    if err != nil {
        fmt.Println("Error getting user:", err)
        return err
    }
    
    umkmIDs, err := service.hakaksesrepository.GetUmkmIdsByUserId(iduser)
    if err != nil {
        fmt.Println("Error getting UMKM IDs:", err)
        return fmt.Errorf("gagal mendapatkan daftar UMKM: %w", err)
    }
    fmt.Println("UMKM IDs to delete:", umkmIDs) // Log daftar UMKM IDs

    if err := service.hakaksesrepository.DeleteUser(iduser); err != nil {
        fmt.Println("Error deleting user from hakaksesrepository:", err)
        return err
    }

    // Hapus semua UMKM dari tabel umkm berdasarkan ID
    for _, umkmID := range umkmIDs {
        fmt.Println("Deleting UMKM with ID:", umkmID)
        if err := service.umkmrepository.DeleteUmkmId(umkmID); err != nil {
            fmt.Println("Error deleting UMKM:", err)
            return fmt.Errorf("gagal menghapus UMKM dengan ID %d: %w", umkmID, err)
        }

        fmt.Println("Deleting Produk with UMKM ID:", umkmID)
        if err := service.produkRepository.DeleteProdukUmkmId(umkmID); err != nil {
            fmt.Println("Error deleting Produk:", err)
            return fmt.Errorf("gagal menghapus produk dengan ID %d: %w", umkmID, err)
        }

        fmt.Println("Deleting Dokumen with UMKM ID:", umkmID)
        if err := service.dokumenrepository.DeleteDokumenUmkmId(umkmID); err != nil {
            fmt.Println("Error deleting Dokumen:", err)
            return err
        }

        fmt.Println("Deleting Omset with UMKM ID:", umkmID)
        if err := service.omsetrepository.DeleteUserOmzet(umkmID); err != nil {
            fmt.Println("Error deleting Omset:", err)
            return err
        }

        fmt.Println("Deleting Transaksi with UMKM ID:", umkmID)
        if err := service.transaksiRepository.DeleteTransaksiUmkmId(umkmID); err != nil {
            fmt.Println("Error deleting Transaksi:", err)
            return err
        }
    }

    fmt.Println("Deleting User with ID:", iduser)
    if err := service.authrepository.DeleteUser(iduser); err != nil {
        fmt.Println("Error deleting user from authrepository:", err)
        return err
    }

    fmt.Println("User deleted successfully")
    return nil
}
