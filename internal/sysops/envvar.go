package sysops

import (
	"fmt"
	"rawrpk/internal/common"
	"strings"

	"golang.org/x/sys/windows/registry"
)

func AddPath(pkg *common.PkgData) {
	key, err := registry.OpenKey(registry.CURRENT_USER, `Environment`, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		panic(err)
	}
	defer key.Close()

	// Read the current value of PATH
	pathValue, _, err := key.GetStringValue("Path")
	if err != nil {
		panic(err)
	}

	if !strings.Contains(strings.ToLower(pathValue), strings.ToLower(pkg.InstallLoc)) {
		updatedPath := pathValue + ";" + pkg.InstallLoc
		if err := key.SetStringValue("Path", updatedPath); err != nil {
			panic(err)
		}
		fmt.Println("PATH updated successfully")
	} else {
		fmt.Println("The path is already in PATH; no changes made.")
	}
}
