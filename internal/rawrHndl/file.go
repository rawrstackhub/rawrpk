package rawrHndl

import (
	"bufio"
	"fmt"
	"net/http"
	"rawrpk/internal/common"
	"strings"
)

func File(pkg *common.PkgData) []string {
	response, err := http.Get(pkg.PkgURL)
	if err != nil {
		panic("Error fetching file: " + err.Error())
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		panic("Error fetching file: " + response.Status)
	}

	scanner := bufio.NewScanner(response.Body)

	var export []string
	for scanner.Scan() {
		line := scanner.Text()
		export = append(export, line)
	}
	if err := scanner.Err(); err != nil {
		panic("Error reading file: " + err.Error())
	}
	fmt.Println("File read successfully, number of lines: ", len(export))
	return export
}

func ParseLine(pkg *common.PkgData, line string) int8 {
	tokens := strings.Fields(line)
	if len(tokens) < 2 {
		panic("Invalid instruction: " + line)
	}

	switch tokens[0] {
	case "INSTALL":
		if len(tokens) != 2 {
			panic("invalid INSTALL format: " + line)
		}
		pkg.DwnSrc = tokens[1]
		return common.Install

	case "ENVVAR":
		if len(tokens) != 4 || tokens[1] != "ADD" {
			panic("invalid ENVVAR format: " + line)
		}
		return common.EnvVar
	default:
		panic("Unknown instruction: " + tokens[0])
	}
}
