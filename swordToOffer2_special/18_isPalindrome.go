package main

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		for ; i < j && !isLetterOrNumber(s[i]); i++ {
		}
		for ; i < j && !isLetterOrNumber(s[j]); j-- {
		}
		// 用位运算代替判断大小写
		if s[i] != s[j] && s[i] != (s[j]^32) && i < j {
			return false
		}
		i++
		j--
	}
	return true
}

func equal(b1 uint8, b2 uint8) bool {
	// 判断两个字符是否相等
	if b1 == b2 {
		// b1和b2相等
		return true
	}
	if (b1 >= 'A' && b1 <= 'Z') && (b2 >= 'a' && b2 <= 'z') && b1 == b2-'a'+'A' {
		// b1大写, b2小写
		return true
	}
	if (b1 >= 'a' && b1 <= 'z') && (b2 >= 'A' && b2 <= 'Z') && b1 == b2-'A'+'a' {
		// b1小写, b2大写
		return true
	}
	return false
}

func isLetterOrNumber(b uint8) bool {
	// 判断一个字符是否为字母或数字
	return (b >= '0' && b <= '9') || (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z')
}
