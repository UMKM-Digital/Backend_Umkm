package route

import (
	"net/http"
	"os"

	"umkm/app"
	beritacontroller "umkm/controller/berita"
	daerahcontoller "umkm/controller/daerah"
	datacontroller "umkm/controller/data"
	dokumenumkmcontroller "umkm/controller/dokumenumkm"
	hakaksescontroller "umkm/controller/hakakses"
	homepagecontroller "umkm/controller/homepage"
	aboutuscontroller "umkm/controller/homepage/aboutus"
	brandlogo "umkm/controller/homepage/logo"
	slidercontroller "umkm/controller/homepage/slider"
	omsetcontroller "umkm/controller/omset"
	sektorusahacontroller "umkm/controller/sektorusaha"
	storagegambarcontroller "umkm/controller/storagegambar"

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

	daerahrepo "umkm/repository/daerah"
	datarepo "umkm/repository/data"
	dokumenumkmrepo "umkm/repository/dokumenumkm"
	hakaksesrepo "umkm/repository/hakakses"
	testimonialrepo "umkm/repository/homepage"
	aboutusrepo "umkm/repository/homepage/aboutus"
	beritarepo "umkm/repository/homepage/berita"
	brandrepo "umkm/repository/homepage/brandlogo"
	sliderrepo "umkm/repository/homepage/slider"
	omsetrepo "umkm/repository/omset"
	sektorusaharepo "umkm/repository/sektorusaha"
	storagegambarrepo "umkm/repository/storagegambar"

	kategoriprodukrepo "umkm/repository/kategori_produk"
	repokategoriumkm "umkm/repository/kategori_umkm"
	masterdokumenlegalrepo "umkm/repository/masterdokumenlegal"
	produkrepo "umkm/repository/produk"
	transaksirepo "umkm/repository/transaksi"
	umkmrepo "umkm/repository/umkm"
	"umkm/repository/userrepo"
	daerahservice "umkm/service/daerah"
	dataservice "umkm/service/data"
	dokumenumkmservice "umkm/service/dokumenumkm"
	hakaksesservice "umkm/service/hak_akses"
	homepageservice "umkm/service/homepage"
	aboutusservice "umkm/service/homepage/aboutus"
	beritaservice "umkm/service/homepage/berita"
	brandlogoservice "umkm/service/homepage/brandlogo"
	sliderservice "umkm/service/homepage/slider"
	omsetservice "umkm/service/omset"
	sektorusahaservice "umkm/service/sektorusaha"
	storagegambarservice "umkm/service/storagegambarproduk"

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


	userHakAksesRepo := hakaksesrepo.NewHakAksesRepositoryImpl(db)
	userHakAksesService := hakaksesservice.NewKHakAkesService(userHakAksesRepo)
	userHakAksesController := hakaksescontroller.NewHakAksesController(userHakAksesService)


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
	umkmQueryBuilderlo := query_builder_umkm.NewUmkmQueryBuilder(db)
	userUmkmRepolo := umkmrepo.NewUmkmRepositoryImpl(db, umkmQueryBuilderlo)
	userHakAksesRepopro := hakaksesrepo.NewHakAksesRepositoryImpl(db)  
	userProdukService := produkservice.NewProdukService(userProdukrepo, userHakAksesRepopro, userUmkmRepolo)
	userProdukController := produkcontroller.NewProdukController(userProdukService)

	userDokumenuMKM := dokumenumkmrepo.NewDokumenRepositoryImpl(db)
	userDokumenUmkmService := dokumenumkmservice.NewDokumenUmkmService(userDokumenuMKM)
	userDokumenUmkmController := dokumenumkmcontroller.NewDokumenUmkmController(userDokumenUmkmService)
	transaksiQueryBuilder := general_query_builder.NewTransaksiQueryBuilder(db)

	userTransaksiRepo := transaksirepo.NewTransaksiRepositoryImpl(db, transaksiQueryBuilder)
	userTransaksiService := transaksiservice.NewTransaksiservice(userTransaksiRepo,db, userHakAksesRepo)
	userTransaksiController := transaksicontroller.NewTransaksiController(userTransaksiService, db)
	
	KategoriProdukQueryBuilder := query_builder_kategori_produk.NewKategoriProdukQueryBuilder(db)
	userKategoriProdukRepo := kategoriprodukrepo.NewKategoriProdukRepo(db,KategoriProdukQueryBuilder)
	userKategoriProdukService := kategoriprodukservice.NewKategoriProdukService(userKategoriProdukRepo)
	userKategoriProdukController := kategoriprodukcontroller.NewKategeoriProdukController(*userKategoriProdukService)

	useromsetRepo	:= omsetrepo.NewomsetRepositoryImpl(db)
	


	//umkm
	umkmQueryBuilder := query_builder_umkm.NewUmkmQueryBuilder(db)
	userUmkmRepo := umkmrepo.NewUmkmRepositoryImpl(db, umkmQueryBuilder)
	// userHakAksesRepo := hakaksesrepo.NewHakAksesRepositoryImpl(db)  
	userUmkmService := umkmservice.NewUmkmService(userUmkmRepo, userHakAksesRepo, db, userProdukrepo, userTransaksiRepo, userDokumenuMKM, userMasterLegal, useromsetRepo)
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

	//daerah
	userDaerah := daerahrepo.NewDaerah(db)
	userDaerahService := daerahservice.NewDaerahService(userDaerah)
	userDaerahController := daerahcontoller.NewDaerahController(userDaerahService)

dataUmkm := datarepo.NewDataRepositoryImpl(db)
userDataService := dataservice.NewDataservice(dataUmkm) // Ensure this constructor matches the signature
userDataController := datacontroller.NewUmkmController(userDataService)


//omset
omsetUmkm := omsetrepo.NewomsetRepositoryImpl(db)
userOmsetService := omsetservice.NewOmsetService(omsetUmkm, userHakAksesRepo)
userOmsetController := omsetcontroller.NewOmsetController(userOmsetService)

	userAuthRepo := userrepo.NewAuthRepositoryImpl(db)
	userAuthService := userservice.Newauthservice(userAuthRepo, tokenUseCase, db, userHakAksesRepo,userUmkmRepo,userProdukrepo,userDokumenuMKM,userTransaksiRepo, useromsetRepo)
	userAuthController := usercontroller.NewAuthController(userAuthService, tokenUseCase)

	repo := storagegambarrepo.NewStorageGambarRepo()
	service := storagegambarservice.NewStorageGambarService(repo)
	controller := storagegambarcontroller.NewStorageGambarController(service)

	g := e.Group(prefix)

	linkRoute := g.Group("/link")
	linkRoute.POST("/produk", func(c echo.Context) error {
		r := c.Request()
		w := c.Response()
		controller.UploadFiles(w, r)
		return nil
	})

	authRoute := g.Group("/auth")
	authRoute.POST("/register", userAuthController.Register)
	authRoute.POST("/login", userAuthController.Login)
	authRoute.POST("/send-otp", userAuthController.SendOtp)
	authRoute.POST("/verifyOtp", userAuthController.VerifyOTPHandler)
	authRoute.POST("/sendotp-register", userAuthController.SendOtpRegister)
	authRoute.POST("/verifyOtpRegister", userAuthController.VerifyOTPHandlerRegister)
	authRoute.POST("/cekpassword", userAuthController.CekPassword, JWTProtection())
	authRoute.POST("/updatepassword", userAuthController.ChangePassword, JWTProtection())
	authRoute.POST("/login-google", userAuthController.HandleGoogleLoginOrRegister) 
	authRoute.POST("/send_email", userAuthController.HandlePasswordResetRequest) 
	authRoute.POST("/send_email", userAuthController.HandlePasswordResetRequest) 
	authRoute.PUT("/edit_profile", userAuthController.Update, JWTProtection()) 
	authRoute.GET("/list", userAuthController.GetUser) 
	authRoute.GET("/count", userAuthController.GetUserCountByGender) 
	authRoute.DELETE("/delete", userAuthController.DeleteUser, JWTProtection())

	meRoute := g.Group("/me")
	meRoute.GET("", userAuthController.View, JWTProtection())

	KatUmkmRoute := g.Group("/kategori")
	KatUmkmRoute.POST("/create", userKategoriUmkmController.Create, JWTProtection())
	KatUmkmRoute.GET("/list", userKategoriUmkmController.GetKategoriList)
	KatUmkmRoute.GET("/:id", userKategoriUmkmController.GetKategoriId, JWTProtection())
	KatUmkmRoute.PUT("/update/:id", userKategoriUmkmController.UpdateKategoriId, JWTProtection())
	KatUmkmRoute.DELETE("/delete/:id", userKategoriUmkmController.DeleteKategoriId, JWTProtection())
	KatUmkmRoute.GET("/sektor/:id", userKategoriUmkmController.GetSektorUsaha)

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
	Umkm.GET("/list/activeback", userUmkmController.ListActveBack)
	Umkm.PUT("/updateactive/:id", userUmkmController.UpdateSldierActive)
	Umkm.GET("/active/list", userUmkmController.GetUmkmActive)

	//transaksi
	Transaksi := g.Group("/transaksi")
	Transaksi.POST("/umkm", userTransaksiController.Create, JWTProtection())
	Transaksi.GET("/:id", userTransaksiController.GetKategoriId, JWTProtection())
	Transaksi.GET("/:umkm_id/:date", userTransaksiController.GetTransaksiFilterList, JWTProtection())
	Transaksi.GET("/rekap-transaksi-tahunan", userTransaksiController.GetTransaksiByYear, JWTProtection())
	Transaksi.GET("/rekap-transaksi-bulanan", userTransaksiController.GetTransaksiByMonth, JWTProtection())
	Transaksi.GET("/rekap-transaksi-harian", userTransaksiController.GetTransaksiByDate, JWTProtection())

	//kategoriproduk
	KatProdukRoute := g.Group("/kategoriproduk")
	KatProdukRoute.POST("/create", userKategoriProdukController.Create)
	KatProdukRoute.GET("/list", userKategoriProdukController.GetKategoriList)
	KatProdukRoute.GET("/:id", userKategoriProdukController.GetKategoriId)
	KatProdukRoute.PUT("/update/:id", userKategoriProdukController.UpdateKategoriProduk)
	KatProdukRoute.DELETE("/delete/:id", userKategoriProdukController.Delete)

	//produk
	Produk := g.Group("/produk")
	Produk.POST("/create", userProdukController.CreateProduk, JWTProtection())
	Produk.DELETE("/delete/:id", userProdukController.DeleteProdukId, JWTProtection())
	//untuk mobile
	Produk.GET("/list/all/:umkm_id", userProdukController.GetprodukList, JWTProtection())
	//
	Produk.GET("/:id", userProdukController.GetProdukId, JWTProtection())
	Produk.PUT("/update/:id", userProdukController.UpdateProduk, JWTProtection())
	Produk.GET("/list", userProdukController.GetProdukListWeb)
	Produk.GET("/list/:id", userProdukController.GetProdukWebId)
	Produk.GET("/list/login", userProdukController.GetProdukByLogin, JWTProtection())
	Produk.GET("/umkm/baru/:id", userProdukController.GetProdukBaru)
	Produk.GET("/list/activeback", userProdukController.GetTopProuduk)
	Produk.PUT("/updateactive/:id", userProdukController.UpdateTopProduk)
	Produk.GET("/active/list", userProdukController.GetProdukActive)


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
	
	//
	bentukusaha := g.Group("/bentukusaha")
	bentukusaha.GET("/list", userSektorUsahaController.GetBentukUsaha)

	//statustempatusaha
	statustempatusaha := g.Group("/statustempatusaha")
	statustempatusaha.GET("/list", userSektorUsahaController.GetStatusTempatUsaha)
	
	//daerah
	daerah := g.Group("/daerah")
	daerah.GET("/provinsi", userDaerahController.GetDaerah)
	daerah.GET("/kabupaten/:id", userDaerahController.GetKabupaten)
	daerah.GET("/kecamatan/:id", userDaerahController.GetKecamatan)
	daerah.GET("/kelurahan/:id", userDaerahController.GetKelurahan)

	data := g.Group("/data")
	data.GET("/list", userDataController.CountData)
	data.GET("/grafik", userDataController.GrafikKategoriBySektorHandler)
	data.GET("/grafikbinaan", userDataController.TotalUmkmKriteriaUsahaPerBulanHandler)
	data.GET("/umkmlist", userDataController.CountUmkmBulan)
	data.GET("/omset", userDataController.CountOmzets)
	data.GET("/grafikbinaantahun",userDataController.TotalUmkmKriteriaUsahaPertahun)
	data.GET("/umkm", userDataController.CountPengggunaUmkm, JWTProtection())
	data.GET("/omzet_bulan", userDataController.CountPenggunaOmzet, JWTProtection())

	hakakses := g.Group("/hakakses")
	hakakses.PUT("/update", userHakAksesController.UpdateHakAksesIds)

	omset := g.Group("/omset")
	omset.POST("/create/:umkm_id", userOmsetController.CreateOmsetcontroller, JWTProtection())
	omset.GET("/list/:umkm_id", userOmsetController.LisOmsetController, JWTProtection())
	omset.GET("/:id", userOmsetController.GetOmsetController, JWTProtection())
	omset.PUT("/update/:id", userOmsetController.UpdateOmset, JWTProtection())
	omset.GET("/grafik/list/:umkm_id", userOmsetController.ListOmsetGrafik, JWTProtection())
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
	

