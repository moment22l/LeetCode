package main

// 用一个map来存储各种异位词的当前strings
// 时间O(n*(k+∣Σ∣)) n=len(strs), k=最长str的长度, O(∣Σ∣)的时间生成哈希表的键, 此题中∣Σ∣=26也就是字符集长度
// 空间O(n(k+∣Σ∣))
func groupAnagrams(strs []string) (ans [][]string) {
	m := make(map[[26]int][]string, 0)
	for _, str := range strs {
		var temp [26]int
		for _, b := range str {
			temp[b-'a']++
		}
		if _, ok := m[temp]; ok {
			m[temp] = append(m[temp], str)
		} else {
			m[temp] = []string{str}
		}
	}
	for _, val := range m {
		ans = append(ans, val)
	}
	return
}
