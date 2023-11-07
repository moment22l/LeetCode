package main

func vowelStrings(words []string, left int, right int) int {
	ans := 0
	for _, word := range words[left : right+1] {
		if isVowel(word[0]) && isVowel(word[len(word)-1]) {
			ans++
		}
	}
	return ans
}

func isVowel(ch byte) bool {
	if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
		return true
	}
	return false
}
