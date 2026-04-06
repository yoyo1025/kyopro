package main

import "fmt"

// ============================================================
// 2章: グラフの表現
// ============================================================
//
// グラフの表現方法は大きく2種類あります。
//   - 隣接リスト: ノードごとに「繋がっているノードの一覧」を持つ
//   - 隣接行列:   N×N の 2次元配列で辺の有無を表す
//
// 競プロ・実務ともに隣接リストが一般的です。
// ノード数 N が大きく辺が少ないスパースグラフでも無駄がありません。
//
// 本サンプルでは map[int][]int を使った実装を示します。

// Graph は隣接リストによる無向グラフを表す
type Graph struct {
	nodes int           // ノード数
	edges map[int][]int // 隣接リスト
}

// NewGraph は nodes 個のノードを持つ空グラフを作る
func NewGraph(nodes int) *Graph {
	return &Graph{
		nodes: nodes,
		edges: make(map[int][]int),
	}
}

// AddEdge は無向グラフの辺 u-v を追加する
func (g *Graph) AddEdge(u, v int) {
	g.edges[u] = append(g.edges[u], v)
	g.edges[v] = append(g.edges[v], u)
}

// Print はグラフの隣接リストを表示する
func (g *Graph) Print() {
	for i := 0; i < g.nodes; i++ {
		fmt.Printf("  node %d -> %v\n", i, g.edges[i])
	}
}

func main() {
	g := NewGraph(6)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 3)
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)

	fmt.Println("隣接リスト:")
	g.Print()
}
