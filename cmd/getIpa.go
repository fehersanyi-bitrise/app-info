package cmd

import (
	"io/ioutil"
	"strings"

	"github.com/bitrise-io/go-utils/log"
)

func getIpa(file, path string) (map[string]string, error) {
	appInfo := make(map[string]string)
	zip := strings.Split(file, ".") // path may contain multiple dots ./app.xcarchive/info.plist, instead: strings.TrimSuffix(file, filepath.Ext(file))
	zip[1] = "zip"
	newFile := strings.Join(zip, ".")

	copy := command{Command: "cp", Flag: path + "/" + file, Path: path + "/" + newFile} // command model is a bit strict and over abstraction
	unzip := command{Command: "unzip", Flag: "-oa", Path: path + "/" + newFile, extraFlag: "-d", outPath: path}
	removeZip := command{Command: "rm", Flag: "-f", Path: path + "/" + newFile} // declare var close to its usage
	removePayload := command{Command: "rm", Flag: "-rf", Path: path + "/" + "Payload/"}

	if err := do(copy); err != nil {
		return appInfo, err
	}
	if err := do(unzip); err != nil {
		return appInfo, err
	}

	infoPlist := path + "/Payload/" + zip[0] + ".app/Info.plist" // filepath.Join
	infoXML, err := ioutil.ReadFile(infoPlist)
	if err != nil {
		log.Warnf("here")
		return appInfo, err
	}

	infoArray := strings.Split(string(infoXML), "\n") // use plutil -convert xml1 Payload/Runner.app/Info.plist, to convert to standard plsit format, bit better to use a plist parser package like "howett.net/plist"
	for i := 0; i < len(infoArray); i++ {
		if strings.Contains(infoArray[i], "BundleIdentifier") {
			appInfo["Bundle ID"] = trimXML(infoArray[i+1])
		}
		if strings.Contains(infoArray[i], "BundleShortVersionString") {
			appInfo["Version Number"] = trimXML(infoArray[i+1])
		}
		if strings.Contains(infoArray[i], "CFBundleVersion") {
			appInfo["Build Number"] = trimXML(infoArray[i+1])
		}
		if strings.Contains(infoArray[i], "AppIcon-260") {
			appInfo["App Icon"] = trimXML(infoArray[i])
		}
	}
	if err := do(removeZip); err != nil {
		return appInfo, err
	}
	if err := do(removePayload); err != nil {
		return appInfo, err
	}
	return appInfo, err
}
