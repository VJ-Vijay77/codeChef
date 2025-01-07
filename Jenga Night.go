package main

import "fmt"

func maibw() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var n,x int
		fmt.Scan(&n, &x)
        if x < n {
			fmt.Println("NO")
			continue
		}
		if x % n == 0{
			fmt.Println("YES")
		}else{
			fmt.Println("NO")
		}
		
	}
}
