package routers

import (
	"TouristAppSolution/api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	authorGroup := r.Group("/recommendation")
	{
		authorGroup.GET("/city/:city", controllers.GetDestinationsByCity)
		authorGroup.GET("/landmarks", controllers.GetDestinationsByLandmarks)
	}
	return r
}