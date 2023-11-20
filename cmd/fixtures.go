package cmd

import (
	"encoding/json"
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
			competitionTerm := args[utils.IndexOf(args, "-l")+1]
			competition := utils.ConvertCompetitionId(competitionTerm)
			fixturesCompetition(competition)
		}
		if teamFlagFixtures {
			teamTerm := args[utils.IndexOf(args, "-t")+1]
			team := utils.ConvertClubId(teamTerm)
			fixturesTeam(team)
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
