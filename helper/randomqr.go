package helper

import (
    "crypto/rand"
    "math/big"
)
// GenerateValidationTicket generates a random 32-character validation ticket
func GenerateValidationTicket() (string, error) {
    const length = 32
    const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    ticket := make([]byte, length)
    
    for i := range ticket {
        randIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
        if err != nil {
            return "", err
        }
        ticket[i] = charset[randIndex.Int64()]
    }
    
    return string(ticket), nil
}
