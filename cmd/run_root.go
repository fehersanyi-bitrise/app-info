package cmd

import (
	"errors"
	"path/filepath"
	"strings"
)

func runRoot(arg string) error {
	path := filepath.Dir(arg)
	file := filepath.Base(arg)
	if !strings.Contains(arg, ".apk") && !strings.Contains(arg, ".ipa") {
		return errors.New("Incorrect input")
	}
	if strings.Contains(arg, ".apk") {
		if err := APK(file, path, arg); err != nil {
			return err
		}
	} else if strings.Contains(arg, ".ipa") {
		if err := IPA(arg); err != nil {
			return err
		}
	}
	return nil
}
