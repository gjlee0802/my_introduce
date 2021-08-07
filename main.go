package main

import (
	"github.com/gin-gonic/gin"
)

func main(){

	// mysql DB


	// Router
	r := gin.Default()
	r.LoadHTMLGlob("view/*")
	//r.GET("/", Need Handler Func here) // should call handler.showIndexPage

	r.Run()
}