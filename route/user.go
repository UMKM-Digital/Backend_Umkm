package route

import (
	"net/http"
	"os"

	"umkm/app"
	beritacontroller "umkm/controller/berita"
	dokumenumkmcontroller "umkm/controller/dokumenumkm"
	homepagecontroller "umkm/controller/homepage"
	aboutuscontroller "umkm/controller/homepage/aboutus"
	brandlogo "umkm/controller/homepage/logo"
	slidercontroller "umkm/controller/homepage/slider"
	sektorusahacontroller "umkm/controller/sektorusaha"

	kategoriprodukcontroller "umkm/controller/kategoriproduk"
	kategoriumkmcontroller "umkm/controller/kategoriumkm"
	masterlegalcontroller "umkm/controller/masterlegal"
	produkcontroller "umkm/controller/produk"
	transaksicontroller "umkm/controller/transaksi"
	umkmcontroller "umkm/controller/umkm"
	"umkm/controller/usercontroller"
	"umkm/helper"
	"umkm/model"

	query_builder_berita "umkm/query_builder/berita"
	query_builder_kategori_produk "umkm/query_builder/kategoriproduk"
	query_builder_kategori_umkm "umkm/query_builder/kategoriumkm"
	query_builder_masterlegal "umkm/query_builder/masterlegal"
	query_builder_produk "umkm/query_builder/produk"
	general_query_builder "umkm/query_builder/transaksi"
	query_builder_umkm "umkm/query_builder/umkm"

	dokumenumkmrepo "umkm/repository/dokumenumkm"
	hakaksesrepo "umkm/repository/hakakses"
	testimonialrepo "umkm/repository/homepage"
	aboutusrepo "umkm/repository/homepage/aboutus"
	beritarepo "umkm/repository/homepage/berita"
	brandrepo "umkm/repository/homepage/brandlogo"
	sliderrepo "umkm/repository/homepage/slider"
	sektorusaharepo "umkm/repository/sektorusaha"

	kategoriprodukrepo "umkm/repository/kategori_produk"
	repokategoriumkm "umkm/repository/kategori_umkm"
	masterdokumenlegalrepo "umkm/repository/masterdokumenlegal"
	produkrepo "umkm/repository/produk"
	transaksirepo "umkm/repository/transaksi"
	umkmrepo "umkm/repository/umkm"
	"umkm/repository/userrepo"
	dokumenumkmservice "umkm/service/dokumenumkm"
	homepageservice "umkm/service/homepage"
	aboutusservice "umkm/service/homepage/aboutus"
	beritaservice "umkm/service/homepage/berita"
	brandlogoservice "umkm/service/homepage/brandlogo"
	sliderservice "umkm/service/homepage/slider"
	sektorusahaservice "umkm/service/sektorusaha"

	kategoriprodukservice "umkm/service/kategori_produk"
	kategoriumkmservice "umkm/service/kategori_umkm"
	masterdokumenlegalservice "umkm/service/masterdokumenlegal"
	produkservice "umkm/service/produk"
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

	kategoriUmkmQueryBuilder := query_builder_kategori_umkm.NewKategoriUmkmQueryBuilder(db)
	userKategoriUmkmRepo := repokategoriumkm.NewKategoriUmkmRepositoryImpl(db, kategoriUmkmQueryBuilder)
	userKatgoriUmkmService := kategoriumkmservice.NewKategoriUmkmService(userKategoriUmkmRepo)
	userKategoriUmkmController := kategoriumkmcontroller.NewKategeoriUmkmController(userKatgoriUmkmService)

	masterlegalQueryBuilder := query_builder_masterlegal.NewMasteLegalQueryBuilder(db)
	userMasterLegal := masterdokumenlegalrepo.NewDokumenLegalRepoImpl(db, masterlegalQueryBuilder)
	userMasterLegalService := masterdokumenlegalservice.NewMasterLegalService(userMasterLegal)
	userMasterLegalController := masterlegalcontroller.NewKategeoriProdukController(userMasterLegalService)


	//userproduk
	produkQueryBuilder := query_builder_produk.NewProdukQueryBuilder(db)
	userProdukrepo := produkrepo.NewProdukRepositoryImpl(db, produkQueryBuilder)
	userProdukService := produkservice.NewProdukService(userProdukrepo)
	userProdukController := produkcontroller.NewProdukController(userProdukService)

	userDokumenuMKM := dokumenumkmrepo.NewDokumenRepositoryImpl(db)
	userDokumenUmkmService := dokumenumkmservice.NewDokumenUmkmService(userDokumenuMKM)
	userDokumenUmkmController := dokumenumkmcontroller.NewDokumenUmkmController(userDokumenUmkmService)
	transaksiQueryBuilder := general_query_builder.NewTransaksiQueryBuilder(db)

	userTransaksiRepo := transaksirepo.NewTransaksiRepositoryImpl(db, transaksiQueryBuilder)
	userTransaksiService := transaksiservice.NewTransaksiservice(userTransaksiRepo, db)
	userTransaksiController := transaksicontroller.NewTransaksiController(userTransaksiService, db)
	
	KategoriProdukQueryBuilder := query_builder_kategori_produk.NewKategoriProdukQueryBuilder(db)
	userKategoriProdukRepo := kategoriprodukrepo.NewKategoriProdukRepo(db,KategoriProdukQueryBuilder)
	userKategoriProdukService := kategoriprodukservice.NewKategoriProdukService(userKategoriProdukRepo)
	userKategoriProdukController := kategoriprodukcontroller.NewKategeoriProdukController(*userKategoriProdukService)


	//umkm
	umkmQueryBuilder := query_builder_umkm.NewUmkmQueryBuilder(db)
	userUmkmRepo := umkmrepo.NewUmkmRepositoryImpl(db, umkmQueryBuilder)
	userHakAksesRepo := hakaksesrepo.NewHakAksesRepositoryImpl(db)  
	userUmkmService := umkmservice.NewUmkmService(userUmkmRepo, userHakAksesRepo, db, userProdukrepo, userTransaksiRepo, userDokumenuMKM, userMasterLegal)
	userUmkmController := umkmcontroller.NewUmkmController(userUmkmService)
 

	//testimonial
	userTestimonial := testimonialrepo.NewTestimonal(db)
	userTesimonialService := homepageservice.NewTestimonialService(userTestimonial)
	userTesimonialController := homepagecontroller.NewTestimonialController(*userTesimonialService)

	//logo
	userLogo := brandrepo.NewBrandlogo(db)
	userLogoBrandService := brandlogoservice.NewBrandLogoService(userLogo)
	userLogoBrandController := brandlogo.NewBrandLogoController(userLogoBrandService)

	//aboutus
	userAboutUs := aboutusrepo.NewAboutUS(db)
	userAboutUsService := aboutusservice.NewAboutUsService(userAboutUs)
	userAboutUsController := aboutuscontroller.NewAboutUsController(*userAboutUsService)
	
	//slider
	userSlider := sliderrepo.NewSlider(db)
	userSliderService := sliderservice.NewSliderService(userSlider)
	userSliderController := slidercontroller.NewTestimonialController(*userSliderService)
	

	//berita
	beritaQueryBuilder := query_builder_berita.NewBeritaQueryBuilder(db)
	userBerita := beritarepo.NewBerita(db, beritaQueryBuilder)
	userBeritaService := beritaservice.NewBeritaService(userBerita)
	userBeritaController := beritacontroller.NewBeritaController(*userBeritaService)

	//sektorusaha
	userSektorUsaha := sektorusaharepo.NewSektorUsaha(db)
	userSektorUsahaService := sektorusahaservice.NewSektorUsahaService(userSektorUsaha)
	userSektorUsahaController := sektorusahacontroller.NewSektorUsahaController(userSektorUsahaService)

	g := e.Group(prefix)

	authRoute := g.Group("/auth")
	authRoute.POST("/register", userAuthController.Register)
	authRoute.POST("/login", userAuthController.Login)
	authRoute.POST("/send-otp", userAuthController.SendOtp)
	authRoute.POST("/verifyOtp", userAuthController.VerifyOTPHandler)
	authRoute.POST("/sendotp-register", userAuthController.SendOtpRegister)
	authRoute.POST("/verifyOtpRegister", userAuthController.VerifyOTPHandlerRegister)
	// authRoute.POST("/logout", userAuthController.Logout, JWTProtection())
	authRoute.POST("/cekpassword", userAuthController.CekPassword, JWTProtection())
	authRoute.POST("/updatepassword", userAuthController.ChangePassword, JWTProtection())

	meRoute := g.Group("/me")
	meRoute.GET("", userAuthController.View, JWTProtection())

	KatUmkmRoute := g.Group("/kategori")
	KatUmkmRoute.POST("/create", userKategoriUmkmController.Create, JWTProtection())
	KatUmkmRoute.GET("/list", userKategoriUmkmController.GetKategoriList)
	KatUmkmRoute.GET("/:id", userKategoriUmkmController.GetKategoriId, JWTProtection())
	KatUmkmRoute.PUT("/update/:id", userKategoriUmkmController.UpdateKategoriId, JWTProtection())
	KatUmkmRoute.DELETE("/delete/:id", userKategoriUmkmController.DeleteKategoriId, JWTProtection())

	//umkm
	Umkm := g.Group("/umkm")
	Umkm.Static("/uploads", "uploads")

	Umkm.POST("/create", userUmkmController.Create, JWTProtection())
	Umkm.GET("/list", userUmkmController.GetUmkmList,JWTProtection())
	Umkm.GET("/filter", userUmkmController.GetUmkmFilter,JWTProtection())
	Umkm.GET("/web/list", userUmkmController.GetUmkmListWeb, JWTProtection())
	Umkm.GET("/:id", userUmkmController.GetUmkmId)
	Umkm.PUT("/edit/:umkm_id", userUmkmController.UpdateUmkm)
	Umkm.GET("/all/list", userUmkmController.GetUmmkmList)
	Umkm.GET("/detail/list/:id", userUmkmController.GetUmkmListDetial)
	Umkm.DELETE("/:id", userUmkmController.DeleteUmkmId)

	//transaksi
	Transaksi := g.Group("/transaksi")
	Transaksi.POST("/umkm", userTransaksiController.Create)
	Transaksi.GET("/:id", userTransaksiController.GetKategoriId)
	Transaksi.GET("/:umkm_id/:date", userTransaksiController.GetTransaksiFilterList)
	Transaksi.GET("/web/:umkm_id", userTransaksiController.GetTransaksiByYear)
	Transaksi.GET("/web/mounth/:umkm_id", userTransaksiController.GetTransaksiByMonth)
	Transaksi.GET("/web/date/:umkm_id", userTransaksiController.GetTransaksiByDate)

	//kategoriproduk
	KatProdukRoute := g.Group("/kategoriproduk")
	KatProdukRoute.POST("/create", userKategoriProdukController.Create)
	KatProdukRoute.GET("/list", userKategoriProdukController.GetKategoriList)
	KatProdukRoute.GET("/:id", userKategoriProdukController.GetKategoriId)
	KatProdukRoute.PUT("/update/:id", userKategoriProdukController.UpdateKategoriProduk)
	KatProdukRoute.DELETE("/delete/:id", userKategoriProdukController.Delete)

	//produk
	Produk := g.Group("/produk")
	Produk.POST("/create", userProdukController.CreateProduk)
	Produk.DELETE("/delete/:id", userProdukController.DeleteProdukId)
	//untuk mobile
	Produk.GET("/list/all/:umkm_id", userProdukController.GetprodukList)
	//
	Produk.GET("/:id", userProdukController.GetProdukId)
	Produk.PUT("/update/:id", userProdukController.UpdateProduk)
	Produk.GET("/list", userProdukController.GetProdukListWeb)
	Produk.GET("/list/:id", userProdukController.GetProdukWebId)


	//testimonial
	Testimonial := g.Group("/testimonial")
	Testimonial.POST("/create", userTesimonialController.Create)
	Testimonial.GET("/list", userTesimonialController.GetTestimonial)
	Testimonial.DELETE("/delete/:id", userTesimonialController.DeleteTestimonial)
	Testimonial.GET("/:id", userTesimonialController.GetTestimonialId)
	Testimonial.PUT("/update/:id", userTesimonialController.UpdateTestimonial)
	Testimonial.GET("/list/active", userTesimonialController.GetTestimonialActive)
	Testimonial.PUT("/update/active/:id", userTesimonialController.UpdateTestimonialActive)

	//brandlogo
	Brandlogo := g.Group("/brandlogo")
	Brandlogo.POST("/create", userLogoBrandController.Create)
	Brandlogo.GET("/list", userLogoBrandController.GetBrandLogoList)
	Brandlogo.DELETE("/delet/:id", userLogoBrandController.DeleteProdukId)
	Brandlogo.GET("/:id", userLogoBrandController.GetBrandLogoId)
	Brandlogo.PUT("/edit/:id", userLogoBrandController.UpdateBrandLogo)

	//aboutus
	Aboutus := g.Group("/aboutus")
	Aboutus.POST("/create", userAboutUsController.Create)
	Aboutus.GET("/:id", userAboutUsController.GetAboutId)
	Aboutus.GET("/list", userAboutUsController.GetAboutUs)
	Aboutus.PUT("/edit/:id", userAboutUsController.UpdateAboutUs)

	Slider := g.Group("/slider")
	Slider.POST("/create", userSliderController.Create)
	Slider.GET("/list", userSliderController.List)
	Slider.GET("/:id", userSliderController.GetSlideId)
	Slider.DELETE("/delete/:id", userSliderController.DelSlideId)
	Slider.PUT("/edit/:id",userSliderController.Update)
	Slider.PUT("/edit/active/:id",userSliderController.UpdateSldierActive)
	Slider.GET("/list/active", userSliderController.GetSlideralActive)

	masterlegal := g.Group("/masterlegal")
	masterlegal.POST("/create", userMasterLegalController.Create)
	masterlegal.GET("/list",userMasterLegalController.GetMasterLegalList)
	masterlegal.DELETE("/delete/:id",userMasterLegalController.Delete)
	masterlegal.GET("/:id", userMasterLegalController.GetIdMasterLegalId)
	masterlegal.PUT("/edit/:id", userMasterLegalController.UpdateMasterLegalId)
	masterlegal.GET("/list/dokumenumkm/:umkm_id", userMasterLegalController.List)
	masterlegal.GET("/dokumen-legal-umkm", userMasterLegalController.ListAll, JWTProtection())

	dokumenumkm := g.Group("/dokumenumkm")
	dokumenumkm.POST("/dokumen-legal-by-umkm/:dokumen_id/:umkm_id", userDokumenUmkmController.Create)
	dokumenumkm.GET("/:id/:umkm_id", userDokumenUmkmController.GetDokumenId)
	dokumenumkm.PUT("/edit/:id/:umkm_id", userDokumenUmkmController.UpdateProduk)

	berita := g.Group("/berita")
	berita.POST("/create", userBeritaController.Create, JWTProtection())
	berita.GET("/list", userBeritaController.LIst)
	berita.DELETE("/delete/:id", userBeritaController.Delete)
	berita.GET("/:id", userBeritaController.GetId)
	berita.PUT("/edit/:id",userBeritaController.Update)

	sektorusaha := g.Group("/sektorusaha")
	sektorusaha.POST("/create", userSektorUsahaController.Create)
	sektorusaha.GET("/list", userSektorUsahaController.GetSektorUsaha)
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
