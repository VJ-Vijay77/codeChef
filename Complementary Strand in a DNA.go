package main

import "fmt"

func main() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var x int
		var s string
		fmt.Scan(&x)
		fmt.Scan(&s)
		output := make([]byte,x)
		for i:=0; i<x; i++ {
			switch s[i] {
			case 'A':
				output[i] = 'T'
			case 'T':
				output[i] = 'A'
			case 'C':
				output[i] = 'G'
			case 'G':
				output[i] = 'C'
			}
		}

		fmt.Println(string(output))
	}
}
