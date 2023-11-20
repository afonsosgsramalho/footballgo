package cmd

import (
	"encoding/json"
	"fmt"
	"footgo/internal/datastructures"
	"footgo/utils"
	"strconv"

	"github.com/spf13/cobra"
)

var standingsFlag bool
var exportStandings bool

// standingsCmd represents the standings command
var standingsCmd = &cobra.Command{
	Use:   "standings",
	Short: "football standings",
	Long:  `Get standings of particular league`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.CreateHeader()
		var league string

		if standingsFlag {
			league = utils.ConvertCompetitionId(args[len(args)-1])
			getStandings(league)
		} else {
			fmt.Println("Provide argument")
		}
	},
}

func init() {
	rootCmd.AddCommand(standingsCmd)

	standingsCmd.Flags().BoolVarP(&standingsFlag, "standings", "s", false, "Standings flag")
	standingsCmd.MarkFlagRequired("standings")
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
			tmp_string := strconv.Itoa(arg2.Position) + " " + arg2.Team.ShortName
			lines = append(lines, tmp_string)
			fmt.Println(lines)
		}
	}

	if exportStandings {
		utils.ExportFile("StandingsCompetition"+competition+".txt", lines)
	}
}
