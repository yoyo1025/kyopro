package main

import "fmt"

func main() {
	var num float64
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("エラー: ", err)
		return
	}

	// 求める実数Xの範囲が0以上100以下である
	var left float64 = 0
	var right float64 = 100
	var mid float64 = (left + right) / 2

	// 許容誤差
	var tolerance float64 = 0.01

	// 探索開始
	for {
		result := calcResult(mid)
		// 誤差が+-0.01までは許容
		if (num-tolerance) <= result && result <= (num+tolerance) {
			fmt.Println(mid)
			break
		}
		// num より result が小さい場合は x を大きくする必要があるので left = mid
		if result < num {
			left = mid
		}
		// num より result が大きい場合は x を小さくする必要があるので right = mid
		if result > num {
			right = mid
		}
		mid = (left + right) / 2
	}
}

// 条件式の左辺
func calcResult(x float64) float64 {
	return x*(x*(x+1)+2) + 3
}
