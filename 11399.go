package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	var a int
	fmt.Fscan(r, &a)
	b := make([]int, a)
	for i := 0; i < a; i++ {
		fmt.Fscanf(r, "%d ", &b[i])
	}
	sort.Ints(b)
	fmt.Println(b)
	sum := 0
	for i := 0; i < a; i++ {
		sum += b[i] * (a - i)
	}
	fmt.Fprintf(w, "%d\n", sum)
	w.Flush()
}
