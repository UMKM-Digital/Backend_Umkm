package helper

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
)

const otpLength = 6

// GenerateOTP generates a random OTP
func GenerateOTP() (string, error) {
	otp, err := generateRandomOTP(otpLength)
	if err != nil {
		return "", err
	}
	return otp, nil
}

// GenerateRandomOTP generates a random OTP of given length
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

// SendWhatsAppOTP sends OTP via WhatsApp
func SendWhatsAppOTP(phone, otp string) error {
	url := "http://103.162.60.86:3000/kirim-pesan"

	requestData := map[string]interface{}{
		"app_token_id": "ea5be31c-d4ec-4d44-8163-948a8e528ff6",
		"service":      "whatsapp",
		"penerima":     phone,
		"konten":       "kode OTP Anda adalah " + otp,
		"optional_data": map[string]string{
			"callback": "http://localhost:8001/user/auth/send-otp",
		},
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return err
	}

	_, err = http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	return nil
}

// ResponseToJson standardizes the response format

