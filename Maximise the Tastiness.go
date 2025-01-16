package main

import "fmt"

func mainsss() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var a,b,c,d int
		fmt.Scan(&a,&b,&c,&d)
		 firstIng := a
		secondIng := c
		if b > a {firstIng = b}
		if d > c {secondIng = d}
		fmt.Println(firstIng+secondIng)
	}
}
