package cmd

import "strings"

func trimXML(s string) string {
	s = strings.Replace(s, "<string>", "", -1)
	ts := strings.TrimSpace((strings.Replace(s, "</string>", "", -1)))
	return ts
}
