package main

import "fmt"

func mauiin() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var n int
		fmt.Scan(&n)
		if n % 4 == 0 {
			fmt.Println("GOOD")
		}else{
			fmt.Println("NOT GOOD")
		}
		
	}
}
