package helper

import (
    "fmt"
    "math/rand"
    "time"
)

// GenerateInvoiceNumber generates an invoice number with the format INV<YYYY><BULAN><NOACK>
func GenerateInvoiceNumber() (string, error) {
    now := time.Now()
    year := now.Format("2006")
    month := now.Format("01")
    noAck := generateRandomNumber(3) // Generate a 3-digit number
    
    return fmt.Sprintf("INV%s%s%s", year, month, noAck), nil
}

// generateRandomNumber generates a random number with a fixed number of digits
func generateRandomNumber(digits int) string {
    rand.Seed(time.Now().UnixNano()) // Seed the random number generator
    min := 1
    max := 9
    for i := 1; i < digits; i++ {
        max *= 10
    }
    number := rand.Intn(max-min+1) + min
    return fmt.Sprintf("%0*d", digits, number)
}
