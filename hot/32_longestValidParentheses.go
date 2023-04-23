package main

// 动态规划
// 当s[i]=='('时, dp[i]=0
// 当s[i]==')'&&s[i-1]=='('时, dp[i] = dp[i-2]+2
// 当s[i]==')'&&s[i-1]==')'时, 若dp[i-dp[i-1]-1]=='(', 则dp[i] = dp[i-1]+2+dp[i-dp[i-1]-2] 否则dp[i]=0
//func longestValidParentheses(s string) (ans int) {
//	dp := make([]int, len(s))
//	for i := 1; i < len(s); i++ {
//		if s[i] == ')' {
//			if s[i-1] == '(' {
//				if i >= 2 {
//					dp[i] = dp[i-2] + 2
//				} else {
//					dp[i] = 2
//				}
//			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
//				if i-dp[i-1]-2 > 0 {
//					dp[i] = dp[i-1] + 2 + dp[i-dp[i-1]-2]
//				} else {
//					dp[i] = dp[i-1] + 2
//				}
//			}
//			ans = max32(ans, dp[i])
//		}
//	}
//	return
//}
//
//func max32(x, y int) int {
//	if x > y {
//		return x
//	}
//	return y
//}

// 滑动窗口，从前往后以及从后往前，一旦count<0以及count>0了都将count置0并从下一位置开始遍历
// 时间O(n) 空间O(1)
func longestValidParentheses(s string) (ans int) {
	n := len(s)
	if n == 0 || n == 1 {
		return 0
	}
	index, count := 0, 0
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			count++
		} else {
			count--
		}
		if count == 0 {
			ans = max32(ans, i-index+1)
		} else if count < 0 {
			index = i + 1
			count = 0
		}
	}
	index, count = n-1, 0
	for i := n - 1; i >= 0; i-- {
		if s[i] == '(' {
			count++
		} else {
			count--
		}
		if count == 0 {
			ans = max32(ans, index-i+1)
		} else if count > 0 {
			index = i - 1
			count = 0
		}
	}
	return
}

func max32(x, y int) int {
	if x > y {
		return x
	}
	return y
}
