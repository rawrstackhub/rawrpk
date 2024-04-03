package pkCmds

import (
	"fmt"
	"os"
	"rawrpk/src/pkParse"
	"strings"
)

func CmdsHandler(sourceIdentifier string) {
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
		pkParse.ParseGit(repository)
	case "url":
		fmt.Printf("Fetching from URL: %s\n", identifier)
	default:
		fmt.Printf("Unsupported source: %s\n", source)
		os.Exit(1)
	}
}
