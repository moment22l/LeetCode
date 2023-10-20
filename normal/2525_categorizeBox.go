package main

import "math"

func categorizeBox(length int, width int, height int, mass int) string {
	m := map[int]string{
		0: "Neither",
		1: "Bulky",
		2: "Heavy",
		3: "Both",
	}
	num := 0
	if length >= int(math.Pow10(4)) || width >= int(math.Pow10(4)) ||
		height >= int(math.Pow10(4)) || length*width*height >= int(math.Pow10(9)) {
		num += 1
	}
	if mass >= 100 {
		num += 2
	}
	return m[num]
}
