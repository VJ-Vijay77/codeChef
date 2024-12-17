package main

import "fmt"

func maine() {
	var T, X,Y int
	fmt.Scan(&T)

	for i := 0; i < int(T); i++ {
		fmt.Scan(&X)
		fmt.Scan(&Y)
		fmt.Println((X*4)+Y)
	}
}
