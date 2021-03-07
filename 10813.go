package main

import "fmt"

func main() {
	var m, n int
	fmt.Scanln(&n, &m)
	var test []int
	for i := 1; i <= n; i++ {
		test = append(test, i)
	}
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scanln(&a, &b)
		test[a-1], test[b-1] = test[b-1], test[a-1]
	}
	for _, val := range test {
		fmt.Printf("%d ", val)
	}
	fmt.Printf("\n")
}
