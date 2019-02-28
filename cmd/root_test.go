package cmd

import "testing"

func TestExecute(t *testing.T) {
	tests := []struct {
		name string
	}{
		{""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Execute()
		})
	}
}

func Test_RUNApp(t *testing.T) {

	tests := []struct {
		name    string
		args    string
		wantErr bool
	}{
		{"", "../test.apk", false},
		{"", "../semmi.apk", true},
		{"", "../Grability.ipa", false},
		{"", "../semmi.ipa", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RUNApp(tt.args); (err != nil) != tt.wantErr {
				t.Errorf("runRoot() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
