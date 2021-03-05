package main

import (
	"fmt"
)

func main() {
	var a string
	fmt.Scan(&a)
	var b int = len(a)
	for i := 0; i < b/10; i++ {
		fmt.Println(a[i*10 : (i*10)+10])
	}
	if b%10 != 0 {
		fmt.Println(a[(b/10)*10 : len(a)])
	}
}
