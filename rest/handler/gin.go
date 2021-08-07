package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	//"github.com/gin-gonic/gin"
	"github.com/gjlee0802/my_introduce/usecase"
)

func catchPanic() {
	if p := recover(); p != nil {
		log.Printf("%+v\n", p)
	}
}

type GinHandler struct {
	ruc usecase.RegistrationUsecase
}

func NewGinHandler(ruc usecase.RegistrationUsecase) *GinHandler {
	return &GinHandler{
		ruc:  ruc,
	}
}
func (h *GinHandler) ShowIndexPage(c *gin.Context) {

	// Current articles include 'Writer.Password'.
	// Since this is information to be hidden in the client, you need to use an adapter,
	// but it is an example code, so I will skip it.

	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title":   "Home Page",
	}, "index.html")
}

/*
func (h *GinHandler) Show~~Page(c *gin.Context) {
	render(c, gin.H{
	}, "~~~.html")
}
*/