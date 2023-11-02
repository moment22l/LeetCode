package main

func countPoints(rings string) int {
	// 状态：
	// R:001 G:010 B:100 RG:011 RB:101 GB:110 RGB:111
	m := map[byte]int{
		'R': 0b001,
		'G': 0b010,
		'B': 0b100,
	}
	total := make([]int, 10)
	index := 0
	for index < len(rings) {
		total[rings[index+1]-'0'] |= m[rings[index]]
		index += 2
	}
	ans := 0
	for _, num := range total {
		if num == 0b111 {
			ans++
		}
	}
	return ans
}
