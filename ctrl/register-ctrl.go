package ctrl

import (
	"kk-backend/backend-project/db"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// RegisterCtrl :
type RegisterCtrl struct{}

var DB *gorm.DB

// IDS :
type IDS struct {
	ID []int `json:"id"`
}

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

// Read :
func (x RegisterCtrl) Read(c *gin.Context) {
	var data []RegisterPeople
	db.GetDB().Find(&data)
	c.JSON(200, data)
}

// Insert :
func (x RegisterCtrl) Insert(c *gin.Context) {
	var item RegisterPeople
	c.ShouldBind(&item)
	db.GetDB().Create(&item)
	c.String(200, "Success")
}

// Update :
func (x RegisterCtrl) Update(c *gin.Context) {
	data := RegisterPeople{}
	c.ShouldBind(&data)
	db.GetDB().Save(&data)
	c.String(200, "Update Success")
}

// Delete :
func (x RegisterCtrl) Delete(c *gin.Context) {
	data := RegisterPeople{}
	id := c.Query("id")
	i, _ := strconv.Atoi(id)
	data.ID = uint(i)
	db.GetDB().Delete(&data)
	c.String(200, "Success")
}

// DeleteSelect :
func (x RegisterCtrl) DeleteSelect(c *gin.Context) {
	var ids []int
	data := RegisterPeople{}
	c.ShouldBind(&ids)
	for i := 0; i < len(ids); i++ {
		data.ID = uint(ids[i])
		db.GetDB().Delete(&data)
	}
	c.String(200, "Success")
}
