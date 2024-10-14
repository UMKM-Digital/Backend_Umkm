package helper

import (
    "time"
    "errors"
    "fmt"
)

// ParseDate parses a date string in ISO 8601 format to a time.Time object
func ParseDate(dateStr string) (time.Time, error) {
    date, err := time.Parse(time.RFC3339, dateStr)
    if err != nil {
        return time.Time{}, errors.New("invalid date format, expected ISO 8601")
    }
    return date, nil
}



func ParseDateLahir(dateStr string) (time.Time, error) {
    layout := "2006-01-02"
    parsedDate, err := time.Parse(layout, dateStr)
    if err != nil {
        return time.Time{}, fmt.Errorf("invalid date format, expected YYYY-MM-DD: %v", err)
    }
    return parsedDate, nil
}
