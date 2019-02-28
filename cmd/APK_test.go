package cmd

import "testing"

func TestAPK(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		wantErr bool
	}{
		{"1", "file.apk", true},
		{"2", "", true},
		{"3", "file.apk", true},
		{"4", "../test.apk", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := APK(tt.args); (err != nil) != tt.wantErr {
				t.Errorf("APK() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
