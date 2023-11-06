package cmd

import (
	"encoding/json"
	"footgo/config"
	"footgo/internal/datastructures"
	"io"
	"net/http"
	"strconv"

	"github.com/fatih/color"
)

func createHeader() {
	color.Green(
		` 
			  _____                __     ____          
			_/ ____\____    ____ _/  |_  / ___\  ____   
			\   __\/  _ \  /  _ \\   __\/ /_/ > /  _ \  
			|  |  (  <_> )(  <_> )|  |  \___  /(  <_> ) 
			|__|   \____/  \____/ |__| /_____/  \____/	
																	 
		`)
}

func getData(endpoint string) ([]byte, error) {
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

func getClubs() map[string]string {
	m := make(map[string]string)

	responseBytes, err := getData("teams/")
	if err != nil {
		panic(err)
	}

	var team datastructures.Team
	err = json.Unmarshal(responseBytes, &team)
	if err != nil {
		panic(err)
	}

	// Add not only names but also codes
	for _, arg := range team.Teams {
		m[arg.Name] = strconv.Itoa(arg.ID)
		m[arg.ShortName] = strconv.Itoa(arg.ID)
	}

	return m

}

func getLeagues() map[string]string {
	m := make(map[string]string)

	responseBytes, err := getData("competitions/")
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
func convertClubId(club string) string {
	clubs := getClubs()
	return clubs[club]
}

// TODO: cache the results, intead always querying API
func convertCompetitionId(competition string) string {
	competitions := getLeagues()
	return competitions[competition]
}
