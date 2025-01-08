package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xarick/golang-memcached-example/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/create", handlers.CreateShortURL)
	r.GET("/:shortCode", handlers.HandleShortURL)

	return r
}
