package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc *bufio.Scanner

func init() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
}

func scanInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func main() {
	n := scanInt()
	min := 10000000
	max := 0
	for i := 0; i < n; i++ {
		a := scanInt()
		if min > a {
			min = a
		}
		if max < a {
			max = a
		}
	}
	fmt.Printf("%d\n", min*max)
}
