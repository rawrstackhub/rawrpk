package main

import (
	"flag"
	"fmt"
	"os"
	"rawrpk/src/pkCmds"
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
		pkCmds.CmdsHandler(sourceIdentifier)
	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}
