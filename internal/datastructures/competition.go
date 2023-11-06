package datastructures

import "time"

type Competition struct {
	Count   int `json:"count"`
	Filters struct {
	} `json:"filters"`
	Competitions []struct {
		ID   int `json:"id"`
		Area struct {
			ID   int         `json:"id"`
			Name string      `json:"name"`
			Code string      `json:"code"`
			Flag interface{} `json:"flag"`
		} `json:"area"`
		Name          string      `json:"name"`
		Code          string      `json:"code"`
		Type          string      `json:"type"`
		Emblem        interface{} `json:"emblem"`
		Plan          string      `json:"plan"`
		CurrentSeason struct {
			ID              int         `json:"id"`
			StartDate       string      `json:"startDate"`
			EndDate         string      `json:"endDate"`
			CurrentMatchday int         `json:"currentMatchday"`
			Winner          interface{} `json:"winner"`
		} `json:"currentSeason"`
		NumberOfAvailableSeasons int       `json:"numberOfAvailableSeasons"`
		LastUpdated              time.Time `json:"lastUpdated"`
	} `json:"competitions"`
}
