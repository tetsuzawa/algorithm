package combinationsum

import "reflect"

var g [][]int

func hasSameElems(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	ma := make(map[int]int)
	mb := make(map[int]int)
	for i := range a {
		ma[a[i]]++
		mb[b[i]]++
	}
	return reflect.DeepEqual(ma, mb)
}

func dfs(candidates []int, target int, combs []int) {
	for _, v := range candidates {
		if v < target {
			dfs(candidates, target-v, append(combs, v))
		} else if v == target {
			for _, sl := range g{
				if hasSameElems(sl, append(combs, v)){
					return
				}
			}
			g = append(g, append(combs, v))
		}
	}
}

func combinationSum(candidates []int, target int) [][]int {
	g = make([][]int, 0)
	dfs(candidates, target, nil)
	return g
}
