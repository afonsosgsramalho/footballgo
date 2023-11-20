package utils

import (
	"encoding/json"
	"footgo/config"
	"footgo/internal/datastructures"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func GetData(endpoint string) ([]byte, error) {
	endpointAux := config.APIEndpoint + endpoint
	request, err := http.NewRequest(http.MethodGet, endpointAux, nil)
	if err != nil {
		panic(err)
	}

	configData, err := config.LoadConfig("config/config.json")
	if err != nil {
		panic(err)
	}

	// Set the "X-Auth-Token" header with your API token
	request.Header.Set("X-Auth-Token", configData.APIToken)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	return responseBytes, nil
}

// Define a function to fetch competition data for a given competition code.
func GetCompetitionData(compCode string) (*datastructures.Team, error) {
	url := "competitions/" + compCode + "/teams"

	responseBytes, err := GetData(url)
	if err != nil {
		return nil, err
	}

	var team datastructures.Team
	err = json.Unmarshal(responseBytes, &team)
	if err != nil {
		return nil, err
	}

	return &team, nil
}

type Teams struct {
	Team []struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		ShortName string `json:"shortName"`
	} `json:"team"`
}

func GetClubs() map[string]string {
	jsonData, err := os.ReadFile("utils/teams.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	m := make(map[string]string)
	var teamsData Teams
	err = json.Unmarshal([]byte(jsonData), &teamsData)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	for _, team := range teamsData.Team {
		m[team.Name] = strconv.Itoa(team.ID)
		m[team.ShortName] = strconv.Itoa(team.ID)
	}

	return m
}

type Competitions struct {
	Competition []struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		ShortName string `json:"shortName"`
	} `json:"team"`
}

// leagues
func GetLeagues() map[string]string {
	jsonData, err := os.ReadFile("utils/competitions.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var competitionsData Competitions
	m := make(map[string]string)
	err = json.Unmarshal([]byte(jsonData), &competitionsData)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	for _, comp := range competitionsData.Competition {
		m[comp.Name] = strconv.Itoa(comp.ID)
		m[comp.ShortName] = strconv.Itoa(comp.ID)
	}

	return m
}

// TODO: cache the results, intead always querying API
func ConvertClubId(club string) string {
	clubs := GetClubs()
	return clubs[club]
}

// TODO: cache the results, intead always querying API
func ConvertCompetitionId(competition string) string {
	competitions := GetLeagues()
	return competitions[competition]
}
