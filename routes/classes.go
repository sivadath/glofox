package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sivadath/glofox/controllers"
	"github.com/sivadath/glofox/storage"
)

func RegisterClassRoutes(r *gin.Engine) {
	cc := controllers.NewClassController(storage.DB)
	classRoutes := r.Group(Version + "/classes")

	classRoutes.POST("/", cc.CreateClass)
	classRoutes.GET("/", cc.GetClasses)
}
