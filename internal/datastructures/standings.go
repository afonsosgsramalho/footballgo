package datastructures

type Standings struct {
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
		Stages          []string    `json:"stages"`
	} `json:"season"`
	Standings []struct {
		Stage string      `json:"stage"`
		Type  string      `json:"type"`
		Group interface{} `json:"group"`
		Table []struct {
			Position int `json:"position"`
			Team     struct {
				ID        int    `json:"id"`
				Name      string `json:"name"`
				ShortName string `json:"shortName"`
				Tla       string `json:"tla"`
				Crest     string `json:"crest"`
			} `json:"team"`
			PlayedGames    int    `json:"playedGames"`
			Form           string `json:"form"`
			Won            int    `json:"won"`
			Draw           int    `json:"draw"`
			Lost           int    `json:"lost"`
			Points         int    `json:"points"`
			GoalsFor       int    `json:"goalsFor"`
			GoalsAgainst   int    `json:"goalsAgainst"`
			GoalDifference int    `json:"goalDifference"`
		} `json:"table"`
	} `json:"standings"`
}
