// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"

// 	"umkm/helper"
// 	"umkm/route"

// 	"github.com/go-playground/validator/v10"
// 	"github.com/joho/godotenv"
// 	"github.com/labstack/echo/v4"
// )

// type CustomValidator struct {
//     validator *validator.Validate
// }

// func (cv *CustomValidator) Validate(i interface{}) error {
//     return cv.validator.Struct(i)
// }

// func main() {
//     if err := godotenv.Load(".env"); err != nil {
//         log.Fatal("Error loading .env file!")
//     }

//     r := echo.New()

//     // Register custom validator
//     r.Validator = &CustomValidator{validator: validator.New()}
// 	r.HTTPErrorHandler = helper.BindAndValidate

//     route.RegisterUserRoute("/user", r)

//     r.Logger.Fatal(r.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
// }


package main

import (
	"fmt"
	"log"
	"os"

	"umkm/helper"
	"umkm/route"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file!")
	}

	r := echo.New()

	// Register custom validator
	r.Validator = helper.NewCustomValidator()

	// Set custom error handler
	r.HTTPErrorHandler = helper.BindAndValidate

	route.RegisterUserRoute("/user", r)

	r.Logger.Fatal(r.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
