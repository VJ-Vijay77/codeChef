package main

import "fmt"

func maikn() {
	var T, P,Q,R,S int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&P)
		fmt.Scan(&Q)
		fmt.Scan(&R)
		fmt.Scan(&S)
		
		if P > (Q+R+S) || Q > (P+R+S) || R > (P+Q+S) || S > (P+Q+R) {
			fmt.Println("YES")
		}else{
			fmt.Println("NO")
		}
	}
}
