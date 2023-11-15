/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/spf13/cobra"
)

var videoFlag bool

// videosCmd represents the videos command
var videosCmd = &cobra.Command{
	Use:   "videos",
	Short: "videos from games",
	Long:  `Get videos from games of specific game`,
	Run: func(cmd *cobra.Command, args []string) {
		if videoFlag {
			scrapeMonth()
		}
	},
}

func init() {
	rootCmd.AddCommand(videosCmd)

	videosCmd.Flags().BoolVarP(&videoFlag, "video", "v", false, "Video flag")

}

func scrapeMonth() {
	var searchString string = "https://sportdaylight.com/wp-content/uploads/"
	links := make([]string, 0)

	y, m, _ := time.Now().Date()
	month := int(m)
	year := strconv.Itoa(int(y))

	for i := 1; i <= month; i++ {
		if i <= 9 {
			links = append(links, searchString+year+"/"+"0"+strconv.Itoa(i))
		} else {
			links = append(links, searchString+year+"/"+strconv.Itoa(i))
		}
	}

	teste_url := "https://sportdaylight.com/wp-content/uploads/2023/04/"
	mp4_links := make([]string, 0)

	// Instantiate default collector
	c := colly.NewCollector()

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.HasSuffix(link, ".mp4") {
			mp4_links = append(mp4_links, link)
		}
	})

	// Start scraping on the specific page
	c.Visit(teste_url)

	//Crate an example for downloading the first mp4 file
	download_file("https://sportdaylight.com" + mp4_links[3])
}
