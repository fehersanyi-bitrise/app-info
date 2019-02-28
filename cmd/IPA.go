package cmd

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/bitrise-io/go-utils/log"
)

//IPA ...
func IPA(arg string) error {
	log.Infof("Retrieving IPA Info:")
	appInfo, err := getIpa(arg)
	if err != nil {
		return fmt.Errorf("Error retrieving IPA info %s", err)
	}
	printAppInfo(appInfo, "ipa")
	return nil
}

func getIpa(arg string) (map[string]string, error) {
	appInfo := make(map[string]string)
	file := filepath.Base(arg)
	path := filepath.Dir(arg)
	zip := strings.TrimSuffix(file, filepath.Ext(file))
	newFile := zip + ".zip"

	if err := do("cp", arg, filepath.Join(path, newFile)); err != nil {
		return appInfo, err
	}

	if err := do("unzip", "-oa", filepath.Join(path, newFile), "-d", path); err != nil {
		return appInfo, err
	}

	infoPlist := filepath.Join(path, "Payload", zip+".app", "Info.plist")

	infoXML, err := exec.Command("plutil", "-p", infoPlist).Output()
	if err != nil {
		return appInfo, err
	}

	infoArray := strings.Split(string(infoXML), "\n")
	for _, line := range infoArray {
		if strings.Contains(line, "BundleIdentifier") {
			appInfo["Bundle ID"] = trimPlist(line)
		}
		if strings.Contains(line, "BundleShortVersionString") {
			appInfo["Version Number"] = trimPlist(line)
		}
		if strings.Contains(line, "CFBundleVersion") {
			appInfo["Build Number"] = trimPlist(line)
		}
		if strings.Contains(line, "AppIcon-260") {
			appInfo["App Icon"] = trimPlist(line)
		}
	}

	if err := do("rm", "-fr", filepath.Join(path, newFile)); err != nil {
		return appInfo, err
	}
	if err := do("rm", "-rf", filepath.Join(path, "Payload")); err != nil {
		return appInfo, err
	}
	return appInfo, err
}

func trimPlist(s string) string {
	r := strings.Split(s, " ")
	return r[len(r)-1]
}
