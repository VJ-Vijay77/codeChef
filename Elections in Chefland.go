package main

import "fmt"

func mai67() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var n,x int
		fmt.Scan(&n, &x)
		count := 0
		for j:=0; j<n; j++ {
			var s int
			fmt.Scan(&s)
			if s >= x {
				count++
			}
		}
		fmt.Println(count)
	}
}
