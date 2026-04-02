package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func makeString(n int) string {
	if n <= 0 {
		return ""
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	c1 := byte('a' + r.Intn(26))

	// N が奇数なら1文字だけでOK
	if n%2 == 1 {
		return strings.Repeat(string(c1), n)
	}

	// N が偶数なら別の文字を1回だけ使う
	c2 := c1
	for c2 == c1 {
		c2 = byte('a' + r.Intn(26))
	}

	return strings.Repeat(string(c1), n-1) + string(c2)
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(makeString(n))
}
