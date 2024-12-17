package main

import "fmt"

func mainsd() {
	var T, X,Y,Z float32
	fmt.Scan(&T)

	for i := 0; i < int(T); i++ {
		fmt.Scan(&X)
		fmt.Scan(&Y)
		fmt.Scan(&Z)
		passed :=  (Z/(X*Y))*100
		if passed > 50 {
			fmt.Println("yes")
		}else{
			fmt.Println("no")
		}
	}
}
