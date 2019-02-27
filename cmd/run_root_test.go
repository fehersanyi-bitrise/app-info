package cmd

import "testing"

func Test_run_root(t *testing.T) {

	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{"", []string{"../test.apk"}, false},
		{"", []string{"../semmi.apk"}, true},
		{"", []string{"../Grability.ipa"}, false},
		{"", []string{"../semmi.ipa"}, true},
		{"", []string{}, true},
		{"", []string{"fdsad"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := runRoot(tt.args); (err != nil) != tt.wantErr {
				t.Errorf("runRoot() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
