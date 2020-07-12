package main

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "0",
			args: args{s: "abcba"},
			want: true,
		},
		{
			name: "1",
			args: args{s: "abccba"},
			want: true,
		},
		{
			name: "2",
			args: args{s: "abcdba"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBreakPalindrome(t *testing.T) {
	type args struct {
		palindromeStr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "0",
			args: args{palindromeStr:"bab"},
			want: "aab",
		},
		{
			name: "1",
			args: args{palindromeStr:"aaa"},
			want: "IMPOSSIBLE",
		},
		{
			name: "2",
			args: args{palindromeStr:"acca"},
			want: "aaca",
		},
		{
			name: "3",
			args: args{palindromeStr:"aczibizca"},
			want: "aazibizca",
		},
		{
			name: "4",
			args: args{palindromeStr:"zzzz"},
			want: "azzz",
		},
		{
			name: "5",
			args: args{palindromeStr:"aba"},
			want: "IMPOSSIBLE",
		},
		{
			name: "6",
			args: args{palindromeStr:"zbzbz"},
			want: "abzbz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BreakPalindrome(tt.args.palindromeStr); got != tt.want {
				t.Errorf("BreakPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}