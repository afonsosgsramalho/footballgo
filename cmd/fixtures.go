package cmd

import (
	"encoding/json"
	"footgo/internal/datastructures"
	"footgo/utils"
	"log"

	"github.com/spf13/cobra"
)

var leagueFlagFixtures string
var teamFlagFixtures string
var exportFixtures bool

// fixturesCmd represents the fixtures command
var fixturesCmd = &cobra.Command{
	Use:   "fixtures",
	Short: "football fixtures",
	Long:  `Get upcoming and past fixtures of a league and team`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.CreateHeader()

		if (len(args) == 0 && leagueFlagFixtures == "") || (len(args) == 0 && teamFlagFixtures == "") {
			log.Fatal("You have to provide at least one argument")
		}

		if leagueFlagFixtures != "" {
			competition := utils.ConvertCompetitionId(leagueFlagFixtures)
			fixturesCompetition(competition)
		}
		if teamFlagFixtures != "" {
			team := utils.ConvertClubId(teamFlagFixtures)
			fixturesTeam(team)
		}
	},
}

func init() {
	rootCmd.AddCommand(fixturesCmd)

	fixturesCmd.Flags().StringVarP(&leagueFlagFixtures, "leagues", "l", "", "League flag")
	fixturesCmd.Flags().StringVarP(&teamFlagFixtures, "teams", "t", "", "Team flag")
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
		fixture := utils.PrintFixtures(arg.Competition.Name, arg.HomeTeam.ShortName, arg.AwayTeam.ShortName, arg.UtcDate)
		lines = append(lines, fixture)
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
		fixture := utils.PrintFixtures(arg.Competition.Name, arg.HomeTeam.ShortName, arg.AwayTeam.ShortName, arg.UtcDate)
		lines = append(lines, fixture)
	}

	if exportFixtures {
		utils.ExportFile("fixturesTeam"+team+".txt", lines)
	}
}
