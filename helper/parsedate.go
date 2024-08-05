package helper

import (
    "time"
    "errors"
)

// ParseDate parses a date string in the format "YYYY-MM-DD" to a time.Time object
func ParseDate(dateStr string) (time.Time, error) {
    date, err := time.Parse("2006-01-02", dateStr)
    if err != nil {
        return time.Time{}, errors.New("invalid date format, expected YYYY-MM-DD")
    }
    return date, nil
}
