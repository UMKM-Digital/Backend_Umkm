package route

import (
	"net/http"
	"os"

	"umkm/app"
	// homepagecontroller "umkm/controller/homepage"
	// brandlogo "umkm/controller/homepage/logo"
	kategoriprodukcontroller "umkm/controller/kategoriproduk"
	kategoriumkmcontroller "umkm/controller/kategoriumkm"
	// produkcontroller "umkm/controller/produk"
	// transaksicontroller "umkm/controller/transaksi"
	umkmcontroller "umkm/controller/umkm"
	"umkm/controller/usercontroller"
	"umkm/helper"
	"umkm/model"

	// querybuilder "umkm/query_builder"

	// querybuildertransaksi "umkm/query_builder/transaksi"
	hakaksesrepo "umkm/repository/hakakses"
	// testimonialrepo "umkm/repository/homepage"
	// brandrepo "umkm/repository/homepage/brandlogo"
	kategoriprodukrepo "umkm/repository/kategori_produk"
	repokategoriumkm "umkm/repository/kategori_umkm"
	// produkrepo "umkm/repository/produk"
	// transaksirepo "umkm/repository/transaksi"
	umkmrepo "umkm/repository/umkm"
	"umkm/repository/userrepo"
	// homepageservice "umkm/service/homepage"
	// brandlogoservice "umkm/service/homepage/brandlogo"
	kategoriprodukservice "umkm/service/kategori_produk"
	kategoriumkmservice "umkm/service/kategori_umkm"
	// produkservice "umkm/service/produk"
	// transaksiservice "umkm/service/transaksi"
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
	userUmkmService := umkmservice.NewUmkmService(userUmkmRepo, userHakAksesRepo, db)
	userUmkmController := umkmcontroller.NewUmkmController(userUmkmService)

	//  
	// userQuerBuilder := querybuilder.NewBaseQueryBuilder(db)
	// userTransaksiRepo := transaksirepo.NewTransaksiRepositoryImpl(db)
	// userTransaksiService := transaksiservice.NewTransaksiservice(userTransaksiRepo, db)
	// userTransaksiController := transaksicontroller.NewUmkmController(userTransaksiService, db)

	//userkategori produk
	userKategoriProdukRepo := kategoriprodukrepo.NewKategoriProdukRepo(db)
	userKategoriProdukService := kategoriprodukservice.NewKategoriProdukService(userKategoriProdukRepo)
	userKategoriProdukController := kategoriprodukcontroller.NewKategeoriProdukController(*userKategoriProdukService)

	//userproduk
	// userProdukrepo := produkrepo.NewProdukRepositoryImpl(db)
	// userProdukService := produkservice.NewProdukService(userProdukrepo)
	// userProdukController := produkcontroller.NewProdukController(userProdukService)

	//testimonial
	// userTestimonial := testimonialrepo.NewTestimonal(db)
	// userTesimonialService := homepageservice.NewTestimonialService(userTestimonial)
	// userTesimonialController := homepagecontroller.NewTestimonialController(*userTesimonialService)

	//logo
	// userLogo := brandrepo.NewBrandlogo(db)
	// userLogoBrandService := brandlogoservice.NewBrandLogoService(userLogo)
	// userLogoBrandController := brandlogo.NewBrandLogoController(userLogoBrandService)
	
	g := e.Group(prefix)

	authRoute := g.Group("/auth")
	authRoute.POST("/register", userAuthController.Register)
	authRoute.POST("/login", userAuthController.Login)
	authRoute.POST("/send-otp", userAuthController.SendOtp)
	authRoute.POST("/verifyOtp", userAuthController.VerifyOTPHandler)
	authRoute.POST("/sendotp-register", userAuthController.SendOtpRegister)
	authRoute.POST("/verifyOtpRegister", userAuthController.VerifyOTPHandlerRegister)
	// authRoute.POST("/logout", userAuthController.Logout, JWTProtection())

	meRoute := g.Group("/me")
	meRoute.GET("", userAuthController.View, JWTProtection())

	KatUmkmRoute := g.Group("/kategori")
	KatUmkmRoute.POST("/umkm", userKategoriUmkmController.Create, JWTProtection())
	KatUmkmRoute.GET("/umkm/list", userKategoriUmkmController.GetKategoriList, JWTProtection())
	KatUmkmRoute.GET("/umkm/:id", userKategoriUmkmController.GetKategoriId, JWTProtection())
	KatUmkmRoute.PUT("/umkm/:id", userKategoriUmkmController.UpdateKategoriId, JWTProtection())
	KatUmkmRoute.DELETE("/umkm/delete/:id", userKategoriUmkmController.DeleteKategoriId, JWTProtection())

	Umkm := g.Group("/umkm")
	Umkm.Static("/uploads", "uploads")

	Umkm.POST("/create", userUmkmController.Create, JWTProtection())
	Umkm.GET("/list", userUmkmController.GetUmkmList,JWTProtection())
	Umkm.GET("/filter", userUmkmController.GetUmkmFilter,JWTProtection())

	// Transaksi := g.Group("/transaksi")
	// Transaksi.POST("/umkm", userTransaksiController.Create)
	// Transaksi.GET("/:id", userTransaksiController.GetKategoriId)
	// Transaksi.GET("/:umkm_id/:date", userTransaksiController.GetTransaksiFilterList)

	KatProdukRoute := g.Group("/kategoriproduk")
	KatProdukRoute.POST("/create	", userKategoriProdukController.Create)
	KatProdukRoute.GET("/:umkm_id", userKategoriProdukController.GetKategoriList)

	//produk
	// Produk := g.Group("/produk")
	// Produk.POST("/create", userProdukController.CreateProduk)
	// Produk.DELETE("/delete/:id", userProdukController.DeleteProdukId)
	// Produk.GET("/list/:umkm_id", userProdukController.GetprodukList)

	//testimonial
	// Testimonial := g.Group("/testimonial")
	// Testimonial.POST("/create", userTesimonialController.Create)
	// Testimonial.GET("/list", userTesimonialController.GetTestimonial)
	// Testimonial.DELETE("/delete/:id", userTesimonialController.DeleteTestimonial)
	// Testimonial.GET("/:id", userTesimonialController.GetTestimonialId)
	// Testimonial.PUT("/update/:id", userTesimonialController.UpdateTestimonial)
	// Testimonial.GET("/list/active", userTesimonialController.GetTestimonialActive)
	// Testimonial.PUT("/update/active/:id", userTesimonialController.UpdateTestimonialActive)

	//brandlogo
	// Brandlogo := g.Group("/brandlogo")
	// Brandlogo.POST("/create", userLogoBrandController.Create)
	// Brandlogo.GET("/list", userLogoBrandController.GetBrandLogoList)
	// Brandlogo.DELETE("/delet/:id", userLogoBrandController.DeleteProdukId)
}

	func JWTProtection() echo.MiddlewareFunc {
		return echojwt.WithConfig(echojwt.Config{
			NewClaimsFunc: func(c echo.Context) jwt.Claims {
				return new(helper.JwtCustomClaims)
			},
			SigningKey: []byte(os.Getenv("SECRET_KEY")),
			ErrorHandler: func(c echo.Context, err error) error {
				return c.JSON(http.StatusUnauthorized, model.ResponseToClient(http.StatusUnauthorized, false, "unauthorized", nil))
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
