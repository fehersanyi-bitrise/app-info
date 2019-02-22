package cmd

import (
	"fmt"

	"github.com/bitrise-io/go-utils/log"
)

//APK ...
func APK(file, path string, args []string) error {
	log.Infof("Retrieving APK Info:")
	printInfo(file, path)
	appInfo, err := getAPK(args[0])
	if err != nil {
		return fmt.Errorf("Failed to retrieve APK info: %s", err)
	}
	printAppInfo(appInfo, "apk")
	return nil
}
