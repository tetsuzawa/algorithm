import os

// eg. x=18532, returns 5
fn nDigits(x int) int {
	cnt := 0
	for x != 0 {
		cnt++
		x /= 10
	}
	return cnt
}

fn main() {
	a := 18532
	cnt := nDigits(a)
	println(cnt)
}
