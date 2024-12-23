package main

import "fmt"

func main34() {
	var  X,A,B,C int
	fmt.Scan(&A)
	fmt.Scan(&B)
	fmt.Scan(&C)
	fmt.Scan(&X)

	if X==A || X==B || X==C {
		fmt.Println("YES")
	}	else{
		fmt.Println("NO")
	}
}
