package main

import "fmt"

func main43() {
	var T, N int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&N)
		count := 0
		for j := 0; j < N; j++ {
			var NV int
			fmt.Scan(&NV)
			if NV >= 1000 {
				count++
			}
		}
		fmt.Println(count)
	}
}
