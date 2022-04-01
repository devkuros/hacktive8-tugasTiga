package models

type Cuacas struct {
	Status struct {
		Wind  int `json:"wind"`
		Water int `json:"water"`
	} `json:"status"`
}

type Stats struct {
	Wind        int    `json:"wind"`
	StatusWind  string `json:"status_wind"`
	Water       int    `json:"water"`
	StatusWater string `json:"status_water"`
}
