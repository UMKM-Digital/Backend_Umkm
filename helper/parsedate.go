package helper

import (
    "time"
    "errors"
)

// ParseDate parses a date string in ISO 8601 format to a time.Time object
func ParseDate(dateStr string) (time.Time, error) {
    date, err := time.Parse(time.RFC3339, dateStr)
    if err != nil {
        return time.Time{}, errors.New("invalid date format, expected ISO 8601")
    }
    return date, nil
}
