package routers

import (
	"trialtugastiga/repositories"
	"trialtugastiga/services"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	routers := gin.Default()

	cuacaRouter := routers.Group("/weather")
	{
		cuacaRouter.GET("/", repositories.CuacaUpdate)
		cuacaRouter.GET("/stats", services.GetCuacaStatus)
		cuacaRouter.GET("/json", services.GetJsonStats)
	}

	return routers
}
