package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	githubURL = "https://raw.githubusercontent.com/username/repo/branch/README.md" // Replace with actual URL
	interval  = 10 * time.Second                                                  // Fetch interval
)

func fetchReadme(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch the README.md: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read the response body: %v", err)
	}

	return string(body), nil
}

func extractTargetString(content, target string) (string, error) {
	if strings.Contains(content, target) {
		return target, nil
	}
	return "", fmt.Errorf("target string '%s' not found", target)
}

func main() {
	var previousContent string

	for {
		content, err := fetchReadme(githubURL)
		if err != nil {
			log.Printf("Error fetching README.md: %v", err)
			time.Sleep(interval)
			continue
		}

		targetString, err := extractTargetString(content, "Hello World")
		if err != nil {
			log.Printf("Error extracting target string: %v", err)
		} else if targetString != previousContent {
			log.Printf("Current string: %s", targetString)
			previousContent = targetString
		}

		time.Sleep(interval)
	}
}
