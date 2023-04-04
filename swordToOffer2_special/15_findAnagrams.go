package main

func findAnagrams(s string, p string) []int {
	// 与题14方法一致, 只是返回的东西不同
	// 判断特殊情况: p长度大于s
	if len(p) > len(s) {
		return []int{}
	}
	var result []int
	// arr1记录p的字符出现情况, arr2记录sub的字符出现情况
	var arr1, arr2 [26]int32
	for _, b := range p {
		arr1[b-97]++
	}
	sub := s[0:len(p)]
	for _, b := range sub {
		arr2[b-97]++
	}
	for i := 0; i < len(s)-len(p)+1; i++ {
		// 判断sub是否为p的变位词
		j := 0
		for ; j < 26; j++ {
			if arr1[j] != arr2[j] {
				break
			}
		}
		if j == 26 {
			result = append(result, i)
		}
		// 删除sub最左边的字符并把后一个字符添加进来
		arr2[s[i]-97]--
		if i+len(p) < len(s) {
			arr2[s[i+len(p)]-97]++
		}
	}
	return result
}
