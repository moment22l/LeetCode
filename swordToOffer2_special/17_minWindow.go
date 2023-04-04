package main

func minWindow(s string, t string) string {
	// 因为只有大小写英文字母, 理论上用len=52的数组即可, 但为了方便这里使用了len=58的数组
	// arrT记录了t的字母出现次数情况, arrS则记录滑动窗口也就是s的子串的字母出现次数情况
	var arrT [58]int
	for _, r := range t {
		arrT[r-65]++
	}
	var arrS [58]int
	left, right := -1, 0 // 双指针
	var ans string
	for ; right < len(s); right++ {
		arrS[s[right]-65]++
		if contain(arrS, arrT) {
			for left < right && contain(arrS, arrT) {
				left++
				arrS[s[left]-65]--
			}
			if right-left+1 < len(ans) || len(ans) == 0 {
				ans = s[left : right+1]
			}
		}
	}
	return ans
}

func contain(arr1 [58]int, arr2 [58]int) bool {
	// 判断arr1是否包含arr2(在本例中表示滑动窗口(即s的子串)是否包含t)
	for i := 0; i < 58; i++ {
		if arr1[i] < arr2[i] {
			return false
		}
	}
	return true
}
