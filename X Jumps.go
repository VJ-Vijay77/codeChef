package main

import "fmt"

func main() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var x,y int
		fmt.Scan(&x,&y)
		a := x/y
		if x % y !=0 {
			a = a + (x-(a*y))
	}
		fmt.Println(a)
	}
}
