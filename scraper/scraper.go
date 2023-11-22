package scraper

import (
	"fmt"
	"footgo/utils"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly"
	"github.com/schollz/closestmatch"
)

var (
	mp4_links []string
	mutex     sync.Mutex
)

func _parseLinks(urls []string) map[string]string {
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

func _scrapeLink(link string) {
	c := colly.NewCollector()

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.HasSuffix(link, ".mp4") {
			link, _ = url.QueryUnescape(link)
			// Critical section
			mutex.Lock()
			mp4_links = append(mp4_links, link)
			mutex.Unlock()
		}
	})

	// Start scraping on the specific page
	c.Visit(link)
}

func Scrape(url string, game string) {
	links := make([]string, 0)

	y, m, _ := time.Now().Date()
	month := int(m)
	year := strconv.Itoa(int(y))

	if month > 6 {
		for i := 8; i <= month; i++ {
			if i <= 9 {
				links = append(links, url+year+"/"+"0"+strconv.Itoa(i))
			} else {
				links = append(links, url+year+"/"+strconv.Itoa(i))
			}
		}
	} else {
		year_ly := y - 1
		for i := 6; i <= 12; i++ {
			if i <= 9 {
				links = append(links, url+strconv.Itoa(int(year_ly))+"/"+"0"+strconv.Itoa(i))
			} else {
				links = append(links, url+strconv.Itoa(int(year_ly))+"/"+strconv.Itoa(i))
			}
		}
		for i := 1; i <= month; i++ {
			links = append(links, url+year+"/"+"0"+strconv.Itoa(i))
		}
	}

	var wg sync.WaitGroup

	for _, link := range links {
		wg.Add(1)
		go func(link string) {
			defer wg.Done()
			_scrapeLink(link)
		}(link)
	}

	wg.Wait()

	parsedLinks := _parseLinks(mp4_links)

	keys := make([]string, 0, len(parsedLinks))
	for key := range parsedLinks {
		keys = append(keys, key)
	}

	game_converted := _parseGame(game)

	// Closest match
	bagSizes := []int{2}
	cm := closestmatch.New(keys, bagSizes)
	closest := cm.Closest(game_converted)

	word_sim := utils.Distance_words_ratio(game_converted, closest)
	if word_sim < 0.7 {
		fmt.Println("No matching game")
		return
	}

	// Download the file
	utils.Download_file("https://sportdaylight.com" + parsedLinks[closest])
}
