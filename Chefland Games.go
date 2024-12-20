package main

import "fmt"

func mainerw() {
	var T, X,Y,Z,ZZ int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&X)
		fmt.Scan(&Y)
		fmt.Scan(&Z)
		fmt.Scan(&ZZ)
		if X==0 && Y==0 && Z==0 && ZZ==0 {
			fmt.Println("IN")
		}else{
			fmt.Println("OUT")
		}
	}
}
