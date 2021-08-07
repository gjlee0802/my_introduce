package main

import (
	//"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"github.com/gjlee0802/my_introduce/dataService/mysql"
	"github.com/gjlee0802/my_introduce/usecase/registration"
	"github.com/gjlee0802/my_introduce/rest/handler"
	//"github.com/gjlee0802/my_introduce/rest/middleware"
)

func main(){

	// mysql DB
	mysql.Setup()

	// Repo
	ur := mysql.NewUserRepo()
	// Usecases
	ruc := registration.NewRegistrationUsecase(ur)
	// Handler
	h := handler.NewGinHandler(ruc)

	// Router
	r := gin.Default()
	r.Use(favicon.New("view/assets/img/favicon.ico"))
	//r.LoadHTMLGlob("view/*")
	r.LoadHTMLFiles("view/index.html")
	r.Static("/css", "./view/css")
	r.Static("/assets", "./view/assets")
	r.Static("/js", "./view/js")
	//r.GET("/", Need Handler Func here) // should call handler.showIndexPage

	//
	r.GET("/", h.ShowIndexPage)
 	/*
	user := r.Group("/u")
	{
		user.GET("/login", notLoggedIn, h.ShowLoginPage)
		user.POST("/login", notLoggedIn, h.Login)
		user.GET("/logout", loggedIn, h.Logout)
		user.GET("/register", notLoggedIn, h.ShowRegistrationPage)
		user.POST("/register", notLoggedIn, h.Register)
	}
 	*/

	r.Run(":8080")
}