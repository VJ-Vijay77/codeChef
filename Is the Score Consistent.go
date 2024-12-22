package main

import "fmt"

func main() {
	var T, A,B,C,D int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&A)
		fmt.Scan(&B)
		fmt.Scan(&C)
		fmt.Scan(&D)
		
		if (C>=A) && (D>=B) {
			fmt.Println("POSSIBLE")
		}else{
			fmt.Println("IMPOSSIBLE")
		}
	}
}
