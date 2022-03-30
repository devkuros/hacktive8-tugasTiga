package models

type Cuacas struct {
	Status struct {
		Wind  int `json:"wind"`
		Water int `json:"water"`
	} `json:"status"`
}
