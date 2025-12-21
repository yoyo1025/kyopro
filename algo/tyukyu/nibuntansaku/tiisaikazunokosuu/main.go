package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	if _, err := fmt.Fscan(in, &n, &m); err != nil {
		fmt.Println("読み取りエラーが発生しました。")
		return
	}

	// n 個の A、m 個の B を読む（入力が行区切りでなくても安全）
	A := make([]int, n)
	for i := 0; i < n; i++ {
		if _, err := fmt.Fscan(in, &A[i]); err != nil {
			fmt.Println("A の読み取りエラーが発生しました。")
			return
		}
	}

	B := make([]int, m)
	for i := 0; i < m; i++ {
		if _, err := fmt.Fscan(in, &B[i]); err != nil {
			fmt.Println("B の読み取りエラーが発生しました。")
			return
		}
	}

	// Aを昇順にソート
	sort.Ints(A)

	for i := 0; i < m; i++ {
		x := B[i]
		l, r := 0, len(A) // [l, r)
		for l < r {
			mid := (l + r) / 2
			if A[mid] > x {
				r = mid
			} else {
				l = mid + 1
			}
		}

		fmt.Println(l)
	}
}
