package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sivadath/glofox/controllers"
	"github.com/sivadath/glofox/storage"
	"path"
)

const EndPointClasses = "classes"

func RegisterClassRoutes(r *gin.Engine, db storage.Storage) {
	cc := controllers.NewClassController(db)
	classRoutes := r.Group(path.Join(Version, EndPointClasses))

	classRoutes.POST("", cc.CreateClass)
	classRoutes.GET("", cc.GetClasses)
}
