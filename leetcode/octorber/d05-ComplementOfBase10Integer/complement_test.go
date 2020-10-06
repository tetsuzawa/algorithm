package d05_ComplementOfBase10Integer

import (
	"reflect"
	"testing"
)

func Test_decToBin(t *testing.T) {
	type args struct {
		dec int
	}
	tests := []struct {
		name         string
		args         args
		wantBitSlice []int
	}{
		{
			name: "",
			args: args{
				dec: 5,
			},
			wantBitSlice: []int{1, 0, 1},
		},
		{
			name: "",
			args: args{
				dec: 7,
			},
			wantBitSlice: []int{1, 1, 1},
		},
		{
			name: "",
			args: args{
				dec: 8,
			},
			wantBitSlice: []int{1, 0, 0, 0},
		},
		{
			name: "",
			args: args{
				dec: 0,
			},
			wantBitSlice: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotBitSlice := decToBits(tt.args.dec); !reflect.DeepEqual(gotBitSlice, tt.wantBitSlice) {
				t.Errorf("decToBits() = %v, want %v", gotBitSlice, tt.wantBitSlice)
			}
		})
	}
}

func Test_reversedInts(t *testing.T) {
	type args struct {
		sl []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				sl: []int{1, 2, 3},
			},
			want: []int{3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reversedInts(tt.args.sl); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reversedInts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bitwiseComplement(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				N: 5,
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				N: 7,
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				N: 10,
			},
			want: 5,
		},
		{
			name: "",
			args: args{
				N: 0,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bitwiseComplement(tt.args.N); got != tt.want {
				t.Errorf("bitwiseComplement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bitsToDec(t *testing.T) {
	type args struct {
		bitSlice []int
	}
	tests := []struct {
		name    string
		args    args
		wantDec int
	}{
		{
			name: "",
			args: args{
				bitSlice: []int{1, 0, 1},
			},
			wantDec: 5,
		},
		{
			name: "",
			args: args{
				bitSlice: []int{1, 1, 1},
			},
			wantDec: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDec := bitsToDec(tt.args.bitSlice); gotDec != tt.wantDec {
				t.Errorf("bitsToDec() = %v, want %v", gotDec, tt.wantDec)
			}
		})
	}
}

