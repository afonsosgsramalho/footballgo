package cmd

import (
	"footgo/scraper"
	"footgo/utils"
	"log"

	"github.com/spf13/cobra"
)

var gameFlag string

// videosCmd represents the videos command
var videosCmd = &cobra.Command{
	Use:   "videos",
	Short: "videos from games",
	Long:  `Get videos from games of specific game`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.CreateHeader()

		if gameFlag != "" {
			searchUrl := "https://sportdaylight.com/wp-content/uploads/"
			if len(gameFlag) > 6 || len(gameFlag) < 6 {
				log.Fatalf("Not allowed length different than 6. The string you provided has size %v", len(gameFlag))
			} else {
				scraper.Scrape(searchUrl, gameFlag)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(videosCmd)

	videosCmd.PersistentFlags().StringVarP(&gameFlag, "game", "g", "", "Game flag")
	videosCmd.MarkPersistentFlagRequired("game")
}
