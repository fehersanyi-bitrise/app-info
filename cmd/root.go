package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/bitrise-io/go-utils/log"
	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "app-info",
	Short: "provides some info on artifacts",
	Long:  `This app provides information on a given APK or IPA file.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			if strings.Contains(args[0], ".apk") || strings.Contains(args[0], ".ipa") {
				if err := RUNApp(args[0]); err != nil {
					log.Errorf(err.Error())
				}
			} else {
				log.Errorf("Incorrect input must be .apk or .ipa file")
			}
		} else {
			log.Errorf("No app provided")
		}
	},
}

//RUNApp ...
func RUNApp(arg string) error {
	file := filepath.Base(arg)
	if strings.Contains(file, ".apk") {
		if err := APK(arg); err != nil {
			return err
		}
	} else if strings.Contains(file, ".ipa") {
		if err := IPA(arg); err != nil {
			return err
		}
	}
	return nil
}

//Execute will run the command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
