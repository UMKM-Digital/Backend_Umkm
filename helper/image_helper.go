package helper

import (
    "fmt"
    "io"
    "mime/multipart"
    "os"
    "path/filepath"
)

func HandleFileUpload(file *multipart.FileHeader, uploadDir string) (string, error) {
    // Open the file
    src, err := file.Open()
    if err != nil {
        return "", err
    }
    defer src.Close()

    // Create the upload directory if it does not exist
    if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
        return "", err
    }

    // Define the file path
    filename := file.Filename
    dst, err := os.Create(filepath.Join(uploadDir, filename))
    if err != nil {
        return "", err
    }
    defer dst.Close()

    // Copy the file content to the destination
    if _, err := io.Copy(dst, src); err != nil {
        return "", err
    }

    // Return the URL of the uploaded file
    return fmt.Sprintf("/uploads/%s", filename), nil
}
