package cmd

import "testing"

func Test_printAppInfo(t *testing.T) {
	type args struct {
		appInfo map[string]string
		appType string
	}
	tests := []struct {
		name string
		args args
	}{
		{"", args{appInfo: map[string]string{"packageName": "ez", "versionCode": "1", "versionName": "6.6.6", "icon": "./././"}, appType: "apk"}},
		{"", args{appInfo: map[string]string{"Bundle ID": "ez", "Version Number": "1", "Build Number": "6.6.6", "icon": "./././"}, appType: "ipa"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printAppInfo(tt.args.appInfo, tt.args.appType)
		})
	}
}
