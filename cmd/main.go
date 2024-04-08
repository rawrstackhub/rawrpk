package main

import (
	"rawrpk/internal/cli"
	"rawrpk/internal/common"
	pkgSrc "rawrpk/internal/pkgsources"
	"rawrpk/internal/rawrHndl"
	"rawrpk/internal/sysops"
)

var Pack common.PkgData

func main() {
	switch cli.CliHandle(&Pack) {
	case common.Github:
		pkgSrc.Github(&Pack)
	}

	cmds := rawrHndl.File(&Pack)
	for _, cmd := range cmds {
		switch rawrHndl.ParseLine(&Pack, cmd) {
		case common.Install:
			sysops.Install(&Pack)
		case common.EnvVar:
			sysops.AddPath(&Pack)
		}
	}
}
