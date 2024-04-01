package pkSystem

import (
	"fmt"

	"golang.org/x/sys/windows/registry"
)

func SetEnvVar(name, value string) error {
	fmt.Println("Setting Environment Variable:", name, value)

	key, err := registry.OpenKey(registry.CURRENT_USER, `Environment`, registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer key.Close()

	if err := key.SetStringValue(name, value); err != nil {
		return err
	}
	fmt.Println("Environment Variable Set:", name, value)
	return nil
}
