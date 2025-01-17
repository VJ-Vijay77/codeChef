package main

import "fmt"

func mainpo() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var n, m int
		fmt.Scan(&n, &m)
		if m > n {
			fmt.Println(n)
			continue
		}
		leftShoes := n-m
		fmt.Println(leftShoes+n) 
		
	}
}
