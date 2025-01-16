package main

import (
	"fmt"
)

func mai42() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var a, b, c int
		fmt.Scan(&a, &b, &c)
		maxScore := 0
		if a > b {
			maxScore = a
		} else {
			maxScore = b
		}
		if c > maxScore {
			maxScore = c
		}
		fmt.Println(maxScore)
	}
}
