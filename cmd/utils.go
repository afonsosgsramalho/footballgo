package cmd

import (
	"encoding/json"
	"fmt"
	"footgo/config"
	"footgo/internal/datastructures"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/schollz/progressbar/v3"
)

func createHeader() {
	color.Green(
		` 
			  _____                __     ____          
			_/ ____\____    ____ _/  |_  / ___\  ____   
			\   __\/  _ \  /  _ \\   __\/ /_/ > /  _ \  
			|  |  (  <_> )(  <_> )|  |  \___  /(  <_> ) 
			|__|   \____/  \____/ |__| /_____/  \____/	
																	 
		`)
}

func getData(endpoint string) ([]byte, error) {
	endpointAux := config.APIEndpoint + endpoint
	request, err := http.NewRequest(http.MethodGet, endpointAux, nil)
	if err != nil {
		panic(err)
	}

	configData, err := config.LoadConfig("config/config.json")
	if err != nil {
		panic(err)
	}

	// Set the "X-Auth-Token" header with your API token
	request.Header.Set("X-Auth-Token", configData.APIToken)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	return responseBytes, nil
}

// Define a function to fetch competition data for a given competition code.
func getCompetitionData(compCode string) (*datastructures.Team, error) {
	url := "competitions/" + compCode + "/teams"

	responseBytes, err := getData(url)
	if err != nil {
		return nil, err
	}

	var team datastructures.Team
	err = json.Unmarshal(responseBytes, &team)
	if err != nil {
		return nil, err
	}

	return &team, nil
}

func getClubs() map[string]string {
	m := make(map[string]string)

	comps := [3]string{"CL", "PPL", "PL"} //, "DED", "ABL", "FL1"} //"SA", "PD", "ELC", "BSA", "WC", "EC"}

	for _, compCode := range comps {
		competition, err := getCompetitionData(compCode)
		if err != nil {
			panic(err)
		}

		// Add not only names but also codes
		for _, arg := range competition.Teams {
			m[arg.Name] = strconv.Itoa(arg.ID)
			m[arg.ShortName] = strconv.Itoa(arg.ID)
		}
	}

	return m
}

// leagues
func getLeagues() map[string]string {
	m := make(map[string]string)

	responseBytes, err := getData("competitions/")
	if err != nil {
		panic(err)
	}

	var competition datastructures.Competition
	err = json.Unmarshal(responseBytes, &competition)
	if err != nil {
		panic(err)
	}

	// Add not only names but also codes
	for _, arg := range competition.Competitions {
		m[arg.Name] = strconv.Itoa(arg.ID)
		m[arg.Code] = strconv.Itoa(arg.ID)
	}

	return m
}

// TODO: cache the results, intead always querying API
func convertClubId(club string) string {
	clubs := getClubs()
	return clubs[club]
}

// TODO: cache the results, intead always querying API
func convertCompetitionId(competition string) string {
	competitions := getLeagues()
	return competitions[competition]
}

func exportFile(path string, lines []string) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for _, line := range lines {
		_, err := f.WriteString(line + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}

func indexOf(arr []string, val string) int {
	for pos, v := range arr {
		if v == val {
			return pos
		}
	}
	return -1
}

func download_file(game_url string) {
	// Build fileName from fullPath
	fileURL, err := url.Parse(game_url)
	if err != nil {
		log.Fatal(err)
	}
	path := fileURL.Path
	segments := strings.Split(path, "/")
	fileName := segments[len(segments)-1]
	// filePath := "/home/vboxuser/Videos/" + fileName
	filePath := fileName

	// Create blank file
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	resp, err := client.Get(game_url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Put content on file
	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		"downloading",
	)

	size, err := io.Copy(io.MultiWriter(file, bar), resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	fmt.Printf("Downloaded file %s with size %d %s", fileName, size/(1<<20), "MB")
	fmt.Print("\n")

}

//WORD SIMILARITY UTILS

func word_similarity_Levenshtein(word1 string, word2 string) int {
	pre := make([]int, len(word2)+1)
	cur := make([]int, len(word2)+1)

	for i := 1; i <= len(word1); i++ {
		cur[0] = 1
		for j := 1; j < len(pre); j++ {
			if word1[i-1] != word2[j-1] {
				cur[j] = min(cur[j-1], pre[j-1], pre[j]) + 1
			} else {
				cur[j] = pre[j-1]
			}
		}

		tmp := make([]int, len(cur))
		copy(tmp, cur)
		pre = tmp
	}

	return pre[len(word2)]
}

func distance_words_ratio(word1 string, word2 string) float64 {
	distance := word_similarity_Levenshtein(word1, word2)
	maxLen := max(len(word1), len(word2))

	return 1 - float64(distance)/float64(maxLen)
}
