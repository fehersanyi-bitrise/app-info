package cmd

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/bitrise-io/go-utils/log"
)

//APK ...
func APK(arg string) error {
	if err := checkDependency(); err != nil {
		log.Errorf("%s", err)
	}
	log.Infof("Retrieving APK Info:")
	if arg != "" {
		appInfo, err := getAPK(arg)
		if err != nil {
			return fmt.Errorf("Failed to retrieve APK info: %s", err)
		}
		printAppInfo(appInfo, "apk")
		return nil
	}
	return errors.New("Index out of bounds")
}

func getAPK(path string) (map[string]string, error) {
	command, err := exec.Command("aapt", "dump", "badging", path).Output()
	appInfo := make(map[string]string)
	if err != nil {
		return appInfo, err
	}
	info := strings.Split(string(command), "\n")
	for _, line := range info {
		if strings.Contains(line, "package") {
			appInfo["packageName"] = strings.Split(line, "'")[1]
			appInfo["versionCode"] = strings.Split(line, "'")[3]
			appInfo["versionName"] = strings.Split(line, "'")[5]
		} else if strings.Contains(line, "application-icon") {
			appInfo["icon"] = strings.Split(line, "'")[1]
		}
	}
	return appInfo, nil
}

func checkDependency() error {
	log.Infof("Checking dependencies")
	d, err := exec.Command("mdfind", "-name", "aapt").Output()
	if err != nil {
		return err
	}
	if string(d) != "" {
		log.Successf("aapt packege found in PATH")
		fmt.Println()
		return nil
	}
	log.Warnf("aapt not installed in PATH")
	log.Printf("installing android tools")
	_, err = exec.Command("brew", "cask", "install", "android-sdk").Output()
	if err != nil {
		return err
	}
	log.Successf("android tools installed successfully")
	return nil
}
