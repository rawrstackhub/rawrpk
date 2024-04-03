package sysops

import (
	"fmt"
	"strings"

	"golang.org/x/sys/windows/registry"
)

func AddPath(newPath string) error {
	fmt.Println("Adding to PATH:", newPath)

	key, err := registry.OpenKey(registry.CURRENT_USER, `Environment`, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer key.Close()

	// Read the current value of PATH
	pathValue, _, err := key.GetStringValue("Path")
	if err != nil {
		return err
	}

	if !strings.Contains(strings.ToLower(pathValue), strings.ToLower(newPath)) {
		updatedPath := pathValue + ";" + newPath
		if err := key.SetStringValue("Path", updatedPath); err != nil {
			return err
		}
		fmt.Println("PATH updated successfully")
	} else {
		fmt.Println("The path is already in PATH; no changes made.")
	}

	return nil
}
