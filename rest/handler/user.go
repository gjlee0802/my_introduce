package handler

import (
	"github.com/gjlee0802/my_introduce/domain/model"
	"net/http"
	"os"
	"time"

	//"github.com/gin-gonic/contrib/sessions"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (h *GinHandler) Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	var (
		u	*model.User
		err error
	)

	if u, err = h.ruc.MatchUser(username, password); err != nil {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}
	token, err := CreateToken(u.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, token)
}

func (h *GinHandler) Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	u, err := h.ruc.RegisterUser(username, password)
	if  err != nil{
		c.JSON(http.StatusBadRequest, "Registration Failed")
		return
	}
	token, err := CreateToken(u.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, token)
}

func CreateToken(userID int) (string, error) {
	var err error
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd")
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userID
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix() // 15분 유지됨.
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

