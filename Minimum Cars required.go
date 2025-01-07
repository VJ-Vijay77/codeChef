package main

import "fmt"

func mainio() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var n int
		fmt.Scan(&n)

		if n <= 4 {
			fmt.Println(1)
			continue
		}

		s := n/4
		if n%4 !=0 {
			fmt.Println(s+1)
		}else{
			fmt.Println(s)
		}
		
	}
}
