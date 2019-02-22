package cmd

import (
	"fmt"

	"github.com/bitrise-io/go-utils/log"
)

func printAppInfo(appInfo map[string]string, appType string) {
	if appType == "apk" {
		fmt.Println()
		log.Printf("Package name: %s", appInfo["packageName"])
		log.Printf("Version Code: %s", appInfo["cersionCode"])
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
