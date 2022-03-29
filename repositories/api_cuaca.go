package repositories

import (
	"math/rand"
	"net/http"
	"time"
	"trialtugastiga/models"

	"github.com/gin-gonic/gin"
)

func CuacaUpdate(ctx *gin.Context) {
	weather := &models.Cuaca{
		Wind:  RandomNumber(),
		Water: RandomNumber(),
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": weather,
	})
}

func RandomNumber() (randNum int) {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 100
	return rand.Intn(max-min+1) + min
}

// func DoEvery(d time.Duration, f func(time.Time)) {
// 	for x := range time.Tick(d) {
// 		f(x)
// 	}
// }

// func GetUpdateCuaca(ctx *gin.Context) {
// 	DoEvery(1*time.Second, CuacaUpdate)
// }
