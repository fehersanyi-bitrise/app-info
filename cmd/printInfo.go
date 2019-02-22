package cmd

import (
	"fmt"

	"github.com/bitrise-io/go-utils/log"
)

func printInfo(file, path string) {
	fmt.Println()
	fmt.Print("APP: ")
	log.Infof("%s", file)
	fmt.Print("PATH: ")
	log.Infof("%s", path)
	fmt.Println()
}
