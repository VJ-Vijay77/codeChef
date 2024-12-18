package main

import "fmt"

func maind() {
	var T, X,Y int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&X)
		fmt.Scan(&Y)
		coins := X*Y
		fmt.Println(coins/100)
	}
}
