package main

import "fmt"

func mainrty() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var n,x,y int
		fmt.Scan(&n, &x, &y)

		if y%x == 0{
			fmt.Println("YES")
		}else{
			fmt.Println("NO")
		}
		
	}
}
