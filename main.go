package main

import (
	"kk-backend/backend-project/db"
	"kk-backend/backend-project/route"
	"strconv"

	"github.com/9lon/gonylon"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// RegisterPeople :
type RegisterPeople struct {
	gorm.Model
	Career     string `json:"car" gorm:"type:varchar(20)"`
	Fname      string `json:"fname" gorm:"type:varchar(30)"`
	Lname      string `json:"lname" gorm:"type:varchar(30)"`
	Age        int    `json:"age"`
	Faculty    string `json:"fac" gorm:"type:varchar(30)"`
	Department string `json:"dep" gorm:"type:varchar(30)"`
}

// DeviceModel :
// type DeviceModel struct {
// 	gorm.Model
// 	Project string `json:"pname" gorm:"type:varchar(100)"`
// 	Type    string `json:"spec" gorm:"type:varchar(100)"`
// 	Surname string `json:"sname" gorm:"type:varchar(100)"`
// 	Token   string `json:"key" gorm:"type:varchar(100)"`
// 	State   int    `json:"state"  gorm:"type:int(10)"`
// }

// IDS :
type IDS struct {
	ID []int `json:"id"`
}

// DB :
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

func read(c *gin.Context) {
	// var data []RegisterPeople
	// var one RegisterPeople
	// id := c.Query("id")
	// i, _ := strconv.Atoi(id)

	// if i == 0 {
	// 	DB.Find(&data)
	// 	c.JSON(200, data)
	// } else {
	// 	DB.First(&one, i)
	// 	c.JSON(200, one)
	// }
}
func insert(c *gin.Context) {
	// var item RegisterPeople
	// fmt.Println(item)
	// c.ShouldBind(&item)
	// DB.Create(&item)
	// c.String(200, "Success")
}

func update(c *gin.Context) {
	data := RegisterPeople{}
	c.ShouldBind(&data)
	DB.Save(&data)
	c.String(200, "Update Success")
}

func delete(c *gin.Context) {

	data := RegisterPeople{}

	id := c.Query("id")
	i, _ := strconv.Atoi(id)
	data.ID = uint(i)
	DB.Delete(&data)
	c.String(200, "Success")

}

// func delete2(c *gin.Context) {

// 	data2 := RegisterPeople{}

// 	id := c.Query("id")
// 	i, _ := strconv.Atoi(id)
// 	data2.ID = uint(i)
// 	DB.Delete(&data2)
// 	c.String(200, "Success")

// }

func deleteall(c *gin.Context) {
	var ids []int
	data := RegisterPeople{}
	c.ShouldBind(&ids)
	for i := 0; i < len(ids); i++ {
		data.ID = uint(ids[i])
		DB.Delete(&data)
	}
	c.String(200, "Success")
}
