// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/bitrise-io/go-utils/log"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "app-info",
	Short: "A brief description of your application",
	Long:  `This app provides information on a given APK or IPA file.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		file, path, err := processPath(args)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print("APP: ")
		log.Infof("%s", file)
		fmt.Print("PATH: ")
		log.Infof("%s", path)
		fmt.Println()
		getAPK(args[0])
	},
}

func getAPK(path string) {
	runCmd := []string{"aapt", "dump", "badging"}
	log.Printf("Running command: ")
	log.Donef("%s %s", strings.Join(runCmd, " "), path)
	command, err := exec.Command("aapt", "dump", "badging", path).Output()
	if err != nil {
		fmt.Printf("failed to run command %s", err)
	}

	info := strings.Split(string(command), "\n")

	appInfo := make(map[string]string)
	for i := 0; i < len(info); i++ {
		if strings.Contains(info[i], "package") {
			appInfo["packageName"] = strings.Split(info[i], "'")[1]
			appInfo["cersionCode"] = strings.Split(info[i], "'")[3]
			appInfo["versionName"] = strings.Split(info[i], "'")[5]
		} else if strings.Contains(info[i], "application-icon") {
			appInfo["icon"] = strings.Split(info[i], "'")[1]
		}
	}

	log.Printf("Package name: %s", appInfo["packageName"])
	log.Printf("Version Code: %s", appInfo["cersionCode"])
	log.Printf("Version Name: %s", appInfo["versionName"])
	log.Printf("Path to icon: %s", appInfo["icon"])
}

func processPath(args []string) (string, string, error) {
	if len(args) > 0 {
		path := args[0]
		apkPath := strings.Split(path, "/")
		fileName := apkPath[len(apkPath)-1]
		pathToUse := strings.Join(apkPath[:len(apkPath)-1], "/")
		return fileName, pathToUse, nil
	}
	return "", "", fmt.Errorf("please provide a path to an apk or an ipa file")
}

//Execute will run the command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.app-info.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".app-info" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".app-info")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
