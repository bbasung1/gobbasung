package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var a int
	fmt.Fscan(r, &a)
	b := make(map[string]int)
	for i := a; i > 0; i-- {
		var d string
		var e int
		fmt.Fscan(r, &e, &d)
		b[d] = e
	}
	type kv struct {
		key string
		val int
	}
	var ss []kv
	for k, v := range b {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].val < ss[j].val
	})
	for _, kv := range ss {
		fmt.Println(kv.val, kv.key)
	}
}
