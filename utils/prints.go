package utils

import (
	"fmt"
	"strconv"

	"github.com/fatih/color"
)

func PrintScores(homeTeam string, homeResult int, awayTeam string, awayResult int) string {
	home_team := fmt.Sprintf("%s: %s", homeTeam, strconv.Itoa(homeResult))
	away_team := fmt.Sprintf("%s: %s", awayTeam, strconv.Itoa(awayResult))
	game := fmt.Sprintf("%s vs %s", home_team, away_team)

	if homeResult > awayResult {
		fmt.Fprintf(color.Output, "%s vs %s \n", color.GreenString(home_team), color.RedString(away_team))
	} else if homeResult < awayResult {
		fmt.Fprintf(color.Output, "%s vs %s \n", color.RedString(home_team), color.GreenString(away_team))
	} else {
		fmt.Fprintf(color.Output, "%s vs %s \n", color.YellowString(home_team), color.YellowString(away_team))
	}

	return game
}

func PrintFixtures() string {
	return "asdfsdf"
}

func PrintStandings() string {
	return "asdfsdf"
}
