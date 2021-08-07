// render.go
// 핸들러 로직 수행 결과를 다양한 형태(JSON, XML, Html Template 등)로 응답한다.

package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func render(c *gin.Context, data gin.H, templateName string) {
	//loggedInInterface, _ := c.Get("is_logged_in")
	//data["is_logged_in"] = loggedInInterface.(bool)

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}
}