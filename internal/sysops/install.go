package sysops

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"rawrpk/internal/common"
	"strings"
)

/*
This pksystem package is organized to contain the install process of rawrpk.
The download processor for rawrpk.
The environment setup processor for rawrpk.
*/
func Install(pkg *common.PkgData) {
	url := strings.Trim(pkg.DwnSrc, "\"")
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	err = os.MkdirAll(pkg.InstallLoc, os.ModePerm)
	if err != nil {
		panic(err)
	}

	out, err := os.Create(fmt.Sprintf(pkg.InstallLoc+"\\%s.exe", pkg.Title))
	if err != nil {
		panic(err)
	}
	fmt.Println("Creating file:", pkg.Title)
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Installing from link:", url)
}
