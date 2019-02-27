// Package cmd ...
// go uses snake_case file names -> run_root.go
package cmd

import (
	"errors"
	"strings"
)

func runRoot(args []string) error {
	file, path, err := processPath(args)
	if err != nil {
		return err
	}
	// validate before parse (processPath)
	if !strings.Contains(args[0], ".apk") && !strings.Contains(args[0], ".ipa") {
		return errors.New("Incorrect input")
	}
	if strings.Contains(args[0], ".apk") {
		if err := APK(file, path, args); err != nil {
			return err
		}
	} else if strings.Contains(args[0], ".ipa") {
		if err := IPA(file, path, args); err != nil {
			return err
		}
	}
	return nil
}
