package main

import (
	"fmt"
	"strings"
)

var (
	visited [][]bool
	S       [][]string
	H, W    int
)

var dx = []int{-1, 0, 1, 0}
var dy = []int{0, 1, 0, -1}

func main() {
	fmt.Scanf("%d %d", &H, &W)

	S = make([][]string, H)
	var line string
	for i := 0; i < H; i++ {
		fmt.Scanf("%s", &line)
		S[i] = make([]string, W)
		S[i] = strings.Split(line, "")
	}
	visited = make([][]bool, H)
	for i := 0; i < H; i++ {
		visited[i] = make([]bool, W)
	}
	ans := 0
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if S[i][j] == "." && !visited[i][j] {
				if dfs(i, j) {
					ans++
				}
			}
		}
	}
	fmt.Println(ans)
}

func dfs(i, j int) bool {
	visited[i][j] = true
	touchBoarder := i == 0 || i == H-1 || j == 0 || j == W-1
	for k := 0; k < 4; k++ {
		nextI := i + dx[k]
		nextJ := j + dy[k]
		// エリア外を探索しようとした場合
		if nextI < 0 || nextI > H-1 || nextJ < 0 || nextJ > W-1 {
			continue
		}
		// すでに訪れた場所の場合
		if visited[nextI][nextJ] {
			continue
		}
		// 黒の場合
		if S[nextI][nextJ] == "#" {
			continue
		}
		if !dfs(nextI, nextJ) {
			touchBoarder = true
		}
	}
	return !touchBoarder
}
