package route

import (
	"umkm/app"
	"umkm/controller/usercontroller"
	"umkm/helper"
	"umkm/repository/userrepo"
	userservice "umkm/service/user"

	"github.com/labstack/echo/v4"
)

func RegisterUserRoute(prefix string, e *echo.Echo) {
    db := app.DBConnection()
    token := helper.NewTokenUseCase()
    userAuthRepo := userrepo.NewAuthRepositoryImpl(db)
    userAuthService := userservice.Newauthservice(userAuthRepo, token)
    userAuthController := usercontroller.NewAuthController(userAuthService)

    g := e.Group(prefix)

    authRoute := g.Group("/auth")
    authRoute.POST("/register", userAuthController.Register)
    authRoute.POST("/login", userAuthController.Login)
}
