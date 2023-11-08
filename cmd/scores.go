package cmd

import (
	"encoding/json"
	"fmt"
	"footgo/internal/datastructures"
	"log"

	"github.com/spf13/cobra"
)

var liveFlag bool
var teamFlag bool

// scoresCmd represents the scores command
var scoresCmd = &cobra.Command{
	Use:   "scores",
	Short: "Get scores of past and live fixtures",
	Long:  `Get scores of past and live fixtures from footgo`,
	Run: func(cmd *cobra.Command, args []string) {
		createHeader()
		var team string

		if liveFlag {
			getScoresLive()
		}
		if teamFlag {
			team = convertClubId(args[len(args)-1])
			getScoresForTeam(team)
		}
	},
}

func init() {
	rootCmd.AddCommand(scoresCmd)

	// Add a required flag
	scoresCmd.Flags().BoolVarP(&liveFlag, "live", "l", false, "Live flag")
	scoresCmd.Flags().BoolVarP(&teamFlag, "team", "t", false, "Team flag")
}

func getScoresForTeam(team string) {
	responseBytesFin, err := getData("teams/" + team + "/matches?status=FINISHED")
	if err != nil {
		panic(err)
	}

	var matchesFin datastructures.Match
	err = json.Unmarshal(responseBytesFin, &matchesFin)
	if err != nil {
		log.Printf("Could not unmarshal response - %v", err)
	}

	responseBytesLive, err := getData("teams/" + team + "/matches?status=LIVE")
	if err != nil {
		panic(err)
	}

	var matchesLive datastructures.Match
	err = json.Unmarshal(responseBytesLive, &matchesLive)
	if err != nil {
		log.Printf("Could not unmarshal response - %v", err)
	}

	for _, arg := range matchesFin.Matches {
		fmt.Println(arg.HomeTeam.ShortName, arg.Score.FullTime.Home, " 	vs	", arg.AwayTeam.ShortName, arg.Score.FullTime.Away)
	}
	for _, arg := range matchesLive.Matches {
		fmt.Println(arg.HomeTeam.ShortName, arg.Score.FullTime.Home, " 	vs	", arg.AwayTeam.ShortName, arg.Score.FullTime.Away)
	}

}

func getScoresLive() {
	responseBytes, err := getData("matches?status=LIVE")
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
