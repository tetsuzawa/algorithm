package lib

import (
	"reflect"
	"testing"
)

func TestPermute(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "[1,2,3]",
			args: args{nums: []int{1, 2, 3}},
			want: [][]int{
				{1, 2, 3,},
				{2, 1, 3,},
				{3, 1, 2,},
				{1, 3, 2,},
				{2, 3, 1,},
				{3, 2, 1,},
			},
		},
		{
			name: "[1,2,3,4]",
			args: args{nums: []int{1, 2, 3, 4}},
			want: [][]int{
				{1, 2, 3, 4,},
				{2, 1, 3, 4,},
				{3, 1, 2, 4,},
				{1, 3, 2, 4,},
				{2, 3, 1, 4,},
				{3, 2, 1, 4,},
				{3, 2, 4, 1,},
				{2, 3, 4, 1,},
				{4, 3, 2, 1,},
				{3, 4, 2, 1,},
				{2, 4, 3, 1,},
				{4, 2, 3, 1,},
				{4, 1, 3, 2,},
				{1, 4, 3, 2,},
				{3, 4, 1, 2,},
				{4, 3, 1, 2,},
				{1, 3, 4, 2,},
				{3, 1, 4, 2,},
				{2, 1, 4, 3,},
				{1, 2, 4, 3,},
				{4, 2, 1, 3,},
				{2, 4, 1, 3,},
				{1, 4, 2, 3,},
				{4, 1, 2, 3,},
			},
		},
		{
			name: "[4,5,6]",
			args: args{nums: []int{2, 4, 6}},
			want: [][]int{
				{2, 4, 6,},
				{4, 2, 6,},
				{6, 2, 4,},
				{2, 6, 4,},
				{4, 6, 2,},
				{6, 4, 2,},
			},
		},
		{
			name: "[1,2]",
			args: args{nums: []int{1, 2}},
			want: [][]int{
				{1, 2,},
				{2, 1,},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Permute(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Permute() = %v, want %v", got, tt.want)
				//for _, v := range got {
				//	//fmt.Printf("%v: \t%v\n", i, v)
				//	fmt.Print("{")
				//	for _, vv := range v {
				//		//fmt.Printf("%v: \t%v\n", i, v)
				//		fmt.Printf("%v,", vv)
				//	}
				//	fmt.Printf("},\n")
				//}
				//fmt.Printf("\n")
			}
		})
	}
}

func TestSortPermutation(t *testing.T) {
	type args struct {
		nums [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "[1,2,3]",
			args: args{nums: [][]int{
				{1, 2, 3,},
				{2, 1, 3,},
				{3, 1, 2,},
				{1, 3, 2,},
				{2, 3, 1,},
				{3, 2, 1,},
			}},
			want: [][]int{
				{1, 2, 3,},
				{1, 3, 2,},
				{2, 1, 3,},
				{2, 3, 1,},
				{3, 1, 2,},
				{3, 2, 1,},
			},
		},
		{
			name: "[1,2,3,4]",
			args: args{nums: [][]int{
				{1, 2, 3, 4,},
				{2, 1, 3, 4,},
				{3, 1, 2, 4,},
				{1, 3, 2, 4,},
				{2, 3, 1, 4,},
				{3, 2, 1, 4,},
				{3, 2, 4, 1,},
				{2, 3, 4, 1,},
				{4, 3, 2, 1,},
				{3, 4, 2, 1,},
				{2, 4, 3, 1,},
				{4, 2, 3, 1,},
				{4, 1, 3, 2,},
				{1, 4, 3, 2,},
				{3, 4, 1, 2,},
				{4, 3, 1, 2,},
				{1, 3, 4, 2,},
				{3, 1, 4, 2,},
				{2, 1, 4, 3,},
				{1, 2, 4, 3,},
				{4, 2, 1, 3,},
				{2, 4, 1, 3,},
				{1, 4, 2, 3,},
				{4, 1, 2, 3,},
			}},
			want: [][]int{
				{1, 2, 3, 4,},
				{1, 2, 4, 3,},
				{1, 3, 2, 4,},
				{1, 3, 4, 2,},
				{1, 4, 2, 3,},
				{1, 4, 3, 2,},
				{2, 1, 3, 4,},
				{2, 1, 4, 3,},
				{2, 3, 1, 4,},
				{2, 3, 4, 1,},
				{2, 4, 1, 3,},
				{2, 4, 3, 1,},
				{3, 1, 2, 4,},
				{3, 1, 4, 2,},
				{3, 2, 1, 4,},
				{3, 2, 4, 1,},
				{3, 4, 1, 2,},
				{3, 4, 2, 1,},
				{4, 1, 2, 3,},
				{4, 1, 3, 2,},
				{4, 2, 1, 3,},
				{4, 2, 3, 1,},
				{4, 3, 1, 2,},
				{4, 3, 2, 1,},
			},
		},
		{
			name: "[4,5,6]",
			args: args{nums: [][]int{
				{2, 4, 6,},
				{4, 2, 6,},
				{6, 2, 4,},
				{2, 6, 4,},
				{4, 6, 2,},
				{6, 4, 2,},
			}},
			want: [][]int{
				{2, 4, 6,},
				{2, 6, 4,},
				{4, 2, 6,},
				{4, 6, 2,},
				{6, 2, 4,},
				{6, 4, 2,},
			},
		},
		{
			name: "[1,2]",
			args: args{nums: [][]int{
				{1, 2,},
				{2, 1,},
			}},
			want: [][]int{
				{1, 2,},
				{2, 1,},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortPermutation(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortPermutation() = %v, want %v", got, tt.want)
				//for _, v := range got {
				//	//fmt.Printf("%v: \t%v\n", i, v)
				//	fmt.Print("{")
				//	for _, vv := range v {
				//		//fmt.Printf("%v: \t%v\n", i, v)
				//		fmt.Printf("%v,", vv)
				//	}
				//	fmt.Printf("},\n")
				//}
				//fmt.Printf("\n")
			}
		})
	}
}

func BenchmarkPermute(b *testing.B) {
	num := 5
	nums := make([]int, num, num)
	for i := 0; i < num; i++ {
		nums[i] = num + 1
	}
	for i := 0; i < b.N; i++ {
		Permute(nums)
	}
}

func BenchmarkSortedPermute(b *testing.B) {
	num := 5
	nums := make([]int, num, num)
	for i := 0; i < num; i++ {
		nums[i] = num + 1
	}
	for i := 0; i < b.N; i++ {
		SortedPermute(nums)
	}
}
