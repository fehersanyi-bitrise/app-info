package cmd

import (
	"fmt"
	"strings"
)

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
