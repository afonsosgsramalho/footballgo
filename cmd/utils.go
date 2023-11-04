package cmd

import (
	"footgo/config"
	"io"
	"net/http"

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
