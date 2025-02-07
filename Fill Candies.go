package main

import "fmt"

func mainsfsd() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var a,b,c int
		fmt.Scan(&a, &b, &c)
		
		oneBag := b*c
		totalBag := a/oneBag
		if a%oneBag != 0 {
			totalBag+=1
		}
		fmt.Println(totalBag)
	}
}
