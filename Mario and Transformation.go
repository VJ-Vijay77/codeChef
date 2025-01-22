package main

import "fmt"

func mainwe() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var x int
		fmt.Scan(&x)

		if x%3 == 0 {
			fmt.Println("NORMAL")
		}else if x%3 == 1 {
			fmt.Println("HUGE")
		}else{
			fmt.Println("SMALL")
		}
		
	}
}
