package cmd

import "testing"

func Test_do(t *testing.T) {
	type args struct {
		c command
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"", args{c: command{Command: "echo", Flag: "Hello"}}, false},
		{"", args{c: command{Command: "", Flag: "Hello"}}, true},
		{"", args{c: command{Command: "echo", Flag: "Hello", extraFlag: "there"}}, false},
		{"", args{c: command{Command: "ls", Flag: "Hello", extraFlag: "there"}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := do(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("do() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
