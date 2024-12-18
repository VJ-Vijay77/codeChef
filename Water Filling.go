package main

import "fmt"

func mainrt() {
	var T, X,Y,Z int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		count := 0
		fmt.Scan(&X)
		fmt.Scan(&Y)
		fmt.Scan(&Z)
		if X == 0 {
			count++
		}
		if Y == 0 {
			count++
		}
		if Z == 0 {
			count++
		}
		if count > 1 {
			fmt.Println("Water filling time")
		}else{
			fmt.Println("Not now")
		}
	}
}
