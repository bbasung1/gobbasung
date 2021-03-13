package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type loc struct {
	x, y int
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	var a int
	fmt.Fscan(r, &a)
	c := make([]loc, a)
	for i := 0; i < a; i++ {
		fmt.Fscan(r, &c[i].x, &c[i].y)
	}
	sort.Slice(c, func(i, j int) bool {
		if c[i].y == c[j].y {
			return c[i].x < c[j].x
		} else {
			return c[i].y < c[j].y
		}
	})
	for _, val := range c {
		fmt.Fprintln(w, val.x, val.y)
	}
	w.Flush()
}
