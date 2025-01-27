package main

import "fmt"

func main() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var x,y,z int
		fmt.Scan( &x, &y, &z)

		totalMin := x*y
		if x>3 {
			rest := (x/3) * z
			if x%3 == 0 {
				rest -= z
			}
			totalMin += rest
		}
		fmt.Println(totalMin)
		
	}
}
