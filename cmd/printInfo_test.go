package cmd

import "testing"

func Test_printInfo(t *testing.T) {
	type args struct {
		file string
		path string
	}
	tests := []struct {
		name string
		args args
	}{
		{"", args{file: "this", path: "that"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printInfo(tt.args.file, tt.args.path)
		})
	}
}
