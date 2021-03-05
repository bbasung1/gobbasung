package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)
	var b = 0
	for c := 64; c > 1; c /= 2 {
		if a-c < 0 {
			continue
		}
		a -= c
		b++
	}
	if a == 1 {
		b++
	}
	fmt.Printf("%d\n", b)
}
