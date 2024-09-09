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
    "github.com/labstack/echo/v4/middleware"
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
	r.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // Ganti dengan domain frontend kamu
		AllowMethods: []string{
			echo.GET,
			echo.POST,
			echo.PUT,
			echo.DELETE,
			echo.OPTIONS,
		},
		AllowHeaders: []string{
			echo.HeaderContentType,
			echo.HeaderAuthorization,
			echo.HeaderAccept,
		},
	}))

    route.RegisterUserRoute("/api", r)
    r.Static("/uploads", "uploads")
    r.Static("/uploads/logo", "uploads/logo")
    r.Logger.Fatal(r.Start(fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT"))))
    
    r.GET("/uploads/*", func(c echo.Context) error {
        path := c.Param("*")
        log.Printf("Requested path: %s", path)
        return c.File("/Backend_Umkm/uploads/" + path)
    })
    
    r.GET("/uploads/logo/*", func(c echo.Context) error {
        path := c.Param("*")
        log.Printf("Requested path: %s", path)
        return c.File("/Backend_Umkm/uploads/logo/" + path)
    })

    r.GET("/uploads/about/*", func(c echo.Context) error {
        path := c.Param("*")
        log.Printf("Requested path: %s", path)
        return c.File("/Backend_Umkm/uploads/about/" + path)
    })
    
}
