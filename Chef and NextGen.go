package main

import "fmt"

func maisg() {
	var T, X,Y,A,B int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&A)
		fmt.Scan(&B)
		fmt.Scan(&X)
		fmt.Scan(&Y)
		if (X*Y) >= (A*B) {
			fmt.Println("YES")
		}else{
			fmt.Println("NO")
		}
	}
}
