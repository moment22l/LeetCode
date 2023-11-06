package main

func maxProduct(words []string) int {
	n := len(words)
	nums := make([]int, n)
	for i, word := range words {
		for _, b := range word {
			nums[i] = nums[i] | (1 << (b - 'a'))
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]&nums[j] == 0 {
				if len(words[i])*len(words[j]) > ans {
					ans = len(words[i]) * len(words[j])
				}
			}
		}
	}
	return ans
}
