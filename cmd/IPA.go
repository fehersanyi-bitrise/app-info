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

	copy := command{Command: "cp", Flag: path + "/" + file, Path: path + "/" + newFile} // command model is a bit strict and over abstraction
	unzip := command{Command: "unzip", Flag: "-oa", Path: path + "/" + newFile, extraFlag: "-d", outPath: path}

	if err := do(copy); err != nil {
		return appInfo, err
	}
	if err := do(unzip); err != nil {
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

	removeZip := command{Command: "rm", Flag: "-f", Path: path + "/" + newFile}
	removePayload := command{Command: "rm", Flag: "-rf", Path: path + "/" + "Payload/"}

	if err := do(removeZip); err != nil {
		return appInfo, err
	}
	if err := do(removePayload); err != nil {
		return appInfo, err
	}
	return appInfo, err
}

func trimPlist(s string) string {
	r := strings.Split(s, " ")
	return r[len(r)-1]
}
