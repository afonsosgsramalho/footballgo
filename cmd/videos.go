<<<<<<< HEAD
=======
/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
>>>>>>> 111bb82dbadce9ebe449dd3ff6593fd240fae50e
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// videosCmd represents the videos command
var videosCmd = &cobra.Command{
	Use:   "videos",
<<<<<<< HEAD
	Short: "game videos",
	Long:  `Get videos of specific game for the current month, only if the video is available`,
=======
	Short: "team videos",
	Long:  `Get videos of teams`,
>>>>>>> 111bb82dbadce9ebe449dd3ff6593fd240fae50e
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("videos called")
	},
}

func init() {
	rootCmd.AddCommand(videosCmd)
}
