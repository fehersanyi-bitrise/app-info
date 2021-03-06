package cmd

import (
	"os/exec"
	"strings"
)

func getAPK(path string) (map[string]string, error) {
	command, err := exec.Command("aapt", "dump", "badging", path).Output() // works only if aapt is in PATH, see: https://github.com/bitrise-io/steps-deploy-to-bitrise-io/blob/master/uploaders/apkuploader.go#L69-L82
	appInfo := make(map[string]string)
	if err != nil {
		return appInfo, err
	}
	info := strings.Split(string(command), "\n")
	for i := 0; i < len(info); i++ { // for _, line := range info {}
		if strings.Contains(info[i], "package") {
			appInfo["packageName"] = strings.Split(info[i], "'")[1]
			appInfo["versionCode"] = strings.Split(info[i], "'")[3]
			appInfo["versionName"] = strings.Split(info[i], "'")[5]
		} else if strings.Contains(info[i], "application-icon") {
			appInfo["icon"] = strings.Split(info[i], "'")[1]
		}
	}
	return appInfo, nil
}
