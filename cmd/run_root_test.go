package cmd

import "testing"

func Test_run_root(t *testing.T) {

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
			if err := runRoot(tt.args); (err != nil) != tt.wantErr {
				t.Errorf("runRoot() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
