package main

import "fmt"

func mainsdfs() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var a,b int
		fmt.Scan(&a, &b)
		third := 21 - (a+b)
		if third > 10 {
			third = -1
		} 
		fmt.Println(third)
	}
}
