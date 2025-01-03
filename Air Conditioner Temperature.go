package main

import "fmt"

func maisfn() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var A,B,C int
		fmt.Scan(&A, &B, &C)
		
		if max(A,C) <= B {
			fmt.Println("YES")
		}else {
			fmt.Println("NO")
		}
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
