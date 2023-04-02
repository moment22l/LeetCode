package main

import (
	"fmt"
)

func translateNum(num int) int {
	former, latter := 1, 1
	y := num % 10
	for num != 0 {
		num /= 10
		x := num % 10
		if x*10+y >= 10 && x*10+y <= 25 {
			former, latter = former+latter, former
		} else {
			latter = former
		}
		y = x
	}
	return former
}

func main() {
	num := 1000
	fmt.Println(translateNum(num))
}
