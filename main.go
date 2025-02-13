package main

import (
	"github.com/gin-gonic/gin"
	docs "github.com/sivadath/glofox/docs"
	"github.com/sivadath/glofox/routes"
	"github.com/sivadath/glofox/storage"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

// @title Glofox API
// @version 1.0
// @description API for managing classes and bookings in a fitness studio.
// @host localhost:8080
// @BasePath /
func main() {
	r := gin.Default()

	// Initialize storage with an in-memory implementation
	db := storage.NewInMemoryStorage()
	storage.SetStorage(db)

	// Swagger docs
	docs.SwaggerInfo.BasePath = routes.Version
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Register routes
	routes.RegisterClassRoutes(r, storage.DB)
	routes.RegisterBookingRoutes(r, storage.DB)

	log.Fatal(r.Run(":8080"))
}
