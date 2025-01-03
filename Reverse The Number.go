package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var num int
		fmt.Scanln(&num)

		reversed := 0

		for num != 0 {
			digit := num % 10

			reversed = reversed*10 + digit

			num /= 10
		}
		fmt.Println(reversed)
	}
}
