package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// fixturesCmd represents the fixtures command
var fixturesCmd = &cobra.Command{
	Use:   "fixtures",
	Short: "football fixtures",
	Long:  `Get upcoming and past fixtures of a league and team`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("fixtures called")
	},
}

func init() {
	rootCmd.AddCommand(fixturesCmd)
}
