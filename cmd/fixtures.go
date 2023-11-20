package cmd

import (
	"encoding/json"
	"fmt"
	"footgo/internal/datastructures"
	"footgo/utils"
	"log"

	"github.com/spf13/cobra"
)

var leagueFlagFixtures bool
var teamFlagFixtures bool
var exportFixtures bool

// fixturesCmd represents the fixtures command
var fixturesCmd = &cobra.Command{
	Use:   "fixtures",
	Short: "football fixtures",
	Long:  `Get upcoming and past fixtures of a league and team`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.CreateHeader()
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
	fixturesCmd.Flags().BoolVarP(&exportFixtures, "export", "e", false, "Export flag")
}

func fixturesCompetition(comp string) {
	responseBytesFin, err := utils.GetData("competitions/" + comp + "/matches?status=SCHEDULED")
	if err != nil {
		panic(err)
	}

	lines := make([]string, 0)

	var competition datastructures.Match
	err = json.Unmarshal(responseBytesFin, &competition)
	if err != nil {
		log.Printf("Could not unmarshal response - %v", err)
	}

	for _, arg := range competition.Matches {
		tmp_string := arg.Competition.Name + " " + arg.HomeTeam.ShortName + " vs " + arg.AwayTeam.ShortName + " at " + arg.UtcDate.Format("2006-01-02 15:04:05")
		lines = append(lines, tmp_string)
		fmt.Println(tmp_string)
	}

	if exportFixtures {
		utils.ExportFile("fixturesCompetition"+comp+".txt", lines)
	}
}

func fixturesTeam(team string) {
	responseBytesFin, err := utils.GetData("teams/" + team + "/matches?status=SCHEDULED")
	if err != nil {
		panic(err)
	}

	lines := make([]string, 0)

	var matches datastructures.Match
	err = json.Unmarshal(responseBytesFin, &matches)
	if err != nil {
		log.Printf("Could not unmarshal response - %v", err)
	}

	for _, arg := range matches.Matches {
		tmp_string := arg.Competition.Name + " " + arg.HomeTeam.ShortName + " vs " + arg.AwayTeam.ShortName + " at " + arg.UtcDate.Format("2006-01-02 15:04:05")
		lines = append(lines, tmp_string)
		fmt.Println(tmp_string)
	}

	if exportFixtures {
		utils.ExportFile("fixturesTeam"+team+".txt", lines)
	}
}
