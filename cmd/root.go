package cmd

import (
	"fmt"
	"os"

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
			if err := runRoot(args); err != nil {
				log.Errorf(err.Error())
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
