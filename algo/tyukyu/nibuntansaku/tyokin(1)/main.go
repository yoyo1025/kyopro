package main

import "fmt"

func main() {
	// 初年度の預け額、5年後の預け額
	var n, m float64
	_, err := fmt.Scan(&n, &m)
	if err != nil {
		fmt.Println("読み込みエラーが発生しました。")
		return
	}

	var left float64 = 0
	var right float64 = m
	var mid float64 = (left + right) / 2

	// 許容誤差
	var tolerance float64 = 0.01

	for {
		result := calcDeposit(n, mid)
		if m-tolerance <= result && result <= m+tolerance {
			fmt.Println(mid)
			break
		}
		if result < m {
			left = mid
		}

		if result > m {
			right = mid
		}
		mid = (left + right) / 2
	}
}

func calcDeposit(n float64, rate float64) float64 {
	d0 := n + 1
	d1 := d0*rate + 1
	d2 := d1*rate + 1
	d3 := d2*rate + 1
	d4 := d3*rate + 1
	d5 := d4*rate + 1
	return d5
}
