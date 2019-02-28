package cmd

import (
	"fmt"
	"os/exec"

	"github.com/bitrise-io/go-utils/log"
)

func do(c command) error {
	if c.extraFlag != "" {
		_, err := exec.Command(c.Command, c.Flag, c.Path, c.extraFlag, c.outPath).Output()
		if err != nil {
			return err
		}
	} else {
		_, err := exec.Command(c.Command, c.Flag, c.Path).Output()
		if err != nil {
			return err
		}
	}
	return nil
}

func printAppInfo(appInfo map[string]string, appType string) {
	if appType == "apk" {
		fmt.Println()
		log.Printf("Package name: %s", appInfo["packageName"])
		log.Printf("Version Code: %s", appInfo["versionCode"])
		log.Printf("Version Name: %s", appInfo["versionName"])
		log.Printf("Path to icon: %s", appInfo["icon"])
	} else if appType == "ipa" {
		fmt.Println()
		log.Printf("Bundle ID: %s", appInfo["Bundle ID"])
		log.Printf("Version Number %s: ", appInfo["Version Number"])
		log.Printf("Build Number: %s", appInfo["Build Number"])
		log.Printf("Path to icon: %s", appInfo["App Icon"])
	}
}

type command struct {
	Command   string
	Flag      string
	Path      string
	extraFlag string
	outPath   string
}
