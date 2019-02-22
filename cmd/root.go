package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/bitrise-io/go-utils/log"
	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "app-info",
	Short: "A brief description of your application",
	Long:  `This app provides information on a given APK or IPA file.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			file, path, err := processPath(args)
			if err != nil {
				log.Errorf(err.Error())
			}
			if strings.Contains(args[0], ".apk") {
				log.Infof("Retrieving APK Info:")
				fmt.Println()
				fmt.Print("APP: ")
				log.Infof("%s", file)
				fmt.Print("PATH: ")
				log.Infof("%s", path)
				fmt.Println()
				appInfo, err := getAPK(args[0])
				if err != nil {
					log.Errorf("Failed to retrieve APK info: %s", err)
				}
				log.Printf("Package name: %s", appInfo["packageName"])
				log.Printf("Version Code: %s", appInfo["cersionCode"])
				log.Printf("Version Name: %s", appInfo["versionName"])
				log.Printf("Path to icon: %s", appInfo["icon"])
			} else if strings.Contains(args[0], ".ipa") {
				log.Infof("Retrieving IPA Info:")
				fmt.Println()
				fmt.Print("APP: ")
				log.Infof("%s", file)
				fmt.Print("PATH: ")
				log.Infof("%s", path)
				fmt.Println()

				appInfo, err := getIpa(file, path)
				if err != nil {
					log.Errorf("Error retrieving IPA info %s", err)
					return
				}
				fmt.Println()
				log.Printf("Bundle ID: %s", appInfo["Bundle ID"])
				log.Printf("Version Number %s: ", appInfo["Version Number"])
				log.Printf("Build Number: %s", appInfo["Build Number"])
				log.Printf("Path to icon: %s", appInfo["App Icon"])
			} else {
				log.Errorf("Invalid argument: %s", args[0])
			}
		} else {
			log.Errorf("No app provided")
		}
	},
}

//Execute will run the command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
