/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"footgo/internal/datastructures"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// teamsCmd represents the teams command
var teamsCmd = &cobra.Command{
	Use:   "teams",
	Short: "Teams return",
	Long:  `Provides info about the teams`,
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
	responseBytes, err := getData("/matches")
	if err != nil {
		panic(err)
	}

	var matches datastructures.Team
	err = json.Unmarshal(responseBytes, &matches) // Use = instead of :=
	if err != nil {
		log.Printf("Could not unmarshal response - %v", err)
	}

	for _, arg := range matches.Matches {
		fmt.Println(arg.Area.Name)
	}

	color.Green(string(matches.Matches[0].Area.Name))
}
