package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	h := &IntHeap{}
	heap.Init(h)

	var Q int
	fmt.Scanf("%d", &Q)
	sc := bufio.NewScanner(os.Stdin)

	for i := 0; i < Q; i++ {
		sc.Scan()
		inputs := strings.Split(sc.Text(), " ")
		u, _ := strconv.Atoi(inputs[0])
		v, _ := strconv.Atoi(inputs[1])

		if u == 1 {
			heap.Push(h, v)
		} else {
			for h.Len() > 0 && (*h)[0] <= v {
				heap.Pop(h)
			}
		}
		fmt.Println(h.Len())
	}
}
