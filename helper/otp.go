package helper

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

const otpLength = 6

type SaveOtp struct {
	PhoneNumber string    `gorm:"column:phone_number" json:"phone_number"`
	Otp         string    `gorm:"column:otp_code" json:"otp_code"`
	Status      bool      `gorm:"column:status" json:"status"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	ExpiresAt   time.Time `gorm:"column:expires_at" json:"expires_at"`
}



// GenerateOTP generates a random OTP
func GenerateOTP() (string, error) {
	otp, err := generateRandomOTP(otpLength)
	if err != nil {
		return "", err
	}
	return otp, nil
}

// generateRandomOTP generates a random OTP of given length
func generateRandomOTP(length int) (string, error) {
	var otp string
	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return "", err
		}
		otp += fmt.Sprintf("%d", num)
	}
	return otp, nil
}


func init() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        fmt.Println("Error loading .env file")
    }
}

func SendWhatsAppOTP(db *gorm.DB, phone string, expiresAt time.Time) error {
	otp, err := GenerateOTP()
	if err != nil {
		return err
	}

	// Save OTP to the database
	err = saveOTP(db, phone, otp, expiresAt)
	if err != nil {
		return err
	}

	url := os.Getenv("WHATSAPP_API_URL")
    appTokenID := os.Getenv("WHATSAPP_APP_TOKEN_ID")
	requestData := map[string]interface{}{
		"app_token_id": appTokenID,
		"service":      "whatsapp",
		"penerima":     phone,
		"konten":       "kode OTP Anda adalah " + otp,
		"optional_data": map[string]string{
			"callback": "https://demo-umkm.simda.net/",
		},
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return err
	}

	// Mengirim request ke API
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		// Mengubah pesan error sesuai permintaan
		return fmt.Errorf("Metode login WA sedang mengalami gangguan.")
	}
	defer resp.Body.Close()

	return nil
}


// saveOTP saves OTP to the database
// Simpan OTP
func saveOTP(db *gorm.DB, phoneNumber, otpCode string, expiresAt time.Time) error {
    // Hapus OTP yang sudah kadaluarsa
    db.Where("phone_number = ? AND expires_at < ?", phoneNumber, time.Now()).Delete(&SaveOtp{})

    // Buat atau perbarui OTP
    otp := SaveOtp{
        PhoneNumber: phoneNumber,
        Otp:         otpCode,
        Status:      false,
        CreatedAt:   time.Now(),
        ExpiresAt:   expiresAt,
    }

    // Cek apakah ada OTP yang masih berlaku
    result := db.Where("phone_number = ? AND expires_at >= ?", phoneNumber, time.Now()).First(&otp)
    if result.RowsAffected > 0 {
        return fmt.Errorf("OTP untuk nomor ini masih berlaku")
    }

    // Jika tidak ada OTP yang berlaku, buat yang baru
    return db.Create(&otp).Error
}


func VerifyOTP(db *gorm.DB, phoneNumber, otpCode string) (bool, error) {
    var otp SaveOtp
    result := db.Where("phone_number = ? AND otp_code = ?", phoneNumber, otpCode).
        Order("created_at DESC").
        First(&otp)

    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            fmt.Printf("OTP tidak ditemukan untuk nomor telepon %s dan kode OTP %s\n", phoneNumber, otpCode)
            return false, fmt.Errorf("OTP tidak ditemukan untuk nomor telepon %s", phoneNumber)
        }
        fmt.Printf("Error saat mencari OTP: %v\n", result.Error)
        return false, result.Error
    }

    fmt.Printf("OTP Found: %+v\n", otp)

    if otp.Status {
        return false, fmt.Errorf("OTP sudah digunakan")
    }

    if otpCode != otp.Otp {
        return false, fmt.Errorf("Kode OTP tidak valid")
    }

    if time.Now().After(otp.ExpiresAt) {
        return false, fmt.Errorf("OTP sudah kadaluarsa")
    }

    if err := db.Model(&otp).Where("phone_number = ? AND otp_code = ?", phoneNumber, otpCode).Update("status", true).Error; err != nil {
        fmt.Printf("Error saat memperbarui status OTP: %v\n", err)
        return false, err
    }

    return true, nil
}
