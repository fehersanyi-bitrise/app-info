package cmd

import (
	"path/filepath"
	"strings"
)

func runRoot(arg string) error {
	path := filepath.Dir(arg)
	file := filepath.Base(arg)
	if strings.Contains(file, ".apk") {
		if err := APK(file, path, arg); err != nil {
			return err
		}
	} else if strings.Contains(file, ".ipa") {
		if err := IPA(arg); err != nil {
			return err
		}
	}
	return nil
}
