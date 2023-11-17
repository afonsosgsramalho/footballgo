/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/schollz/closestmatch"
	"github.com/spf13/cobra"
)

var videoFlag bool

// videosCmd represents the videos command
var videosCmd = &cobra.Command{
	Use:   "videos",
	Short: "videos from games",
	Long:  `Get videos from games of specific game`,
	Run: func(cmd *cobra.Command, args []string) {
		createHeader()

		if videoFlag {
			gamearg := args[indexOf(args, "-v")+1]
			scrapeMonth(gamearg)
		}
	},
}

func init() {
	rootCmd.AddCommand(videosCmd)

	videosCmd.Flags().BoolVarP(&videoFlag, "video", "v", false, "Video flag")
}

func parseGame(game string) string {
	teamNameMapping := map[string]string{
		"ars": "Arsenal",
		"bur": "Burnley",
		"bha": "Brighton",
		"cry": "Crystal Palace",
		"mun": "Manchester United",
		"avl": "Aston Villa",
		"eve": "Everton",
		"bou": "Bournemouth",
		"liv": "Liverpool",
		"che": "Chelsea",
		"mci": "Manchester City",
		"ful": "Fulham",
		"bre": "Brentford",
		"lut": "Luton Town",
		"new": "Newcastle",
		"nfo": "Nottingham",
		"tot": "Tottenham",
		"whu": "West Ham",
		"wol": "Wolves",
		"shu": "Sheffield United",
	}

	if len(game) > 6 || len(game) < 6 {
		fmt.Errorf("Not allowed length different than 6. you had %v", len(game))
		return "s"

	}

	homeTeam := teamNameMapping[game[0:3]]
	awayTeam := teamNameMapping[game[3:]]

	// Build similar string to the one of the website
	gameString := homeTeam + "-vs-" + awayTeam

	return gameString
}

func parseLinks(urls []string) map[string]string {
	links := make(map[string]string)

	for _, url_aux := range urls {
		url_temp := strings.ToLower(url_aux)
		//all combinations of the words highlight
		url_temp = strings.Split(url_temp, "-highlights")[0]
		url_temp = strings.Split(url_temp, "-highlight")[0]
		url_temp = strings.Split(url_temp, "-hіghlіghts")[0]
		url_temp = strings.Split(url_temp, "/")[5]
		url_temp = url_temp[0 : len(url_temp)-4]
		links[url_temp] = url_aux
	}
	return links
}

func scrapeMonth(game string) {
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
			//unescape link
			link, _ = url.QueryUnescape(link)
			mp4_links = append(mp4_links, link)
		}
	})

	// Start scraping on the specific page
	c.Visit(teste_url)

	parsedLinks := parseLinks(mp4_links)

	keys := make([]string, 0, len(parsedLinks))
	for key := range parsedLinks {
		keys = append(keys, key)
	}

	game_converted := parseGame(game)
	fmt.Println(game_converted, "converted")

	// Choose a set of bag sizes, more is more accurate but slower
	bagSizes := []int{2}
	// Create a closestmatch object
	cm := closestmatch.New(keys, bagSizes)
	closest := cm.Closest(game_converted)
	fmt.Println(cm.Closest(closest), "closest")

	// Download the file
	download_file("https://sportdaylight.com" + parsedLinks[closest])
}
