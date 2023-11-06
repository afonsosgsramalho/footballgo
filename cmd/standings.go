package cmd

import (
	"encoding/json"
	"fmt"
	"footgo/internal/datastructures"

	"github.com/spf13/cobra"
)

var standingsFlag bool

// standingsCmd represents the standings command
var standingsCmd = &cobra.Command{
	Use:   "standings",
	Short: "football standings",
	Long:  `Get standings of particular league`,
	Run: func(cmd *cobra.Command, args []string) {
		createHeader()
		if standingsFlag {
			getStandings("2021")
		} else {
			fmt.Println("Provide argument")
		}
	},
}

func init() {
	rootCmd.AddCommand(standingsCmd)

	standingsCmd.Flags().BoolVarP(&standingsFlag, "standings", "s", false, "Standings flag")
}

func getStandings(team string) {
	responseBytes, err := getData("competitions/" + team + "/standings")
	if err != nil {
		panic(err)
	}

	var standings datastructures.Standings
	err = json.Unmarshal(responseBytes, &standings)
	if err != nil {
		panic(err)
	}

	for _, arg := range standings.Standings {
		for _, arg2 := range arg.Table {
			fmt.Println(arg2.Position, arg2.Team.ShortName)
		}
	}
}
