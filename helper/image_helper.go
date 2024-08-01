// package helper

// import (
// 	"errors"
// 	"fmt"

// 	"io"
// 	"os"
// 	"path/filepath"

// 	"github.com/labstack/echo/v4"
// )

// func RemoveFile(filePath string) error {
// 	if _, err := os.Stat(filePath); err != nil {
// 		if os.IsNotExist(err) {
// 			return fmt.Errorf("File not found")
// 		}
// 		return err
// 	}

// 	if err := os.Remove(filePath); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func HandleFileUpload(c *echo.Context, uploadPath string, formKey string) (string, error) {
// 	form, err := c.mu

// 	if err != nil {
// 		return "", err
// 	}

// 	files := form.File[formKey]
// 	if len(files) == 0 {
// 		return "", errors.New("No file uploaded")
// 	}

// 	if files == nil {
// 		return "", nil
// 	} else {
// 		file := files[0]
// 		filename := file.Filename

// 		destinationFile, err := os.Create(uploadPath + filename)
// 		if err != nil {
// 			return "", err
// 		}
// 		defer destinationFile.Close()

// 		sourceFile, err := file.Open()
// 		if err != nil {
// 			return "", err
// 		}
// 		defer sourceFile.Close()

// 		_, err = io.Copy(destinationFile, sourceFile)
// 		if err != nil {
// 			return "", err
// 		}

// 		return filename, nil
// 	}
// }

// func IsDataInDirectory(directoryPath, targetData string) (bool, error) {
// 	files, err := filepath.Glob(filepath.Join(directoryPath, "*"))
// 	if err != nil {
// 		return false, err
// 	}

// 	for _, file := range files {
// 		if filepath.Base(file) == targetData {
// 			return true, nil
// 		}
// 	}

// 	return false, nil
// }