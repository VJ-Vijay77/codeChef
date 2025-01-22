package main

import (
	"fmt"
	"math/big"
)

func maiopn() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var n int
		fmt.Scan(&n)
		factorial := big.NewInt(1)
		for i:=2; i<=n; i++{
			factorial.Mul(factorial,big.NewInt(int64(i)))
			
		}
		fmt.Println(factorial)
	}
}
