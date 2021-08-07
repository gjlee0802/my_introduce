package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/gjlee0802/my_introduce/domain/model"
)

const (
	dbms = "mysql"
	user = "root"
	pass = "password"
	db = "my_introduce"
)

var dbConn *gorm.DB

func Setup() {
	var err error
	conn := fmt.Sprintf("%s:%s@/%s?parseTime=true", user, pass, db)
	dbConn, err = gorm.Open(dbms, conn)
	if err != nil{
		panic(err)
	}

	dbConn.AutoMigrate(&model.User{},)
}