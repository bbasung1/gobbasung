package main

import (
	"fmt"
	"sort"
)

func main() {
	var a int
	fmt.Scan(&a)
	list := []string{}
	for i := 0; i < a; i++ {
		var b string
		fmt.Scan(&b)
		list = append(list, b)
	}
	keys := make(map[string]struct{})
	res := make([]string, 0)
	for _, val := range list {
		if _, ok := keys[val]; ok {
			continue
		} else {
			keys[val] = struct{}{}
			res = append(res, val)
		}
	}
	r1 := sort.StringSlice(res)
	r1.Sort()
	for i := 1; i <= 50; i++ {
		for _, val := range r1 {
			if len(val) == i {
				fmt.Println(val)
			}
		}
	}
}
