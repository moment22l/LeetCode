package main

func isMatch(s string, p string) bool {
	// 动态规划
	sb := []byte(s)
	pb := []byte(p)
	i, j := 0, 0
	dp := make([][]int, len(sb))
	for n := 0; n < len(dp); n++ {
		dp[n] = make([]int, len(pb))
	}
	for i < len(sb) && j < len(pb) {

	}
	return true
}

func main() {

}
