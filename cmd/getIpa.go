package cmd

import (
	"io/ioutil"
	"os/exec"
	"strings"
)

func getIpa(file, path string) (map[string]string, error) {
	appInfo := make(map[string]string)
	zip := strings.Split(file, ".")
	zip[1] = "zip"
	newFile := strings.Join(zip, ".")
	//make a temporal copy
	_, err := exec.Command("cp", path+"/"+file, path+"/"+newFile).Output()
	if err != nil {
		return appInfo, err
	}
	//create payload
	_, err = exec.Command("unzip", "-oa", path+"/"+newFile, "-d", path).Output()
	if err != nil {

		return appInfo, err
	}
	//read
	infoPlist := path + "/Payload/" + zip[0] + ".app/Info.plist"
	infoXML, err := ioutil.ReadFile(infoPlist)
	if err != nil {
		return appInfo, err
	}
	infoArray := strings.Split(string(infoXML), "\n")
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
	//remove zip
	_, err = exec.Command("rm", "-f", path+"/"+newFile).Output()
	if err != nil {
		return appInfo, err
	}
	//remove payload
	_, err = exec.Command("rm", "-rf", path+"/"+"Payload/").Output()
	if err != nil {
		return appInfo, err
	}
	return appInfo, nil
}
