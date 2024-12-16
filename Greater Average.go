package main

import (
	"fmt"
)

func mains() {
	// Read the number of test cases
	var T int
	fmt.Scan(&T)

	// Process each test case
	for i := 0; i < T; i++ {
		var A, B, C int
		fmt.Scan(&A, &B, &C)

		// Calculate the average of A and B
		average := float64(A+B) / 2.0

		// Check if the average is strictly greater than C
		if average > float64(C) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
