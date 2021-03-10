package main

import (
	"fmt"
	"math/big"
)

func main() {
	var a, b big.Int
	fmt.Scan(&a, &b)
	mul := new(big.Int)
	mod := new(big.Int)
	mul = mul.Div(&a, &b)
	mod = mod.Mod(&a, &b)
	fmt.Println(mul)
	fmt.Println(mod)
}
