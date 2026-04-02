package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)

	b := make([][]int, n)
	for i := 0; i < n; i++ {
		b[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &b[i][j])
		}
	}

	// 呼ばれた数字を記録
	called := make(map[int]bool)
	for i := 0; i < k; i++ {
		var x int
		fmt.Fscan(in, &x)
		called[x] = true
	}

	// マスを埋める
	marked := make([][]bool, n)
	for i := 0; i < n; i++ {
		marked[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			if called[b[i][j]] {
				marked[i][j] = true
			}
		}
	}
	if n%2 == 1 {
		mid := n / 2
		marked[mid][mid] = true
	}

	ans := 0

	// 各行を判定
	for i := 0; i < n; i++ {
		ok := true
		for j := 0; j < n; j++ {
			if !marked[i][j] {
				ok = false
				break
			}
		}
		if ok {
			ans++
		}
	}

	// 各列を判定
	for j := 0; j < n; j++ {
		ok := true
		for i := 0; i < n; i++ {
			if !marked[i][j] {
				ok = false
				break
			}
		}
		if ok {
			ans++
		}
	}

	// 左上 → 右下
	ok := true
	for i := 0; i < n; i++ {
		if !marked[i][i] {
			ok = false
			break
		}
	}
	if ok {
		ans++
	}

	// 右上 → 左下
	ok = true
	for i := 0; i < n; i++ {
		if !marked[i][n-1-i] {
			ok = false
			break
		}
	}
	if ok {
		ans++
	}

	fmt.Println(ans)
}
