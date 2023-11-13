/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// videosCmd represents the videos command
var videosCmd = &cobra.Command{
	Use:   "videos",
	Short: "team videos",
	Long:  `Get videos of teams`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("videos called")
	},
}

func init() {
	rootCmd.AddCommand(videosCmd)
}
