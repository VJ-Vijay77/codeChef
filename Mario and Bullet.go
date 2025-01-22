package main

import "fmt"

func mainxcv() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var x,y,z int
		fmt.Scan(&x,&y,&z)

		time := y/x
		if time >= z {
			fmt.Println(0)
		}else{
			delay := z-time
			fmt.Println(delay)
		}
	}
}
