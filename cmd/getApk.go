package cmd

import (
	"fmt"
	"os/exec"
	"strings"
)

func getAPK(path string) (map[string]string, error) {
	command, err := exec.Command("aapt", "dump", "badging", path).Output()
	if err != nil {
		fmt.Printf("failed to run command %s", err)
	}
	info := strings.Split(string(command), "\n")
	appInfo := make(map[string]string)
	for i := 0; i < len(info); i++ {
		if strings.Contains(info[i], "package") {
			appInfo["packageName"] = strings.Split(info[i], "'")[1]
			appInfo["cersionCode"] = strings.Split(info[i], "'")[3]
			appInfo["versionName"] = strings.Split(info[i], "'")[5]
		} else if strings.Contains(info[i], "application-icon") {
			appInfo["icon"] = strings.Split(info[i], "'")[1]
		}
	}
	return appInfo, nil
}
