package cli

import (
	"flag"
	"fmt"
	"os"
	"rawrpk/internal/common"
	"strings"
)

var loc []string = []string{"", ""}

func CLIparse(pkg *common.PkgData) error {
	flag.Parse()

	if len(flag.Args()) < 2 {
		fmt.Println("Usage: rawrpk get <source>.<repository>")
		os.Exit(1)
	}

	command := flag.Arg(0)
	subCmd := flag.Arg(1)

	switch command {
	case "get":
		r, err := getHndl(subCmd)
		if err != nil {
			return
		}
	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
	return fmt.Errorf("error parsing command: %s", command)
}

func getHndl(subCmd string) error {
	parts := strings.SplitN(subCmd, ".", 2)
	if len(parts) != 2 {
		fmt.Println("Invalid source identifier. Expected format: <source>.<repository>")
		os.Exit(1)
	}

	source := parts[0]
	identifier := parts[1]

	switch source {
	case "github":
		fmt.Printf("Fetching from GitHub: %s\n", identifier)
		loc = strings.Split(identifier, "/")
		return common.Github, nil
	case "url":
		fmt.Printf("Fetching from URL: %s\n", identifier)
	default:
		fmt.Printf("Unsupported source: %s\n", source)
		os.Exit(1)
	}
	return 0, fmt.Errorf("error fetching from source: %s", subCmd)
}
