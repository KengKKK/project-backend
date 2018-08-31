package route

import (
	"kk-backend/backend-project/ctrl"

	"github.com/gin-gonic/gin"
)

// RegisterRoute :
func RegisterRoute(r *gin.RouterGroup) {
	registerCtrl := ctrl.RegisterCtrl{}
	r.GET("/", registerCtrl.Read)
	r.POST("/", registerCtrl.Insert)
	r.PUT("/", registerCtrl.Update)
	r.DELETE("/", registerCtrl.Delete)
	// r.DELETE("/all", registerCtrl.DeleteSelect)
}
