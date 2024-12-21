package main

import "fmt"

func main5t() {
	var T, X int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&X)
		income := X*50
		sugarcane := income *20/100
		saltAndMint := income *20/100
		rent := income *30/100
		fmt.Println(income-(sugarcane+saltAndMint+rent))

	}
}
