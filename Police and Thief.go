package main

import "fmt"

func maisdfn() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var x,y int
		fmt.Scan(&x,&y)
		time := 0
		for {
			if x == y{
				break
			}
			if x<y {
				x+=2
				y+=1
			}else{
				x-=2
				y-=1
			}
			time++
		}
		fmt.Println(time)
	}
}
