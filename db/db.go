package db

import (
	"kk-backend/backend-project/config"

	"github.com/jinzhu/gorm"
)

// DB :
var DB *gorm.DB

// InitDB :
func InitDB() {
	config := config.GetConfig()
	con := config.User + ":" + config.Pass + "@tcp(" + config.Host + ")/" + config.DBname + "?charset=utf8&parseTime=true"
	db, _ := gorm.Open("mysql", con)
	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
