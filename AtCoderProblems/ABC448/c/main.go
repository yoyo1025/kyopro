package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type ball struct {
	index, val int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var N, Q int
	fmt.Fscan(reader, &N, &Q)

	balls := make([]ball, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &balls[i].val)
		balls[i].index = i + 1
	}
	sort.Slice(balls, func(i, j int) bool { return balls[i].val < balls[j].val })

	for i := 0; i < Q; i++ {
		var K int
		fmt.Fscan(reader, &K)

		exclude := make(map[int]bool, K)
		for j := 0; j < K; j++ {
			var b int
			fmt.Fscan(reader, &b)
			exclude[b] = true
		}
		for _, ball := range balls {
			if !exclude[ball.index] {
				fmt.Fprintln(writer, ball.val)
				break
			}
		}
	}
}
