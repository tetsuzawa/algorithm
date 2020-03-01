package main

import "testing"

func Test_maxDigit(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "a",
			args: args{18523},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDigit(tt.args.x); got != tt.want {
				t.Errorf("maxDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}