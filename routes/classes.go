package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sivadath/glofox/controllers"
	"github.com/sivadath/glofox/storage"
)

const EndPointClasses = "/classes"

func RegisterClassRoutes(r *gin.Engine, db storage.Storage) {
	cc := controllers.NewClassController(db)
	classRoutes := r.Group(Version + EndPointClasses)

	classRoutes.POST("", cc.CreateClass)
	classRoutes.GET("", cc.GetClasses)
}
