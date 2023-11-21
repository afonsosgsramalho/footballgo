package cmd

import (
	"encoding/json"
	"fmt"
	"footgo/internal/datastructures"
	"footgo/utils"
	"log"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var liveFlag bool
var teamFlag bool
var daysFlag bool
var exportScores bool

// scoresCmd represents the scores command
var scoresCmd = &cobra.Command{
	Use:   "scores",
	Short: "Get scores of past and live fixtures",
	Long:  `Get scores of past and live fixtures from footgo`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.CreateHeader()

		if liveFlag {
			getScoresLive()
		}
		if teamFlag {
			teamTerm := args[utils.IndexOf(args, "-t")+1]
			team := utils.ConvertClubId(teamTerm)
			getScoresForTeam(team)
		}
		if daysFlag {
			//get days
			days_str := args[utils.IndexOf(args, "d")+1]
			days_int, _ := strconv.Atoi(days_str)
			//get current date
			data := time.Now()
			//Subtract days
			newT := data.AddDate(0, 0, -days_int).Format("2006-01-02")
			getScoresByDate(newT)

		}
	},
}

func init() {
	rootCmd.AddCommand(scoresCmd)

	// Add a required flag
	scoresCmd.Flags().BoolVarP(&liveFlag, "live", "l", false, "Live flag")
	scoresCmd.Flags().BoolVarP(&teamFlag, "team", "t", false, "Team flag")
	scoresCmd.Flags().BoolVarP(&daysFlag, "days", "d", false, "Days flag")
	scoresCmd.Flags().BoolVarP(&exportScores, "export", "e", false, "Export flag")

}

func getScoresForTeam(team string) {
	responseBytesFin, err := utils.GetData("teams/" + team + "/matches?status=FINISHED")
	if err != nil {
		panic(err)
	}

	lines := make([]string, 0)

	var matchesFin datastructures.Match
	err = json.Unmarshal(responseBytesFin, &matchesFin)
	if err != nil {
		log.Printf("Could not unmarshal response - %v", err)
	}

	responseBytesLive, err := utils.GetData("teams/" + team + "/matches?status=LIVE")
	if err != nil {
		panic(err)
	}

	var matchesLive datastructures.Match
	err = json.Unmarshal(responseBytesLive, &matchesLive)
	if err != nil {
		log.Printf("Could not unmarshal response - %v", err)
	}

	for _, arg := range matchesFin.Matches {
		score := utils.PrintScores(arg.HomeTeam.ShortName, arg.Score.FullTime.Home, arg.AwayTeam.ShortName, arg.Score.FullTime.Away)
		lines = append(lines, score)
	}
	for _, arg := range matchesLive.Matches {
		score := utils.PrintScores(arg.HomeTeam.ShortName, arg.Score.FullTime.Home, arg.AwayTeam.ShortName, arg.Score.FullTime.Away)
		lines = append(lines, score)
	}

	if exportScores {
		utils.ExportFile("scoresTeam"+team+".txt", lines)
	}
}

func getScoresLive() {
	responseBytes, err := utils.GetData("matches?status=LIVE")
	if err != nil {
		panic(err)
	}

	lines := make([]string, 0)

	var matches datastructures.Match
	err = json.Unmarshal(responseBytes, &matches)
	if err != nil {
		log.Printf("Could not unmarshal response - %v", err)
	}

	if len(matches.Matches) == 0 {
		fmt.Println("No live games at the moment")
	}

	for _, arg := range matches.Matches {
		score := utils.PrintScores(arg.HomeTeam.ShortName, arg.Score.FullTime.Home, arg.AwayTeam.ShortName, arg.Score.FullTime.Away)
		lines = append(lines, score)
	}

	if exportScores {
		utils.ExportFile("scoresLive.txt", lines)
	}
}

func getScoresByDate(date string) {
	responseBytesFin, err := utils.GetData("matches?date=" + date)
	if err != nil {
		panic(err)
	}

	lines := make([]string, 0)

	var matchesFin datastructures.Match
	err = json.Unmarshal(responseBytesFin, &matchesFin)
	if err != nil {
		log.Printf("Could not unmarshal response - %v", err)
	}

	for _, arg := range matchesFin.Matches {
		score := utils.PrintScores(arg.HomeTeam.ShortName, arg.Score.FullTime.Home, arg.AwayTeam.ShortName, arg.Score.FullTime.Away)
		lines = append(lines, score)
	}

	if exportScores {
		utils.ExportFile("scoresDate"+date+".txt", lines)
	}
}
