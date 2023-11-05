package main

func findRepeatedDnaSequences(s string) []string {
	bin := map[byte]int{'A': 0, 'C': 1, 'G': 2, 'T': 3}
	L := 10
	n := len(s)
	if n <= L {
		return []string{}
	}
	ans := make([]string, 0)
	x := 0
	for _, ch := range s[:L-1] {
		x = x<<2 | bin[byte(ch)]
	}
	cnt := make(map[int]int)
	for i := L - 1; i < n; i++ {
		x = (x<<2 | bin[s[i]]) & (1<<(L*2) - 1)
		cnt[x]++
		if cnt[x] == 2 {
			ans = append(ans, s[i-L+1:i+1])
		}
	}
	return ans
}
