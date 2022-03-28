package main

import (
	"log"
	"trialtugastiga/configs"
	"trialtugastiga/routers"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conf := configs.GetEnvs()

	var PORT = conf.Port

	routers.StartServer().Run(PORT)
}
