package main

import "fmt"

func main5t() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var x,y int
		fmt.Scan(&x,&y)
		if x>=y {
			fmt.Println(0)
			continue
		}
		diff  := y-x
		rounds :=  (diff + 7)/8
		fmt.Println(rounds)
	}
}
