package app

import (
	"github.com/labstack/echo/v4/middleware"
	"go-svc-tpl/app/Controller"
	"go-svc-tpl/app/Controller/LinkControl"
	"go-svc-tpl/app/Controller/UserControl"
	"go-svc-tpl/model"
)

func addRoutes() {
	//e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/pong", Controller.Pong)
	//重定向
	e.GET("/:shortLink", LinkControl.Redirect)

	e.POST("/create", LinkControl.CreateLink)
	e.POST("/register", UserControl.Register)
	e.POST("/login", UserControl.Login)

	login := e.Group("/user")
	config := middleware.JWTConfig{
		Claims:      &model.JwtCustomClaims{},
		SigningKey:  []byte("secret"),
		TokenLookup: "header:token",
	}
	login.Use(middleware.JWTWithConfig(config))
	login.POST("/info", UserControl.Info)
	login.POST("/logout", UserControl.LogOut)
	login.POST("/info", UserControl.Info)
	login.POST("/url/get", UserControl.GetAllUrl)

	login.POST("/create", LinkControl.CreateLinkLogin)
	login.POST("/query", LinkControl.QueryLink)
	login.POST("/update", LinkControl.UpdateLink)
	login.POST("/delete", LinkControl.DeleteLink)
	login.POST("/pause", LinkControl.PauseLink)

}
