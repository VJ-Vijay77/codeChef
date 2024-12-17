package main

import "fmt"

func main2() {
	var T, X int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&X)
		if (10 - X) >= 3 {
			fmt.Println("yes")
			continue
		} 
		fmt.Println("no")
	}
}
