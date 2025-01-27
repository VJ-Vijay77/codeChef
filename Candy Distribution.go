package main

import "fmt"

func maisgvbnn() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var n,m int
		fmt.Scan(&n, &m)

		if n%m == 0 && (n/m)%2 == 0 {
			fmt.Println("YES")
		}else{
			fmt.Println("NO")
		}
	}
}
