package main

import "fmt"

func mainq() {
	var N,ready,notReady int
	fmt.Scan(&N)


	for i:=0 ; i< N ; i++ {
		var a int
		fmt.Scan(&a)
		if a % 2 == 0 {
			ready++
		}else{
			notReady++
		}
	}
	if ready>notReady{
		fmt.Println("READY FOR BATTLE")
	}else{
		fmt.Println("NOT READY")
	}
}
