package utilitaries

import (
	"encoding/json"
	"footgo/config"
	"footgo/internal/datastructures"
	"io"
	"net/http"
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

func GetClubs() map[string]string {
	m := make(map[string]string)

	comps := [3]string{"CL", "PPL", "PL"} //, "DED", "ABL", "FL1"} //"SA", "PD", "ELC", "BSA", "WC", "EC"}

	for _, compCode := range comps {
		competition, err := GetCompetitionData(compCode)
		if err != nil {
			panic(err)
		}

		// Add not only names but also codes
		for _, arg := range competition.Teams {
			m[arg.Name] = strconv.Itoa(arg.ID)
			m[arg.ShortName] = strconv.Itoa(arg.ID)
		}
	}

	return m
}

// leagues
func GetLeagues() map[string]string {
	m := make(map[string]string)

	responseBytes, err := GetData("competitions/")
	if err != nil {
		panic(err)
	}

	var competition datastructures.Competition
	err = json.Unmarshal(responseBytes, &competition)
	if err != nil {
		panic(err)
	}

	// Add not only names but also codes
	for _, arg := range competition.Competitions {
		m[arg.Name] = strconv.Itoa(arg.ID)
		m[arg.Code] = strconv.Itoa(arg.ID)
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
