/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"footgo/scraper"
	"footgo/utils"
	"log"

	"github.com/spf13/cobra"
)

var gameFlag bool

// videosCmd represents the videos command
var videosCmd = &cobra.Command{
	Use:   "videos",
	Short: "videos from games",
	Long:  `Get videos from games of specific game`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.CreateHeader()

		if gameFlag {
			gamearg := args[utils.IndexOf(args, "-v")+1]
			searchUrl := "https://sportdaylight.com/wp-content/uploads/"

			if len(gamearg) > 6 || len(gamearg) < 6 {
				log.Fatalf("Not allowed length different than 6. The string you provided has size %v", len(gamearg))
			} else {
				scraper.Scrape(searchUrl, gamearg)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(videosCmd)

	videosCmd.PersistentFlags().BoolVarP(&gameFlag, "game", "g", false, "Game flag")
}
