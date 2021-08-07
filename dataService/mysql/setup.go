package mysql

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/gjlee0802/my_introduce/domain/model"
	"github.com/gjlee0802/my_introduce/setting"
)

var dbConn *gorm.DB

func Setup() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", setting.DBsetting.User, setting.DBsetting.Pass, setting.DBsetting.Server, setting.DBsetting.Database)
	//conn := fmt.Sprintf("%s:%s@/%s?parseTime=true", user, pass, db)
	dbConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		panic(errors.New("DB connection failed."))
	}

	dbConn.AutoMigrate(&model.User{},)
}