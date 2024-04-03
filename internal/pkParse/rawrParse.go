package pkParse

import (
	"bufio"
	"fmt"
	"net/http"
	"rawrpk/internal/pkSystem"
	"strings"
)

func ParseFile(rawrpkFile string) error {
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
		fmt.Println("Parsing line:", line)
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
		if len(words) != 2 {
			return fmt.Errorf("invalid INSTALL format: %s", line)
		}
		Pack.fileURL = words[1]
		err := pkSystem.Install(Pack.name, Pack.fileURL)
		if err != nil {
			return err
		}
		fmt.Println("Installing from link:", Pack.fileURL)

	case "ENVVAR":
		if len(words) != 4 || words[1] != "ADD" {
			return fmt.Errorf("invalid ENVVAR format: %s", line)
		}
		err := pkSystem.AddPath(Pack.installLoc)
		if err != nil {
			return err
		}
		fmt.Printf("Adding ENVVAR to %s with location: %s\n", words[2], words[3])
	default:
		return fmt.Errorf("unknown instruction: %s", words[0])
	}

	return nil
}
