package d05_ComplementOfBase10Integer

import (
	"math/bits"
)

func reversedInts(sl []int) []int {
	ret := make([]int, len(sl))
	for i, v := range sl {
		ret[len(sl)-i-1] = v
	}
	return ret
}

func decToBits(dec int) (bitSlice []int) {
	if dec == 0 {
		return []int{0}
	}
	bitSlice = make([]int, 0, bits.Len(uint(dec)))
	for dec > 0 {
		bitSlice = append(bitSlice, dec%2)
		dec /= 2
	}
	bitSlice = reversedInts(bitSlice)
	return
}

func invertBits(bitSlice []int) []int {
	ret := make([]int, len(bitSlice))
	for i, v := range bitSlice {
		if v == 0 {
			ret[i] = 1
		} else {
			ret[i] = 0
		}
	}
	return ret
}

func bitsToDec(bitSlice []int) (dec int) {
	l := len(bitSlice)
	for i, v := range bitSlice {
		dec += v * (1 << (l - i - 1))
	}
	return
}

/*
func bitwiseComplement(N int) int {
	nBits := decToBits(N)
	return bitsToDec(invertBits(nBits))
}
*/

func bitwiseComplement(N int) int {
	if N == 0 {
		return 1
	}
	ans, cur := 0, 1
	for N > 0 {
		if N&1 == 0 {
			ans += cur
		}
		N, cur = N>>1, cur<<1
	}
	return ans
}
