package pkSystem

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

/*
This pksystem package is organized to contain the install process of rawrpk.
The download processor for rawrpk.
The environment setup processor for rawrpk.
*/

// name is for the subdirectory app folder.
func Install(name, url string) error {
	userDir, err := os.UserHomeDir()
	installDir := userDir + "\\rawrpk\\" + name
	if err != nil {
		return err
	}

	url = strings.Trim(url, "\"")
	fmt.Println("URL is ", url)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = os.MkdirAll(installDir, os.ModePerm)
	if err != nil {
		return err
	} else {
		fmt.Println("Creating directory:", installDir)
	}

	out, err := os.Create(fmt.Sprintf(installDir+"\\%s.exe", name))
	if err != nil {
		return err
	} else {
		fmt.Println("Creating file:", name)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("Installing from link:", url)
	return nil
}
