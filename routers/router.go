package routers

import (
	"trialtugastiga/handlers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	routers := gin.Default()

	routers.GET("/", handlers.CuacaUpdate)

	return routers
}
