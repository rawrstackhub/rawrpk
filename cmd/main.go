package main

import (
	"rawrpk/internal/cli"
	"rawrpk/internal/common"
	"rawrpk/internal/gitparse"
)

var Pack common.PkgData

func main() {
	switch command() {
	case common.Github:
		gitparse.ParseGit()
	}
}

func command() int8 {
	i, err := cli.CLIparse(&Pack)
	if err != nil {
		panic(err)
	}
	return i
}
