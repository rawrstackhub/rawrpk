package cli

import (
	"flag"
	"fmt"
	"os"
	"rawrpk/internal/common"
	"strings"
)

func CliHandle(pck *common.PkgData) int8 {
	flag.Parse()

	if len(flag.Args()) < 2 {
		fmt.Println("Usage: rawrpk get <source>.<repository>")
		os.Exit(1)
	}

	command := flag.Arg(0)
	subCmd := flag.Arg(1)

	switch command {
	case "get":
		data := sourceSplit(subCmd)
		pck.Title = data[2]   //EX: lsf
		pck.Source = data[:2] //EX: github / rawrstackhub

		return common.Github
	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
	return 0
}

func sourceSplit(source string) [3]string {
	var export [3]string

	mainSRC := strings.SplitN(source, ".", 2)
	if len(mainSRC) != 2 {
		fmt.Println("Invalid source format")
		os.Exit(1)
	}
	export[0] = mainSRC[0] //EX: github

	subSRC := strings.Split(mainSRC[1], "/")
	if len(subSRC) != 2 {
		fmt.Println("Invalid source format")
		os.Exit(1)
	}
	export[1] = subSRC[0] //EX: rawrstackhub
	export[2] = subSRC[1] //EX: lsf

	return export
}
