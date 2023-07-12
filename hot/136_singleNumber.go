package main

func singleNumber(nums []int) (ans int) {
	ans = nums[0]
	for i := 1; i < len(nums); i++ {
		ans ^= nums[i]
	}
	return ans
}
