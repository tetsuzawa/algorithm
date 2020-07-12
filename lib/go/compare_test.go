package lib

import "testing"

func Test_abs(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "minus",
			args: args{a: -4},
			want: 4,
		},
		{
			name: "zero",
			args: args{a: 0},
			want: 0,
		},
		{
			name: "plus",
			args: args{a: 4},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := abs(tt.args.a); got != tt.want {
				t.Errorf("abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_equals(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "correct",
			args: args{a: []int{1, 2, 3}, b: []int{1, 2, 3}},
			want: true,
		},
		{
			name: "incorrect",
			args: args{a: []int{1, 2, 3}, b: []int{1, 3, 2}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equals(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("equals() = %v, want %v", got, tt.want)
			}
		})
	}
}
