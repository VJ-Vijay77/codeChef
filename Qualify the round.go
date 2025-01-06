package main

import "fmt"

func mapoin() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var x,a,b int
		fmt.Scan(&x,&a,&b)
		if a+(b*2) >= x{
			fmt.Println("Qualify")
		}else{
			fmt.Println("NotQualify")
		}
		
	}
}
