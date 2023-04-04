package main

func checkInclusion(s1 string, s2 string) bool {
	// 判断特殊情况: s1长度大于s2
	if len(s1) > len(s2) {
		return false
	}
	// arr1记录s1的字符出现情况, arr2记录sub的字符出现情况
	var arr1, arr2 [26]int32
	for _, b := range s1 {
		arr1[b-97]++
	}
	sub := s2[0:len(s1)]
	for _, b := range sub {
		arr2[b-97]++
	}
	for i := 0; i < len(s2)-len(s1)+1; i++ {
		// 判断sub是否为s1的变位词
		j := 0
		for ; j < 26; j++ {
			if arr1[j] != arr2[j] {
				break
			}
		}
		if j == 26 {
			return true
		}
		// 删除sub最左边的字符并把后一个字符添加进来
		arr2[s2[i]-97]--
		if i+len(s1) < len(s2) {
			arr2[s2[i+len(s1)]-97]++
		}
	}
	return false
}
