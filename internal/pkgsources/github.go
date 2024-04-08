package pkgSrc

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"rawrpk/internal/common"
	"strings"
)

func Github(repo *common.PkgData) {
	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/", repo.Source[1], repo.Title)

	resp, err := http.Get(apiURL)
	if err != nil {
		panic("Error fetching repository contents: " + err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic("Error reading response body: " + err.Error())
	}

	var files []common.RepoFile
	if err := json.Unmarshal(body, &files); err != nil {
		fmt.Printf("Error parsing JSON: %s\n", err)
	}

	rawrpkgs := 0
	var rawrpk int
	for i, file := range files {
		if strings.HasSuffix(file.Name, ".rawrpk") {
			fmt.Printf("Found .rawrpk file: %s, Download URL: %s\n", file.Name, file.URL)
			rawrpk = i
			rawrpkgs++
		}
	}

	if rawrpkgs != 1 {
		if rawrpkgs == 0 {
			fmt.Println("No .rawrpk file found in repository")
		} else {
			fmt.Println("Multiple .rawrpk files found in repository")
		}
		return
	} else {
		userDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("Error getting user home directory: %s\n", err)
		}
		repo.InstallLoc = userDir + "\\rawrpk\\" + repo.Title
		repo.PkgURL = files[rawrpk].URL

		fmt.Println("Install Location: " + repo.InstallLoc + "\nSourcing URL: " + repo.PkgURL)
	}
}
