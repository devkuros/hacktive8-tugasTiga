package repositories

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"trialtugastiga/models"
)

func GetDataCuaca() {
	res, err := http.Get("http://localhost:8085/weather")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	var getCuaca models.Cuaca

	err = json.Unmarshal(body, &getCuaca)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Wind: ", getCuaca.Wind)
	log.Println("Water: ", getCuaca.Water)
}
