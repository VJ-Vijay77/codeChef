package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var a, b int
		fmt.Scan(&a, &b)

		if a*10 == b*5{
			fmt.Println("ANY")
		}else if a*10 > b*5 {
			fmt.Println("FIRST")
		}else{
			fmt.Println("SECOND")
		}
	}
}
