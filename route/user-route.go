package route

import (
	"kk-backend/backend-project/ctrl"

	"github.com/gin-gonic/gin"
)

// UserRoute :
func UserRoute(r *gin.RouterGroup) {

	userCtrl := ctrl.UserCtrl{}
	r.GET("/", userCtrl.Read)
	r.POST("/", userCtrl.Insert)
	r.PUT("/", userCtrl.Update)
	r.DELETE("/", userCtrl.Delete)
	// r.DELETE("/all", userCtrl.DeleteSelect)

}
