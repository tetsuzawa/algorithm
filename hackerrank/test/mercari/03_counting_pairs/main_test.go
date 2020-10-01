package main

import "testing"

func Test_countPairs(t *testing.T) {
	type args struct {
		numbers []int32
		k       int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "0",
			args: args{numbers: []int32{1, 1, 2, 2, 3, 3}, k: 1},
			want: 2,
		},
		{
			name: "1",
			args: args{numbers: []int32{1, 2, 3, 4, 5, 6}, k: 2},
			want: 4,
		},
		{
			name: "2",
			args: args{numbers: []int32{1, 2, 5, 6, 9, 10}, k: 2},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPairs(tt.args.numbers, tt.args.k); got != tt.want {
				t.Errorf("countPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
