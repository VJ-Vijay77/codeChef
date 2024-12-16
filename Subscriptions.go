package main

import "fmt"

func main() {
	var T, X, Y int
	fmt.Scan(&T)
	var sub int
	for i := 0; i < T; i++ {
		fmt.Scan(&X)
		fmt.Scan(&Y)
		if X%6 == 0 {
			sub = X / 6
		} else {
			sub = (X / 6) + 1
		}
		fmt.Println(sub * Y)
	}

}
