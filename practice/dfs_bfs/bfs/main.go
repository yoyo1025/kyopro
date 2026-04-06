package main

import "fmt"

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
		fmt.Printf(" node %d -> %v\n", i, g.edges[i])
	}
}

func BFS(graph *Graph, start int) {
	visited := make(map[int]bool)
	queue := []int{start}

	fmt.Print("BFS（キュー）訪問順：")

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if visited[node] {
			continue
		}
		visited[node] = true
		fmt.Printf("%d ", node)
		for _, v := range graph.edges[node] {
			queue = append(queue, v)
		}
	}
	fmt.Println()
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

	BFS(g, 0)
}
