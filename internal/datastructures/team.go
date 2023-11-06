package datastructures

import "time"

type Team struct {
	Count   int `json:"count"`
	Filters struct {
		Limit      int    `json:"limit"`
		Offset     int    `json:"offset"`
		Permission string `json:"permission"`
	} `json:"filters"`
	Teams []struct {
		ID          int       `json:"id"`
		Name        string    `json:"name"`
		ShortName   string    `json:"shortName"`
		Tla         string    `json:"tla"`
		Crest       string    `json:"crest"`
		Address     string    `json:"address"`
		Website     string    `json:"website"`
		Founded     int       `json:"founded"`
		ClubColors  string    `json:"clubColors"`
		Venue       string    `json:"venue"`
		LastUpdated time.Time `json:"lastUpdated"`
	} `json:"teams"`
}
