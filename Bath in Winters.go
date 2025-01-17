package main

import "fmt"

func mainssv() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var y, x int
		fmt.Scan(&x,&y)
		forOnePerson := y*2
		fmt.Println(x/forOnePerson)
	}
}
