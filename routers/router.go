package routers

import (
	"trialtugastiga/repositories"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	routers := gin.Default()

	routers.GET("/weather", repositories.CuacaUpdate)

	return routers
}
