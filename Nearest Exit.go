package main

import "fmt"

func mascin() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var x int
		fmt.Scanln(&x)
		if x <= 50 {
			fmt.Println("LEFT")
		}else {
			fmt.Println("RIGHT")
		}

	}
}
