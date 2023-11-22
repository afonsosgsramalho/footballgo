package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/schollz/progressbar/v3"
)

func Download_file(game_url string) {
	// Build fileName from fullPath
	fileURL, err := url.Parse(game_url)
	if err != nil {
		log.Fatal(err)
	}
	path := fileURL.Path
	segments := strings.Split(path, "/")
	fileName := segments[len(segments)-1]
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

func ExportFile(path string, lines []string) {
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
