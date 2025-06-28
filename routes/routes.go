package routes

import (
	"stock-suggester/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/health", controllers.HealthCheck)
	r.GET("/news", controllers.GetNewsSuggestions)
	return r
}
