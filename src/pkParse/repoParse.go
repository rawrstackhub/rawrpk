package pkParse

import (
	"encoding/json"
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

type PkFile struct {
	name       string
	fileURL    string
	installLoc string
}

var Pack PkFile

func ParseGit(repo []string) {
	fmt.Println("Parsing page:", repo)
	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/", repo[0], repo[1])
	Pack.name = repo[1]

	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Printf("Error fetching repository contents: %s\n", err)
		return
	}
	defer resp.Body.Close()

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
	count := 0
	var rawrpk int
	for i, file := range files {
		if strings.HasSuffix(file.Name, ".rawrpk") {
			fmt.Printf("Found .rawrpk file: %s, Download URL: %s\n", file.Name, file.URL)
			rawrpk = i
			count++
		}
	}

	if count != 1 {
		if count == 0 {
			fmt.Println("No .rawrpk file found in repository")
		} else {
			fmt.Println("Multiple .rawrpk files found in repository")
		}
		return
	} else {
		userDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("Error getting user home directory: %s\n", err)
			return
		}
		Pack.installLoc = userDir + "\\rawrpk\\" + Pack.name
		if err := ParseFile(files[rawrpk].URL); err != nil {
			fmt.Printf("Error parsing .rawrpk file: %s\n", err)
			return
		}
	}
}
