package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var daysFlag bool

// fixturesCmd represents the fixtures command
var fixturesCmd = &cobra.Command{
	Use:   "fixtures",
	Short: "football fixtures",
	Long:  `Get upcoming and past fixtures of a league and team`,
	Run: func(cmd *cobra.Command, args []string) {
		if daysFlag {
			fmt.Println("fixtures called")
		} else {
			fmt.Println("no fixtures")
		}
	},
}

func init() {
	rootCmd.AddCommand(fixturesCmd)

	fixturesCmd.Flags().BoolVarP(&daysFlag, "days", "d", false, "Days flag")
}
