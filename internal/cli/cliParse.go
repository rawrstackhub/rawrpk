package cli

import (
	"flag"
	"fmt"
	"os"
	"rawrpk/internal/common"
	"rawrpk/internal/gitparse"
	"strings"
)

func CLIparse() ([]string, int8, error) {
	flag.Parse()

	if len(flag.Args()) < 2 {
		fmt.Println("Usage: rawrpk get <source>.<repository>")
		os.Exit(1)
	}

	command := flag.Arg(0)
	sourceIdentifier := flag.Arg(1)

	switch command {
	case "get":
		x, y, z := getHndl(sourceIdentifier)
		return x, y, z
	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
	return nil, 0, fmt.Errorf("error parsing command: %s", command)
}

func getHndl(sourceIdentifier string) ([]string, int8, error) {
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
		gitparse.ParseGit(repository)
		return repository, common.Github, nil
	case "url":
		fmt.Printf("Fetching from URL: %s\n", identifier)
	default:
		fmt.Printf("Unsupported source: %s\n", source)
		os.Exit(1)
	}
	return nil, 0, fmt.Errorf("error fetching from source: %s", sourceIdentifier)
}
