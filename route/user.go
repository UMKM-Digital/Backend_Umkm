package route

import (
	"fmt"
	"net/http"
	
	"umkm/app"
	kategoriumkmcontroller "umkm/controller/kategoriumkm"
	"umkm/controller/usercontroller"
	"umkm/helper"
	"umkm/model"
	repokategoriumkm "umkm/repository/kategori_umkm"
	"umkm/repository/userrepo"
	kategoriumkmservice "umkm/service/kategori_umkm"
	userservice "umkm/service/user"
	"github.com/labstack/echo/v4"
)

var tokenUseCase helper.TokenUseCase

func RegisterUserRoute(prefix string, e *echo.Echo) {
	db := app.DBConnection()
	tokenUseCase := helper.NewTokenUseCase()

	userAuthRepo := userrepo.NewAuthRepositoryImpl(db)
	userAuthService := userservice.Newauthservice(userAuthRepo, tokenUseCase)
	userAuthController := usercontroller.NewAuthController(userAuthService, tokenUseCase)

	userKategoriUmkmRepo := repokategoriumkm.NewKategoriUmkmRepositoryImpl(db)
	userKatgoriUmkmService := kategoriumkmservice.NewKategoriUmkmService(userKategoriUmkmRepo)
	userKategoriUmkmController := kategoriumkmcontroller.NewKategeoriUmkmController(userKatgoriUmkmService)

	g := e.Group(prefix)

	authRoute := g.Group("/auth")
	authRoute.POST("/register", userAuthController.Register)
	authRoute.POST("/login", userAuthController.Login)
	authRoute.POST("/send-otp", userAuthController.SendOtp)
	authRoute.POST("/logout", userAuthController.Logout, JWTProtection(tokenUseCase))

	meRoute := g.Group("/me")
	meRoute.GET("", userAuthController.View, JWTProtection(tokenUseCase))

	KatUmkmRoute := g.Group("/kategori")
	KatUmkmRoute.POST("/umkm", userKategoriUmkmController.Create, JWTProtection(tokenUseCase))
	KatUmkmRoute.GET("/list", userKategoriUmkmController.GetKategoriList, JWTProtection(tokenUseCase))
	KatUmkmRoute.GET("/:id", userKategoriUmkmController.GetKategoriId, JWTProtection(tokenUseCase))
	KatUmkmRoute.PUT("/umkm/:id", userKategoriUmkmController.UpdateKategoriId, JWTProtection(tokenUseCase))
	KatUmkmRoute.DELETE("/umkm/delete/:id", userKategoriUmkmController.DeleteKategoriId, JWTProtection(tokenUseCase))
}
func JWTProtection(tokenUseCase helper.TokenUseCase) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Request().Header.Get("Authorization")
			if token != "" {
				token = token[7:] // Remove "Bearer " prefix
			}

			if tokenUseCase.IsTokenBlacklisted(token) {
				// Tambahkan log untuk debugging
				fmt.Println("Token ditemukan di blacklist")
				return c.JSON(http.StatusUnauthorized, model.ResponseToClient(http.StatusUnauthorized, "Token sudah di-blacklist", nil))
			}

			// Jika token tidak di-blacklist, lanjutkan ke handler berikutnya
			return next(c)
		}
	}
}
