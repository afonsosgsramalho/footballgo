package cmd

import (
	"encoding/json"
	"footgo/internal/datastructures"
	"footgo/utils"
	"strconv"

	"github.com/spf13/cobra"
)

var leagueFlag string
var exportStandings bool

// standingsCmd represents the standings command
var standingsCmd = &cobra.Command{
	Use:   "standings",
	Short: "football standings",
	Long:  `Get standings of particular league`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.CreateHeader()

		if leagueFlag != "" {
			competition := utils.ConvertCompetitionId(leagueFlag)
			getStandings(competition)
		}
	},
}

func init() {
	rootCmd.AddCommand(standingsCmd)

	standingsCmd.Flags().StringVarP(&leagueFlag, "league", "l", "", "League flag")
	standingsCmd.MarkFlagRequired("league")
	standingsCmd.Flags().BoolVarP(&exportStandings, "export", "e", false, "Export flag")

}

func getStandings(competition string) {
	responseBytes, err := utils.GetData("competitions/" + competition + "/standings")
	if err != nil {
		panic(err)
	}

	lines := make([]string, 0)

	var standings datastructures.Standings
	err = json.Unmarshal(responseBytes, &standings)
	if err != nil {
		panic(err)
	}

	for _, arg := range standings.Standings {
		for _, arg2 := range arg.Table {
			fixture := utils.PrintStandings(strconv.Itoa(arg2.Position), arg2.Team.ShortName, strconv.Itoa(arg2.Points))
			lines = append(lines, fixture)
		}
	}

	if exportStandings {
		utils.ExportFile("StandingsCompetition"+competition+".txt", lines)
	}
}
