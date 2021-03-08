package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string
	fmt.Scanln(&a)
	slice := strings.Split(a, ",")
	fmt.Println(len(slice))
}
