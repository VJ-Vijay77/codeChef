package main

import "fmt"

func main4rf() {
	var T, X,Y int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&X)
		fmt.Scan(&Y)
		if X < Y {
			fmt.Println("BIKE")
		}else if Y< X{
			fmt.Println("CAR")
		}else{
			fmt.Println("SAME")
		}
	}
}
