package main

import (
	"bufio"
	"fmt"
	"os"
)

func lowerBound(a []int, x int) int {
	// 二分探索開始
	// 左端と右端を定義
	l := 0
	r := len(a)
	var m int

	for l < r {
		m = (l + r) / 2
		if a[m] >= x {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

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

	// 例：各 B[i] について lower_bound の位置を出力
	// （必要なら +1 して 1-indexed にしてね）
	for i := 0; i < m; i++ {
		fmt.Println(lowerBound(A, B[i]))
	}
}
