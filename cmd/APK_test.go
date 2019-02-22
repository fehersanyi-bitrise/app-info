package cmd

import "testing"

func TestAPK(t *testing.T) {
	type args struct {
		file string
		path string
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"1", args{"this", "is", []string{"file.apk"}}, true},
		{"2", args{"", "", []string{}}, true},
		{"3", args{"", "", []string{"file.apk"}}, true},
		{"4", args{"test.apk", "..", []string{"../test.apk"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := APK(tt.args.file, tt.args.path, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("APK() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
