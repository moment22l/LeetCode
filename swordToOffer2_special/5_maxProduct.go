package main

//	题目: 单词长度的最大乘积
//	给定一个字符串数组 words，请计算当两个字符串 words[i] 和 words[j] 不包含相同字符时，它们长度的乘积的最大值。
//	假设字符串中只包含英语的小写字母。如果没有不包含相同字符的一对字符串，返回 0。

//func maxProduct(words []string) int {
//	var wordsMap [][]bool
//	wordsMap = make([][]bool, len(words))
//	for i := 0; i < len(words); i++ {
//		wordsMap[i] = make([]bool, 26)
//		for _, word := range words[i] {
//			wordsMap[i][word-'a'] = true
//		}
//	}
//	max := 0
//	for i := 0; i < len(words); i++ {
//		for j := i + 1; j < len(words); j++ {
//			if len(words[i])*len(words[j]) <= max {
//				continue
//			}
//			// 判断是否含相同字符
//			k := 0
//			for ; k < 26; k++ {
//				if wordsMap[i][k] && wordsMap[j][k] {
//					break
//				}
//			}
//			if k == 26 {
//				max = len(words[i]) * len(words[j])
//			}
//		}
//	}
//	return max
//}

func maxProduct(words []string) int {
	wordsMap := make([]int, len(words))
	for i, word := range words {
		for _, ch := range word {
			wordsMap[i] |= 1 << (ch - 'a')
		}
	}
	max := 0
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if len(words[i])*len(words[j]) <= max {
				continue
			}
			// 判断是否含相同字符
			if wordsMap[i]&wordsMap[j] == 0 {
				max = len(words[i]) * len(words[j])
			}
		}
	}
	return max
}
