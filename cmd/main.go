package main

import (
	"flag"
	"fmt"
	"os"
	"rawrpk/src/pkCmds"
)

// Outside of main are 3 packages: pkCmds, pkParse, and pkSystem.
// pkCmds handles the get command.
// pkParse parsing for both the get command & the github repository.
// pkSystem handles the rawrpk procedures.

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
		pkCmds.CmdsHandler(sourceIdentifier)
	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}
