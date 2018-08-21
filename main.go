package main

import (
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/9lon/gonylon"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"nextcorp.co.th/backend-lab/config"
)

// StudentModel :
type StudentModel struct {
	gorm.Model
	Fname string `json:"fname" gorm:"type:varchar(100)"`
	Lname string `json:"lname" gorm:"type:varchar(100)"`
	Age   int    `json:"age"`
}

// DeviceModel :
type DeviceModel struct {
	gorm.Model
	Project string `json:"pname" gorm:"type:varchar(100)"`
	Type    string `json:"spec" gorm:"type:varchar(100)"`
	Surname string `json:"sname" gorm:"type:varchar(100)"`
	Token   string `json:"key" gorm:"type:varchar(100)"`
	State   int    `json:"state"  gorm:"type:int(10)"`
}

// IDS :
type IDS struct {
	ID []int `json:"id"`
}

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
	db.AutoMigrate(&StudentModel{})
	db.AutoMigrate(&DeviceModel{})
	defer db.Close()

	r := gin.Default()
	r.Use(gonylon.CorsMiddleware())

	path := "/api/student"
	path2 := "/api/device"
	pathLogin := "/api/login"
	pathDevice := "/api/mqtt"

	r.GET(path, read)
	r.GET(path2, read2)
	r.GET(pathDevice, readDevice)

	r.POST(path, insert)
	r.POST(path2, insert2)
	r.GET(pathLogin, read3)
	r.PUT(path, update)
	r.PUT(path2, update2)

	r.PUT(path+"/all", deleteall)
	r.PUT(path2+"/all", deleteall2)

	r.DELETE(path, delete)
	r.DELETE(path2, delete2)
	r.Run(":3000") // listen and serve on 127.0.0.1:8080 default
}

func read(c *gin.Context) {
	var data []StudentModel
	var one StudentModel
	id := c.Query("id")
	i, _ := strconv.Atoi(id)

	if i == 0 {
		DB.Find(&data)
		c.JSON(200, data)
	} else {
		DB.First(&one, i)
		c.JSON(200, one)
	}
}

func read2(c *gin.Context) {
	var dataDevice []DeviceModel
	var oneDevice DeviceModel

	id := c.Query("id")
	i, _ := strconv.Atoi(id)

	if i == 0 {
		DB.Find(&dataDevice)
		c.JSON(200, dataDevice)
	} else {
		DB.First(&oneDevice, i)
		c.JSON(200, oneDevice)
	}
}

func read3(c *gin.Context) {

	var one StudentModel
	user := c.Query("user")
	pass := c.Query("password")

	var count int

	DB.Where("fname = ? AND lname = ?", user, pass).First(&one).Count(&count)

	if count == 1 {
		c.JSON(200, &one)
	} else {
		c.String(200, "false")
	}
}
func readDevice(c *gin.Context) {

	var one DeviceModel
	mqtt := c.Query("id")
	var count int

	DB.Where("id = ?", mqtt).First(&one).Count(&count)

	if count == 1 {
		c.JSON(200, &one)
	} else {
		c.String(200, "false")
	}
}

func insert(c *gin.Context) {
	var item StudentModel
	fmt.Println(item)
	c.ShouldBind(&item)
	DB.Create(&item)
	c.String(200, "Success")
}
func insert2(c *gin.Context) {
	var item2 DeviceModel
	fmt.Println(item2)
	c.ShouldBind(&item2)
	u2, _ := uuid.NewV4()
	item2.Token = hex.EncodeToString(u2.Bytes())
	fmt.Println(item2.Token)
	DB.Create(&item2)
	c.String(200, "Success")
}

func update(c *gin.Context) {
	data := StudentModel{}
	c.ShouldBind(&data)
	DB.Save(&data)
	c.String(200, "Update Success")
}
func update2(c *gin.Context) {
	data2 := DeviceModel{}
	c.ShouldBind(&data2)
	DB.Save(&data2)
	c.String(200, "Update Success")
}

func delete(c *gin.Context) {

	data := StudentModel{}

	id := c.Query("id")
	i, _ := strconv.Atoi(id)
	data.ID = uint(i)
	DB.Delete(&data)
	c.String(200, "Success")

}
func delete2(c *gin.Context) {

	data2 := DeviceModel{}

	id := c.Query("id")
	i, _ := strconv.Atoi(id)
	data2.ID = uint(i)
	DB.Delete(&data2)
	c.String(200, "Success")

}

func deleteall(c *gin.Context) {
	var ids []int
	data := StudentModel{}
	c.ShouldBind(&ids)
	for i := 0; i < len(ids); i++ {
		data.ID = uint(ids[i])
		DB.Delete(&data)
	}
	c.String(200, "Success")
}

func deleteall2(c *gin.Context) {
	var ids []int
	data := DeviceModel{}
	c.ShouldBind(&ids)
	for i := 0; i < len(ids); i++ {
		data.ID = uint(ids[i])
		DB.Delete(&data)
	}
	c.String(200, "Success")
}
