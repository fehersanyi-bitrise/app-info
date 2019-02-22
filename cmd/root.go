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
				if err := APK(file, path, args); err != nil {
					log.Errorf(err.Error())
				}
			} else if strings.Contains(args[0], ".ipa") {
				if err := IPA(file, path, args); err != nil {
					log.Errorf(err.Error())
				}
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
