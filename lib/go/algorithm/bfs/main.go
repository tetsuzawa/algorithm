package main

import (
	"fmt"
)

type Graph [][]int

func NewGraph(n int) Graph {
	g := make(Graph, n, n)
	//for i := range g {
	//	g[i] = make([]int, n, n)
	//}
	return g
}

// O(N+M)
func main() {
	// 頂点数と変数
	var n, m int
	fmt.Scan(&n)
	fmt.Scan(&m)

	var a, b int
	// グラフ入力受け取り
	g := NewGraph(n)
	for i := 0; i < m; i++ {
		fmt.Scan(&a)
		fmt.Scan(&b)
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	// BFSのためのデータ構造
	dist := make([]int, n, n)
	for i := range dist {
		dist[i] = -1
	}
	// queue
	que := make([]int, 0)

	// 初期条件 (頂点 0 を初期ノードとする)
	dist[0] = 0
	que = append(que, 0) // 0 を橙色頂点にする

	// BFS 開始 (キューが空になるまで探索を行う)
	for len(que) != 0 {
		var v int
		v, que = que[0], que[1:] // キューから先頭頂点を取り出す

		// v から辿れる頂点をすべて調べる
		for _, nv := range g[v] {
			if dist[nv] != -1 {
				continue
			}
			// 新たな白色頂点 nv について距離情報を更新してキューに追加する
			dist[nv] = dist[v] + 1
			que = append(que, nv)
		}
	}

	// 結果出力 (各頂点の頂点 0 からの距離を見る)
	for v := 0; v < n; v++ {
		fmt.Printf("%v: %v\n", v, dist[v])
	}

}
