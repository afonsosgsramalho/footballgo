package scraper

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func downloadFile(urlVideo string, folderPath interface{}) {
	urlVideo = "https://sportdaylight.com/wp-content/uploads/2023/11/man-city-vs-Bournemouth-6-1-highlights-SPORTDAYLIGHT.COM_.mp4"

	// Build fileName from fullPath
	fileURL, err := url.Parse(urlVideo)
	if err != nil {
		log.Fatal(err)
	}
	path := fileURL.Path
	segments := strings.Split(path, "/")
	fileName := path + segments[len(segments)-1]

	var filePath string
	//If theres no path, download file to user current path
	if str, ok := folderPath.(string); ok {
		filePath = str + fileName
	} else {
		filePath = ""
	}

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
	// Put content on file
	resp, err := client.Get(urlVideo)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	fmt.Printf("Downloaded a file %s with size %d", fileName, size)
}
