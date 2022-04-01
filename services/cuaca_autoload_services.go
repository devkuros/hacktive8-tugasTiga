package services

import (
	"fmt"
	"log"
	"time"
	"trialtugastiga/repositories"

	"github.com/gin-gonic/gin"
)

func DoEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

func CuacaStatus(t time.Time) {
	getWinds := repositories.GetDataCuaca().Status.Wind
	getWaters := repositories.GetDataCuaca().Status.Water

	if getWaters <= 5 {
		fmt.Println("Tingkat Air: ", getWaters)
		log.Println("STATUS AMAN")
	} else if getWaters <= 8 {
		fmt.Println("Tingkat Air: ", getWaters)
		log.Println("STATUS SIAGA")
	} else {
		fmt.Println("Tingkat Air: ", getWaters)
		log.Println("STATUS BAHAYA")
	}

	if getWinds <= 6 {
		fmt.Println("Tingkat Angin: ", getWinds)
		log.Println("STATUS AMAN")
	} else if getWinds <= 15 {
		fmt.Println("Tingkat Angin: ", getWinds)
		log.Println("STATUS SIAGA")
	} else {
		fmt.Println("Tingkat Angin: ", getWinds)
		log.Println("STATUS BAHAYA")
	}
}

func GetCuacaStatus(ctx *gin.Context) {
	DoEvery(15*time.Second, CuacaStatus)
}
