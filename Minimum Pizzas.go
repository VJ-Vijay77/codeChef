package main

import "fmt"

func main() {
	var T, N,X int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&N)
		fmt.Scan(&X)
		noOfSlices := N*X
		quo := noOfSlices/4
		rem := noOfSlices % 4
		if rem != 0 {
			fmt.Println(quo+1)
		}else{
			fmt.Println(quo)
		}
	}
}
