package main

import "fmt"

func maint() {
	var T, X int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&X)
		fmt.Println(X*10)
	}
}
