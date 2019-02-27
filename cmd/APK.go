package cmd

import (
	"errors"
	"fmt"

	"github.com/bitrise-io/go-utils/log"
)

//APK ...
func APK(file, path string, args []string) error {
	log.Infof("Retrieving APK Info:")
	// file and path is redundant if you work with args as well
	printInfo(file, path)
	if len(args) > 0 {
		appInfo, err := getAPK(args[0])
		if err != nil {
			return fmt.Errorf("Failed to retrieve APK info: %s", err)
		}
		printAppInfo(appInfo, "apk")
		return nil
	}
	return errors.New("Index out of bounds")
}
