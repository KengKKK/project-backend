package main

import (
	"kk-backend/backend-project/route"

	"github.com/9lon/gonylon"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// DB : 👍👍👍
var DB *gorm.DB

func main() {

	db.InitDB()

	r := gin.Default()
	r.Use(gonylon.CorsMiddleware())

	// All Route
	route.RegisterRoute(r.Group("/api/register"))
	route.UserRoute(r.Group("/api/user"))

	r.Run(":3000")
}
