package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// videosCmd represents the videos command
var videosCmd = &cobra.Command{
	Use:   "videos",
	Short: "game videos",
	Long:  `Get videos of specific game for the current month, only if the video is available`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("videos called")
	},
}

func init() {
	rootCmd.AddCommand(videosCmd)
}
