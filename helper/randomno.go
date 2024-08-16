package helper

import (
	"fmt"
	"strconv"
	"time"
	"umkm/model/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// GenerateInvoiceNumber generates an invoice number with the format INVYYYYMMDDNNN
func GenerateInvoiceNumber(db *gorm.DB, umkmID uuid.UUID) (string, error) {
	now := time.Now()
	yearMonthDay := now.Format("20060102") // Format as YYYYMMDD
	prefix := "INV" + yearMonthDay         // Prefix for the invoice number

	// Find the latest invoice number for the given UMKM ID
	var latestInvoiceNumber string
	result := db.Model(&domain.Transaksi{}).
		Where("umkm_id = ?", umkmID).
		Select("no_invoice").
		Order("no_invoice desc").
		Limit(1).
		Scan(&latestInvoiceNumber)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return "", result.Error
	}

	// Debug log for verification
	fmt.Printf("Latest Invoice Number from DB: %s\n", latestInvoiceNumber)

	var nextNumber string
	if latestInvoiceNumber == "" {
		// No previous invoice found, start with 001
		nextNumber = "001"
	} else {
		// Validate the format and extract number part
		if len(latestInvoiceNumber) < len(prefix)+3 { // Adjust for 3 digits
			return "", fmt.Errorf("invalid invoice number format: %s", latestInvoiceNumber)
		}
		// Extract number part
		numberPart := latestInvoiceNumber[len(prefix):]
		// Debug log for numberPart
		fmt.Printf("Number part extracted: %s\n", numberPart)

		currentNumber, err := strconv.Atoi(numberPart)
		if err != nil {
			return "", fmt.Errorf("error converting string to int: %s", numberPart)
		}
		nextNumber = fmt.Sprintf("%03d", currentNumber+1)
	}

	// Generate the new invoice number
	invoiceNumber := fmt.Sprintf("%s%s", prefix, nextNumber)
	fmt.Printf("Generated Invoice Number: %s\n", invoiceNumber) // Debug log

	return invoiceNumber, nil
}
