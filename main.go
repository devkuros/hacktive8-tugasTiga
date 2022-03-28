package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type Cuaca struct {
	Wind  int `json:"wind"`
	Water int `json:"water"`
}

func DoEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

func CuacaUpdate(t time.Time) {
	weather := &Cuaca{
		Wind:  8,
		Water: 9,
	}

	data, err := json.Marshal(weather)
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	fmt.Println(string(data))
}

func RandomNumber() (randNum int) {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 100
	return rand.Intn(max-min+1) + min
}

func main() {
	DoEvery(1*time.Second, CuacaUpdate)
}
