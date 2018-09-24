package main

import (
	"fmt"
	"kk-backend/backend-lab/config"
	"kk-backend/backend-project/ctrl"
	"kk-backend/backend-project/route"

	"github.com/9lon/gonylon"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// DB :
var DB *gorm.DB

func main() {

	config := config.GetConfig()
	full := config.User + ":" + config.Pass + "@tcp(" + config.Host + ")/" + config.DBname + "?charset=utf8&parseTime=true"
	db, err := gorm.Open("mysql", full)
	DB = db
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(ctrl.ModelRegsiter())
	defer db.Close()

	db.InitDB()

	r := gin.Default()
	r.Use(gonylon.CorsMiddleware())

	// All Route
	route.RegisterRoute(r.Group("/api/register"))
	// route.UserRoute(r.Group("/api/user"))

	r.Run(":3000")
}
