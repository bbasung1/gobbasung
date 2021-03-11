package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type mb struct {
	o int
	p string
}

func main() {
	r := bufio.NewReader(os.Stdin)
	bw := bufio.NewWriter(os.Stdout)
	var a int
	fmt.Fscan(r, &a)
	b := make([]mb, a)
	for i := 0; i < a; i++ {
		fmt.Fscan(r, &b[i].o, &b[i].p)
	}
	sort.SliceStable(b, func(i, j int) bool {
		return b[i].o < b[j].o
	})
	for i := 0; i < a; i++ {
		fmt.Fprintln(bw, b[i].o, b[i].p)
	}
	bw.Flush()
}
