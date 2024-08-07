package route

import (
	"net/http"
	"os"

	"umkm/app"
	kategoriprodukcontroller "umkm/controller/kategoriproduk"
	kategoriumkmcontroller "umkm/controller/kategoriumkm"
	transaksicontroller "umkm/controller/transaksi"
	umkmcontroller "umkm/controller/umkm"
	"umkm/controller/usercontroller"
	"umkm/helper"
	"umkm/model"
	// querybuildertransaksi "umkm/query_builder/transaksi"
	hakaksesrepo "umkm/repository/hakakses"
	kategoriprodukrepo "umkm/repository/kategori_produk"
	repokategoriumkm "umkm/repository/kategori_umkm"
	transaksirepo "umkm/repository/transaksi"
	umkmrepo "umkm/repository/umkm"
	"umkm/repository/userrepo"
	kategoriprodukservice "umkm/service/kategori_produk"
	kategoriumkmservice "umkm/service/kategori_umkm"
	transaksiservice "umkm/service/transaksi"
	umkmservice "umkm/service/umkm"
	userservice "umkm/service/user"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

// var tokenUseCase helper.TokenUseCase

func RegisterUserRoute(prefix string, e *echo.Echo) {
	db := app.DBConnection()
	tokenUseCase := helper.NewTokenUseCase()

	userAuthRepo := userrepo.NewAuthRepositoryImpl(db)
	userAuthService := userservice.Newauthservice(userAuthRepo, tokenUseCase, db)
	userAuthController := usercontroller.NewAuthController(userAuthService, tokenUseCase)

	userKategoriUmkmRepo := repokategoriumkm.NewKategoriUmkmRepositoryImpl(db)
	userKatgoriUmkmService := kategoriumkmservice.NewKategoriUmkmService(userKategoriUmkmRepo)
	userKategoriUmkmController := kategoriumkmcontroller.NewKategeoriUmkmController(userKatgoriUmkmService)

	userUmkmRepo := umkmrepo.NewUmkmRepositoryImpl(db)
	userHakAksesRepo := hakaksesrepo.NewHakAksesRepositoryImpl(db) // Tambahkan repository HakAkses
	userUmkmService := umkmservice.NewUmkmService(userUmkmRepo, userHakAksesRepo)
	userUmkmController := umkmcontroller.NewUmkmController(userUmkmService)

	//  
	userTransaksiRepo := transaksirepo.NewTransaksiRepositoryImpl(db)
	userTransaksiService := transaksiservice.NewTransaksiservice(userTransaksiRepo)
	userTransaksiController := transaksicontroller.NewUmkmController(userTransaksiService)

	userKategoriProdukRepo := kategoriprodukrepo.NewKategoriProdukRepo(db)
	userKategoriProdukService := kategoriprodukservice.NewKategoriProdukService(userKategoriProdukRepo)
	userKategoriProdukController := kategoriprodukcontroller.NewKategeoriProdukController(*userKategoriProdukService)

	g := e.Group(prefix)

	authRoute := g.Group("/auth")
	authRoute.POST("/register", userAuthController.Register)
	authRoute.POST("/login", userAuthController.Login)
	authRoute.POST("/send-otp", userAuthController.SendOtp)
	authRoute.POST("/verifyOtp", userAuthController.VerifyOTPHandler)
	authRoute.POST("/logout", userAuthController.Logout, JWTProtection())

	meRoute := g.Group("/me")
	meRoute.GET("", userAuthController.View, JWTProtection())

	KatUmkmRoute := g.Group("/kategori")
	KatUmkmRoute.POST("/umkm", userKategoriUmkmController.Create, JWTProtection())
	KatUmkmRoute.GET("/list", userKategoriUmkmController.GetKategoriList, JWTProtection())
	 
	KatUmkmRoute.PUT("/umkm/:id", userKategoriUmkmController.UpdateKategoriId, JWTProtection())
	KatUmkmRoute.DELETE("/umkm/delete/:id", userKategoriUmkmController.DeleteKategoriId, JWTProtection())

	Umkm := g.Group("/create")
	Umkm.Static("/uploads", "uploads")

	Umkm.POST("/umkm", userUmkmController.Create, JWTProtection())

	Transaksi := g.Group("/transaksi")
	Transaksi.POST("/umkm", userTransaksiController.Create)
	Transaksi.GET("/:id", userTransaksiController.GetKategoriId)

	KatProdukRoute := g.Group("/kategoriproduk")
	KatProdukRoute.POST("/poost", userKategoriProdukController.Create)
	KatProdukRoute.GET("/:umkm_id", userKategoriProdukController.GetKategoriList)
}

func JWTProtection() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(helper.JwtCustomClaims)
		},
		SigningKey: []byte(os.Getenv("SECRET_KEY")),
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusUnauthorized, model.ResponseToClient(http.StatusUnauthorized, "unauthorized", nil))
		},
	})
}

// func JWTProtection(tokenUseCase helper.TokenUseCase) echo.MiddlewareFunc {
// 	return func(next echo.HandlerFunc) echo.HandlerFunc {
// 		return func(c echo.Context) error {

// 			token := c.Request().Header.Get("Authorization")
// 			if token != "" {
// 				token = token[7:] // Remove "Bearer " prefix
// 			}

// 			if tokenUseCase.IsTokenBlacklisted(token) {
// 				// Tambahkan log untuk debugging
// 				fmt.Println("Token ditemukan di blacklist")
// 				return c.JSON(http.StatusUnauthorized, model.ResponseToClient(http.StatusUnauthorized, "Token sudah di-blacklist", nil))
// 			}

// 			// Jika token tidak di-blacklist, lanjutkan ke handler berikutnya
// 			return next(c)
// 		}
// 	}
// }

// func JWTProtection(tokenUseCase helper.TokenUseCase) echo.MiddlewareFunc {
// 	return func(next echo.HandlerFunc) echo.HandlerFunc {
// 		return func(c echo.Context) error {
// 			token := c.Request().Header.Get("Authorization")

// 			if token == "" {
// 				// Jika header Authorization tidak ada, kembalikan status Unauthorized
// 				return c.JSON(http.StatusUnauthorized, model.ResponseToClient(http.StatusUnauthorized, "Token tidak ditemukan", nil))
// 			}

// 			// Hapus awalan "Bearer " jika ada
// 			if len(token) > 7 && token[:7] == "Bearer " {
// 				token = token[7:]
// 			} else {
// 				// Jika format token tidak sesuai, kembalikan status Unauthorized
// 				return c.JSON(http.StatusUnauthorized, model.ResponseToClient(http.StatusUnauthorized, "Format token tidak valid", nil))
// 			}

// 			if tokenUseCase.IsTokenBlacklisted(token) {
// 				// Tambahkan log untuk debugging
// 				fmt.Println("Token ditemukan di blacklist")
// 				return c.JSON(http.StatusUnauthorized, model.ResponseToClient(http.StatusUnauthorized, "Token sudah di-blacklist", nil))
// 			}

// 			// Jika token tidak di-blacklist, lanjutkan ke handler berikutnya
// 			return next(c)
// 		}
// 	}
// }
