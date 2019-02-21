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
	"io/ioutil"
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
		if len(args) > 0 {
			file, path := logInfo(args)
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
				logInfo(args)
				//TODO
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

func getIpa(file, path string) (map[string]string, error) {
	appInfo := make(map[string]string)
	zip := strings.Split(file, ".")
	zip[1] = "zip"
	newFile := strings.Join(zip, ".")
	//make a temporal copy
	_, err := exec.Command("cp", path+"/"+file, path+"/"+newFile).Output()
	if err != nil {
		log.Warnf("zipp")
		return appInfo, err
	}
	//create payload
	_, err = exec.Command("unzip", "-oa", path+"/"+newFile, "-d", path).Output()
	if err != nil {
		log.Printf(path + "/" + newFile)
		log.Warnf("payload")
		return appInfo, err
	}
	//cat
	infoPlist := path + "/Payload/" + zip[0] + ".app/Info.plist"
	infoXML, err := ioutil.ReadFile(infoPlist)
	if err != nil {
		log.Printf(infoPlist)
		log.Warnf("cat")
		return appInfo, err
	}
	infoArray := strings.Split(string(infoXML), "\n")
	for i := 0; i < len(infoArray); i++ {
		if strings.Contains(infoArray[i], "BundleIdentifier") {
			first := strings.Replace(infoArray[i+1], "<string>", "", -1)
			appInfo["Bundle ID"] = strings.TrimSpace(strings.Replace(first, "</string>", "", -1))
		}
		if strings.Contains(infoArray[i], "BundleShortVersionString") {
			first := strings.Replace(infoArray[i+1], "<string>", "", -1)
			appInfo["Version Number"] = strings.TrimSpace(strings.Replace(strings.Replace(first, "</string>", "", -1), ":", "", -1))
		}
		if strings.Contains(infoArray[i], "CFBundleVersion") {
			first := strings.Replace(infoArray[i+1], "<string>", "", -1)
			appInfo["Build Number"] = strings.TrimSpace(strings.Replace(first, "</string>", "", -1))
		}
		if strings.Contains(infoArray[i], "AppIcon-260") {
			first := strings.Replace(infoArray[i], "<string>", "", -1)
			appInfo["App Icon"] = strings.TrimSpace(strings.Replace(first, "</string>", "", -1))
		}
	}
	//remove zip
	_, err = exec.Command("rm", "-f", path+"/"+newFile).Output()
	if err != nil {
		return appInfo, err
	}
	//remove payload
	_, err = exec.Command("rm", "-rf", path+"/"+"Payload/").Output()
	if err != nil {
		return appInfo, err
	}
	return appInfo, nil
}

func logInfo(args []string) (file, path string) {
	file, path, err := processPath(args)
	if err != nil {
		fmt.Println(err)
	}
	return file, path
}

func getAPK(path string) (map[string]string, error) {
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
	return appInfo, nil
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
