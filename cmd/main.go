package main

import (
	"rawrpk/internal/cli"
	"rawrpk/internal/common"
	"rawrpk/internal/gitparse"
)

func main() {
	repository, i, err := cli.CLIparse()
	if err != nil {
		panic(err)
	} else if i == common.Github {
		gitparse.ParseGit(repository)
	}
}
