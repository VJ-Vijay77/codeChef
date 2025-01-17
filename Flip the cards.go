package main

import "fmt"

func mainth() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var n, x int
		fmt.Scan(&n, &x)
		up,down,count := x,x,0
		for {
			if up==0 || down==0 ||up==n || down==n {
				break
			}
			up--
			down++
			count++
		}
		
		fmt.Println(count)
	}
}
