package main

func majorityElement(nums []int) int {
	// 摩尔投票法(有验证版本)
	votes, count := 0, 0
	x := nums[0]
	for _, num := range nums {
		if votes == 0 {
			x = num
		}
		if num == x {
			votes++
		} else {
			votes--
		}
	}
	for _, num := range nums {
		if num == x {
			count++
		}
	}
	if count > len(nums)/2 {
		return x
	}
	return 0
}

func main() {
	nums := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	majorityElement(nums)
}
