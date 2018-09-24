package ctrl

import (
	"kk-backend/backend-project/db"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// UserCtrl :
type UserCtrl struct{}

var DB *gorm.DB

// IDS :
type IDS struct {
	ID []int `json:"id"`
}

// UserRegister :
type UserRegister struct {
	gorm.Model
	Email string `json:"car" gorm:"type:varchar(30)"`
	Fname string `json:"fname" gorm:"type:varchar(20)"`
	Lname string `json:"lname" gorm:"type:varchar(20)"`
}

// Read :
func (x UserCtrl) Read(c *gin.Context) {
	var data []UserRegister
	db.GetDB().Find(&data)
	c.JSON(200, data)
}

// Insert :
func (x UserCtrl) Insert(c *gin.Context) {
	var item UserRegister
	c.ShouldBind(&item)
	db.GetDB().Create(&item)
	c.String(200, "Success")
}

// Update :
func (x UserCtrl) Update(c *gin.Context) {
	data := UserRegister{}
	c.ShouldBind(&data)
	db.GetDB().Save(&data)
	c.String(200, "Update Success")
}

// Delete :
func (x UserCtrl) Delete(c *gin.Context) {
	data := UserRegister{}
	id := c.Query("id")
	i, _ := strconv.Atoi(id)
	data.ID = uint(i)
	db.GetDB().Delete(&data)
	c.String(200, "Success")
}

// DeleteSelect :
func (x UserCtrl) DeleteSelect(c *gin.Context) {
	var ids []int
	data := UserRegister{}
	c.ShouldBind(&ids)
	for i := 0; i < len(ids); i++ {
		data.ID = uint(ids[i])
		db.GetDB().Delete(&data)
	}
	c.String(200, "Success")
}
