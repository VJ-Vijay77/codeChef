package main

import "fmt"

func maisdn() {
	var T, X,Y int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&X)
		fmt.Scan(&Y)

		if Y*100/X >= 50 {
			fmt.Println("YES")
		}else{
			fmt.Println("NO")
		}

	}
}
