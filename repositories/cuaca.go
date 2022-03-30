package repositories

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"trialtugastiga/models"
)

func GetDataCuaca() models.Cuacas {
	res, err := http.Get("http://localhost:8085/weather")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	var getCuaca models.Cuacas

	if err := json.Unmarshal(body, &getCuaca); err != nil {
		fmt.Println("Cannot Unmarshal JSON!!")
	}

	return getCuaca
}
