package main

import "fmt"

func maipon() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var N int
		fmt.Scan(&N)
		fmt.Println(checkSeat(N))
	}
}

func checkSeat(num int) string{
	if num >=1 && num <=10{
		return "Lower Double"
	}
	if num >=11 && num <=15{
		return "Lower Single"
	}
	if num >=16 && num <=25{
		return "Upper Double"
	}
	if num >=26 && num <=30{
		return "Upper Single"
	}
	return ""
}