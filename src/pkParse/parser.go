package pkParse

import (
	"bufio"
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
	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/", repo[0], repo[1])

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
		if err := RawrpkParse(files[rawrpk].URL); err != nil {
			fmt.Printf("Error parsing .rawrpk file: %s\n", err)
			return
		}
	}
}

func RawrpkParse(rawrpkFile string) error {
	fmt.Println("Parsing File:", rawrpkFile)
	response, err := http.Get(rawrpkFile)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("error fetching file: %s", response.Status)
	}

	scanner := bufio.NewScanner(response.Body)

	for scanner.Scan() {
		line := scanner.Text()
		if err := parseLine(line); err != nil {
			fmt.Printf("Error parsing line: %s\n", err)
			continue
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func parseLine(line string) error {
	words := strings.Fields(line)

	if len(words) < 2 {
		return fmt.Errorf("invalid line format: %s", line)
	}

	switch words[0] {
	case "INSTALL":
		fmt.Println("Installing from link:", words[1])

	case "ENVVAR":
		if len(words) != 4 || words[1] != "ADD" {
			return fmt.Errorf("invalid ENVVAR format: %s", line)
		}
		fmt.Printf("Adding ENVVAR to %s with location: %s\n", words[2], words[3])
	default:
		return fmt.Errorf("unknown instruction: %s", words[0])
	}

	return nil
}
