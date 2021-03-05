package main

import (
	. "fmt"
)

func main() {
	var a int
	Scan(&a)
	min := 1000001
	max := 0
	for i := a; i > 0; i-- {
		var b int
		Scan(&b)
		if min > b {
			min = b
		}
		if max < b {
			max = b
		}
	}
	Println(min * max)
}
