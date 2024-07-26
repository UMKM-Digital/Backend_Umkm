package route

import (
	"net/http"
	"os"
	"umkm/app"
	kategoriumkmcontroller "umkm/controller/kategoriumkm"
	"umkm/controller/usercontroller"
	"umkm/helper"
	"umkm/model"
	repokategoriumkm "umkm/repository/kategori_umkm"
	"umkm/repository/userrepo"
	kategoriumkmservice "umkm/service/kategori_umkm"
	userservice "umkm/service/user"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func RegisterUserRoute(prefix string, e *echo.Echo) {
    db := app.DBConnection()
    token := helper.NewTokenUseCase()
  
    userAuthRepo := userrepo.NewAuthRepositoryImpl(db)
    userAuthService := userservice.Newauthservice(userAuthRepo, token)
    userAuthController := usercontroller.NewAuthController(userAuthService)
	userKategoriUmkmRepo := repokategoriumkm.NewKategoriUmkmRepositoryImpl(db)
	userKatgoriUmkmService := kategoriumkmservice.NewKategoriUmkmService(userKategoriUmkmRepo)
	userKategoriUmkmController := kategoriumkmcontroller.NewKategeoriUmkmController(userKatgoriUmkmService)
    g := e.Group(prefix)

    authRoute := g.Group("/auth")
    authRoute.POST("/register", userAuthController.Register)
    authRoute.POST("/login", userAuthController.Login)
    authRoute.POST("/send-otp", userAuthController.SendOtp)
    authRoute.POST("/logout", userAuthController.Logout, JWTProtection()) 

	meRoute := g.Group("/me")
	meRoute.GET("", userAuthController.View, JWTProtection())
	// meRoute.PUT("/update", userAuthController.Update, JWTProtection())

	KatUmkmRoute := g.Group("/kategori")
	KatUmkmRoute.POST("/umkm", userKategoriUmkmController.Create)
	KatUmkmRoute.GET("/list", userKategoriUmkmController.GetKategoriList)
	KatUmkmRoute.GET("/:id", userKategoriUmkmController.GetKategoriId)
	KatUmkmRoute.PUT("/umkm/:id", userKategoriUmkmController.UpdateKategoriId)
}

func JWTProtection() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(helper.JwtCustomClaims) 
		},
		SigningKey: []byte(os.Getenv("SECRET_KEY")),
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusUnauthorized,  model.ResponseToClient(http.StatusUnauthorized, "anda harus login untuk mengakses resource ini", nil))
		},
	})
}