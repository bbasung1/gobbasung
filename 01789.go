package main

import (
	"fmt"
)

func main() {
	var s, n uint64
	fmt.Scan(&s)
	n = 0
	for ; s > n*(n+1)/2; n++ {

	}
	if s == n*(n+1)/2 {
		fmt.Println(n)
	} else {
		fmt.Println(n - 1)
	}
}
