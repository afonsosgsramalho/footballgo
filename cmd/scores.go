/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"footgo/internal/datastructures"
	"log"

	"github.com/spf13/cobra"
)

var liveFlag bool

// scoresCmd represents the scores command
var scoresCmd = &cobra.Command{
	Use:   "scores",
	Short: "Get scores of past and live fixtures",
	Long:  `Get scores of past and live fixtures from footgo`,
	Run: func(cmd *cobra.Command, args []string) {
		if liveFlag {
			getScores()
		}
	},
}

func init() {
	rootCmd.AddCommand(scoresCmd)

	// Add a required flag
	scoresCmd.Flags().BoolVarP(&liveFlag, "live", "l", false, "A required flag")
	scoresCmd.MarkFlagRequired("live")
}

func getScores() {
	responseBytes, err := getData("/matches?status=LIVE")
	if err != nil {
		panic(err)
	}

	var matches datastructures.Match
	err = json.Unmarshal(responseBytes, &matches) // Use = instead of :=
	if err != nil {
		log.Printf("Could not unmarshal response - %v", err)
	}

	if len(matches.Matches) == 0 {
		fmt.Println("No live games at the moment")
	}

	for _, arg := range matches.Matches {
		fmt.Println(arg.HomeTeam.ShortName, arg.Score.FullTime.Home, " vs ", arg.AwayTeam.ShortName, arg.Score.FullTime.Away)
	}
}

func isLive(game string) (bool, error) {
	return true, nil
}
