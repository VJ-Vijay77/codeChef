package main

import "fmt"

func main09j() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var n,x,k int
		fmt.Scan(&n, &x, &k)

		if x > k {
			fmt.Println(0)
			continue
		}
		no := k/x
		if no > n {
			fmt.Println(n)
		}else{
		fmt.Println(no)
		}
	}
}
