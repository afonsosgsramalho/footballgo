package cmd

import (
	"encoding/json"
	"fmt"
	"footgo/internal/datastructures"
	"log"

	"github.com/spf13/cobra"
)

var leagueFlagFixtures bool
var teamFlagFixtures bool

// fixturesCmd represents the fixtures command
var fixturesCmd = &cobra.Command{
	Use:   "fixtures",
	Short: "football fixtures",
	Long:  `Get upcoming and past fixtures of a league and team`,
	Run: func(cmd *cobra.Command, args []string) {
		if leagueFlagFixtures {
			fixturesCompetition("2017")
		}
		if teamFlagFixtures {
			fixturesTeam("5543")
		}
	},
}

func init() {
	rootCmd.AddCommand(fixturesCmd)

	fixturesCmd.Flags().BoolVarP(&leagueFlagFixtures, "leagues", "l", false, "League flag")
	fixturesCmd.Flags().BoolVarP(&teamFlagFixtures, "teams", "t", false, "Team flag")
}

func fixturesCompetition(comp string) {
	responseBytesFin, err := getData("competitions/" + comp + "/matches?status=SCHEDULED")
	if err != nil {
		panic(err)
	}

	var competition datastructures.Match
	err = json.Unmarshal(responseBytesFin, &competition)
	if err != nil {
		log.Printf("Could not unmarshal response - %v", err)
	}

	for _, arg := range competition.Matches {
		fmt.Println(arg.HomeTeam.ShortName, " vs ", arg.AwayTeam.ShortName, " at ", arg.UtcDate)
	}
}

func fixturesTeam(team string) {
	responseBytesFin, err := getData("teams/" + team + "/matches?status=SCHEDULED")
	if err != nil {
		panic(err)
	}

	var matches datastructures.Match
	err = json.Unmarshal(responseBytesFin, &matches)
	if err != nil {
		log.Printf("Could not unmarshal response - %v", err)
	}

	for _, arg := range matches.Matches {
		fmt.Println(arg.HomeTeam.ShortName, " vs ", arg.AwayTeam.ShortName, " at ", arg.UtcDate)
	}
}
