package utils

import "testing"

func TestIndexOf(t *testing.T) {
	type args struct {
		find string
		data []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Index 1", args{"Active", []string{"Inactive", "Active"}}, 1},
		{"Index 0", args{"Inactive", []string{"Inactive", "Active"}}, 0},
		{"Not found return -1", args{"inactive", []string{"Inactive", "Active"}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IndexOf(tt.args.find, tt.args.data); got != tt.want {
				t.Errorf("IndexOf() = %v, want %v", got, tt.want)
			}
		})
	}
}
