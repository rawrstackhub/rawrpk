package pkParse

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type RepoFile struct {
	Name string `json:"name"`
	Path string `json:"path"`
	URL  string `json:"download_url"`
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
