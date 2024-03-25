package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type RepoFile struct {
	Name string `json:"name"`
	Path string `json:"path"`
	URL  string `json:"download_url"`
}

func main() {
	flag.Parse()

	if len(flag.Args()) < 2 {
		fmt.Println("Usage: rawrpk get <source>.<repository>")
		os.Exit(1)
	}

	command := flag.Arg(0)
	sourceIdentifier := flag.Arg(1)

	switch command {
	case "get":
		handleGetCommand(sourceIdentifier)
	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}

func handleGetCommand(sourceIdentifier string) {
	parts := strings.SplitN(sourceIdentifier, ".", 2)
	if len(parts) != 2 {
		fmt.Println("Invalid source identifier. Expected format: <source>.<repository>")
		os.Exit(1)
	}

	source := parts[0]
	identifier := parts[1]

	switch source {
	case "github":
		fmt.Printf("Fetching from GitHub: %s\n", identifier)
		repository := strings.Split(identifier, "/")
		ParseGit(repository)
	case "url":
		fmt.Printf("Fetching from URL: %s\n", identifier)
	default:
		fmt.Printf("Unsupported source: %s\n", source)
		os.Exit(1)
	}
}

func ParseGit(repo []string) {
	fmt.Println("Parsing page:", repo)

	// GitHub API endpoint to list repository contents
	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/", repo[0], repo[1])

	// Make an HTTP GET request
	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Printf("Error fetching repository contents: %s\n", err)
		return
	}
	defer resp.Body.Close()

	// Read and parse the JSON response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err)
		return
	}

	var files []RepoFile
	if err := json.Unmarshal(body, &files); err != nil {
		fmt.Printf("Error parsing JSON: %s\n", err)
		return
	}

	// Look for files ending with .rawrpk
	for _, file := range files {
		if strings.HasSuffix(file.Name, ".rawrpk") {
			fmt.Printf("Found .rawrpk file: %s, Download URL: %s\n", file.Name, file.URL)
		}
	}
}
