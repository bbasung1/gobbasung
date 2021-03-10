package main

import (
	"fmt"
)

func main() {
	var a, b, c, e uint64
	var n int
	fmt.Scan(&n)
	a = 0
	b = 1
	for i := n; i > 0; i-- {
		c = a + b
		if i == 1 {
			e = a
		}
		a = b
		b = c
	}
	fmt.Println(4*a + 2*e)
}
