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
	var c [][]int
	for i := a; i > 0; i-- {
		b := []int{0, 0}
		fmt.Fscan(r, &b[0], &b[1])
		c = append(c, b)
	}
	sort.SliceStable(c, func(i, j int) bool {
		if c[i][1] == c[j][1] {
			return c[i][0] < c[j][0]
		} else {
			return c[i][1] < c[j][1]
		}
	})
	for _, val := range c {
		fmt.Fprintln(w, val[0], val[1])
	}
	w.Flush()
}
