package main

import (
    "fmt"
    "log"
    "os"

    "umkm/helper"
    "umkm/route"

    "github.com/joho/godotenv"
    "github.com/labstack/echo/v4"
    // "github.com/labstack/echo/v4/middleware"
)

func main() {
    if err := godotenv.Load(".env"); err != nil {
        log.Fatal("Error loading .env file!")
    }

    r := echo.New()

    // Daftarkan validator kustom
    r.Validator = helper.NewCustomValidator()

    // Atur handler kesalahan kustom
    r.HTTPErrorHandler = helper.BindAndValidate


    route.RegisterUserRoute("/user", r)
    r.Static("/uploads", "uploads")
    r.Logger.Fatal(r.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
    
    r.GET("/uploads/*", func(c echo.Context) error {
        path := c.Param("*")
        log.Printf("Requested path: %s", path)
        return c.File("/Backend_Umkm/uploads/" + path)
    })
    
}
