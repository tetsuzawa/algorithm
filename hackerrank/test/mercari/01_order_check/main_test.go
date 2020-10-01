package main

import "testing"

func Test_countStudents(t *testing.T) {
	type args struct {
		height []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "0",
			args: args{height: []int32{1, 1, 3, 4, 1}},
			want: 3,
		},
		{
			name: "1",
			args: args{height: []int32{1, 2, 3, 4}},
			want: 0,
		},
		{
			name: "2",
			args: args{height: []int32{4, 3, 2, 1}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countStudents(tt.args.height); got != tt.want {
				t.Errorf("countStudents() = %v, want %v", got, tt.want)
			}
		})
	}
}
