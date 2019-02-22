package cmd

import (
	"fmt"

	"github.com/bitrise-io/go-utils/log"
)

//IPA ...
func IPA(file, path string, args []string) error {
	log.Infof("Retrieving IPA Info:")
	printInfo(file, path)
	appInfo, err := getIpa(file, path)
	if err != nil {
		return fmt.Errorf("Error retrieving IPA info %s", err)
	}
	printAppInfo(appInfo, "ipa")
	return nil
}
