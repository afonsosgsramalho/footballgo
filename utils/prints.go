package utils

import (
	"fmt"
	"strconv"
	"time"

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

func PrintFixtures(competition string, homeTeam string, awayTeam string, gameTime time.Time) string {
	fixture := fmt.Sprintf("%s %s vs %s at %s \n", competition, homeTeam, awayTeam, gameTime.Format("2006-01-02 15:04"))
	fmt.Fprintf(color.Output, "%s %s vs %s at %s \n", color.GreenString(competition), color.BlueString(homeTeam), color.RedString(awayTeam), color.YellowString(gameTime.Format("2006-01-02 15:04")))

	return fixture
}

func PrintStandings(position string, team string, points string) string {
	standing := fmt.Sprintf("%s %s with %s points \n", position, team, points)
	fmt.Fprintf(color.Output, "%s %s with %s points \n", color.GreenString(position), color.BlueString(team), color.YellowString(points))

	return standing
}
