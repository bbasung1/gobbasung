package main

import (
	"fmt"
	"strings"
)

func main() {
	var a, b bool
	var c string
	fmt.Scan(&a, &c, &b)
	if strings.Contains(c, "AND") {
		fmt.Println(a && b)
	} else {
		fmt.Println(a || b)
	}

}
