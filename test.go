package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scanln(&a)
	b := make([]int, a+1)
	for i := 0; i < a+1; i++ {
		fmt.Println(b[i])
	}
}
