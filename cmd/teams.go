/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"footgo/internal/datastructures"
	"footgo/config"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)


// teamsCmd represents the teams command
var teamsCmd = &cobra.Command{
	Use:   "teams",
	Short: "Teams return",
	Long: `Provides info about the teams`,
	//has the cmd args passed, iterate over them
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			fmt.Println(arg)
		}
		getMatches()
		
	},
}


func init() {
	rootCmd.AddCommand(teamsCmd)

	teamsCmd.Flags().BoolP("teams help", "e", true, "Teams help message")
}


func getMatches() {
	responseBytes, err := getMatchesData("/matches")
	if err != nil {
		panic(err)
	}

	var matches datastructures.Team
	err = json.Unmarshal(responseBytes, &matches) // Use = instead of :=
	if err != nil {
		log.Printf("Could not unmarshal response - %v", err)
	}

	for _, arg:= range matches.Matches {
		fmt.Println(arg.Area.Name)
	}

	color.Green(string(matches.Matches[30].Area.Name))
}


func getMatchesData(endpoint string) ([]byte, error) {
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
