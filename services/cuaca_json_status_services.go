package services

import (
	"net/http"
	"trialtugastiga/models"
	"trialtugastiga/repositories"

	"github.com/gin-gonic/gin"
)

func JsonStats() models.Stats {
	var getStats models.Stats
	getStatusWater := getStats.StatusWater
	getStatusWind := getStats.StatusWind

	getWaters := repositories.GetDataCuaca().Status.Water
	getWinds := repositories.GetDataCuaca().Status.Wind

	if getWaters <= 5 {
		getStatusWater = "STATUS AMAN"
	} else if getWaters <= 8 {
		getStatusWater = "STATUS SIAGA"
	} else {
		getStatusWater = "STATUS BAHAYA"
	}

	if getWinds <= 6 {
		getStatusWind = "STATUS AMAN"
	} else if getWinds <= 15 {
		getStatusWind = "STATUS SIAGA"
	} else {
		getStatusWind = "STATUS BAHAYA"
	}

	getStats.Water = getWaters
	getStats.Wind = getWinds

	getStats.StatusWater = getStatusWater
	getStats.StatusWind = getStatusWind

	return getStats
}

func GetJsonStats(c *gin.Context) {
	result := JsonStats()

	c.JSON(http.StatusOK, result)
}
