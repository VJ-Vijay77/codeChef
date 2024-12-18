package main

import "fmt"

func mainse() {
	var T, X int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&X)
		if X <= 100 {
			fmt.Println(X)
		} else if X <= 1000 {
			fmt.Println(X - 25)
		} else if X <= 5000 {
			fmt.Println(X - 100)
		} else {
			fmt.Println(X - 500)
		}
	}
}
