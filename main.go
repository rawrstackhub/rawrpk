package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

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
	case "url":
		fmt.Printf("Fetching from URL: %s\n", identifier)
	default:
		fmt.Printf("Unsupported source: %s\n", source)
		os.Exit(1)
	}
}
