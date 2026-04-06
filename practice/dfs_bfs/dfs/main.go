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

func DFSRecursive(graph *Graph, start int) {
	visited := make(map[int]bool)
	fmt.Print("DFS（再帰）訪問順：")
	dfsRec(graph, start, visited)
	fmt.Println()
}

func dfsRec(graph *Graph, node int, visited map[int]bool) {
	visited[node] = true
	fmt.Printf("%d ", node)
	for _, next := range graph.edges[node] {
		if !visited[next] {
			dfsRec(graph, next, visited)
		}
	}
}

func DFSStack(graph *Graph, start int) {
	visited := make(map[int]bool)
	stack := []int{start}

	fmt.Print("DFS（スタック）訪問順：")
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if visited[node] {
			continue
		}
		visited[node] = true
		fmt.Printf("%d ", node)

		for _, v := range graph.edges[node] {
			stack = append(stack, v)
		}
	}
	fmt.Println()
}

func main() {
	g := NewGraph(10)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	g.AddEdge(2, 3)
	g.AddEdge(3, 6)
	g.AddEdge(4, 5)
	g.AddEdge(4, 7)
	g.AddEdge(5, 8)
	g.AddEdge(5, 9)
	g.AddEdge(6, 9)

	fmt.Println("隣接リスト:")
	g.Print()

	DFSRecursive(g, 0)
	DFSStack(g, 0)
}
