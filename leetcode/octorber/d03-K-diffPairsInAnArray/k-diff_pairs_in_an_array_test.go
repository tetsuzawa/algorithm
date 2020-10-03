package d03_K_diffPairsInAnArray

import "testing"

func Test_findPairs(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				nums: []int{1,2,4,4,3,3,0,9,2,3},
				k:    3,
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				nums: []int{0,0,0},
				k:    0,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPairs(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}