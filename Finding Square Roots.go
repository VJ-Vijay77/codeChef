package main

import (
	"fmt"
	"math"
)

func main() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var x int
		fmt.Scan(&x)
		fmt.Println(int(math.Sqrt(float64(x))))
}
}