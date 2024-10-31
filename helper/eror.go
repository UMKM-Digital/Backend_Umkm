package helper

import "fmt"

func ValidateLength(field string, value string, expectedLength int) error {
    if len(value) != expectedLength {
        return fmt.Errorf("%s harus terdiri dari %d angka", field, expectedLength)
    }
    return nil
}

func ValidateFieldsLength(fields []struct {
    FieldName      string
    FieldValue     string
    ExpectedLength int
}) error {
    for _, field := range fields {
        if err := ValidateLength(field.FieldName, field.FieldValue, field.ExpectedLength); err != nil {
            return err
        }
    }
    return nil
}