package cmd

import "testing"

func Test_processPath(t *testing.T) {

	tests := []struct {
		name    string
		args    []string
		want    string
		want1   string
		wantErr bool
	}{
		{"", []string{"here/there.apk"}, "there.apk", "here", false},
		{"", []string{"here/there.ipa"}, "there.ipa", "here", false},
		{"", []string{"here/there.zip"}, "there.zip", "here", false},
		{"", []string{}, "", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := processPath(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("processPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("processPath() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("processPath() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
