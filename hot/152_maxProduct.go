package main

func maxProduct(nums []int) int {
	n := len(nums)
	maxDP := make([]int, n)
	minDP := make([]int, n)
	maxDP[0], minDP[0] = nums[0], nums[0]

	for i := 1; i < n; i++ {
		num := nums[i]
		maxDP[i] = max152_3(maxDP[i-1]*num, minDP[i-1]*num, num)
		minDP[i] = min152_3(maxDP[i-1]*num, minDP[i-1]*num, num)
	}
	ans := maxDP[0]
	for _, m := range maxDP {
		if m > ans {
			ans = m
		}
	}
	return ans
}

func max152_3(x, y, z int) (ans int) {
	ans = x
	if ans < y {
		ans = y
	}
	if ans < z {
		ans = z
	}
	return
}

func min152_3(x, y, z int) (ans int) {
	ans = x
	if ans > y {
		ans = y
	}
	if ans > z {
		ans = z
	}
	return
}
