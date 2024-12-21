package main

import "fmt"

func main() {
	var T, X,N int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&N)
		fmt.Scan(&X)
		if X > N {
			fmt.Println(0)
			continue
		}
	remainingCandiesToBuy := N-X
		buy := remainingCandiesToBuy/4
		if remainingCandiesToBuy%4 != 0 {
			buy += 1
		}
		fmt.Println(buy)
	}
}
