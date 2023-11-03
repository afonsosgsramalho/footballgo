package datastructures

import (
	"time"
)

type Match struct {
	Matches []struct {
		Area struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Code string `json:"code"`
			Flag string `json:"flag"`
		} `json:"area"`
		Competition struct {
			ID     int    `json:"id"`
			Name   string `json:"name"`
			Code   string `json:"code"`
			Type   string `json:"type"`
			Emblem string `json:"emblem"`
		} `json:"competition"`
		Season struct {
			ID              int         `json:"id"`
			StartDate       string      `json:"startDate"`
			EndDate         string      `json:"endDate"`
			CurrentMatchday int         `json:"currentMatchday"`
			Winner          interface{} `json:"winner"`
		} `json:"season"`
		ID          int         `json:"id"`
		UtcDate     time.Time   `json:"utcDate"`
		Status      string      `json:"status"`
		Matchday    int         `json:"matchday"`
		Stage       string      `json:"stage"`
		Group       interface{} `json:"group"`
		LastUpdated time.Time   `json:"lastUpdated"`
		HomeTeam    struct {
			ID        int    `json:"id"`
			Name      string `json:"name"`
			ShortName string `json:"shortName"`
			Tla       string `json:"tla"`
			Crest     string `json:"crest"`
		} `json:"homeTeam"`
		AwayTeam struct {
			ID        int    `json:"id"`
			Name      string `json:"name"`
			ShortName string `json:"shortName"`
			Tla       string `json:"tla"`
			Crest     string `json:"crest"`
		} `json:"awayTeam"`
		Score struct {
			Winner   string `json:"winner"`
			Duration string `json:"duration"`
			FullTime struct {
				Home int `json:"home"`
				Away int `json:"away"`
			} `json:"fullTime"`
			HalfTime struct {
				Home int `json:"home"`
				Away int `json:"away"`
			} `json:"halfTime"`
		} `json:"score"`
		Odds struct {
			Msg string `json:"msg"`
		} `json:"odds"`
		Referees []interface{} `json:"referees"`
	} `json:"matches"`
}
