package cmd

import (
	"errors"
	"path/filepath"
	"strings"
)

func runRoot(args []string) error {
	path := filepath.Dir(args[0])
	file := filepath.Base(args[0])
	if !strings.Contains(args[0], ".apk") && !strings.Contains(args[0], ".ipa") {
		return errors.New("Incorrect input")
	}
	if strings.Contains(args[0], ".apk") {
		if err := APK(file, path, args); err != nil {
			return err
		}
	} else if strings.Contains(args[0], ".ipa") {
		if err := IPA(args); err != nil {
			return err
		}
	}
	return nil
}
