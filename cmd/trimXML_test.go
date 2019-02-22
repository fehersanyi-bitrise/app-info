package cmd

import "testing"

func Test_trimXML(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{"", "<string>something</string>", "something"},
		{"", "something", "something"},
		{"", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trimXML(tt.args); got != tt.want {
				t.Errorf("trimXML() = %v, want %v", got, tt.want)
			}
		})
	}
}
