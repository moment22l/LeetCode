package main

func countDigits(num int) int {
	temp := num
	ans := 0
	for temp != 0 {
		if num%(temp%10) == 0 {
			ans++
		}
		temp /= 10
	}
	return ans
}
