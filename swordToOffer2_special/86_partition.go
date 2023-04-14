package main

// 回溯 时间O(n·2^n) 空间O(n^2)
func partition(s string) (ans [][]string) {
	length := len(s)
	// 预处理 f[i][j]表示从i到j的字符串是否为回文串
	f := make([][]bool, length)
	for i := range f {
		f[i] = make([]bool, length)
		for j := range f[i] {
			f[i][j] = true
		}
	}
	for i := length - 1; i >= 0; i-- {
		for j := i + 1; j < length; j++ {
			f[i][j] = s[i] == s[j] && f[i+1][j-1]
		}
	}
	var arr []string
	var backTrace func(index int)
	backTrace = func(index int) {
		if index == length {
			add := make([]string, len(arr))
			copy(add, arr)
			ans = append(ans, add)
			return
		}
		for i := index; i < length; i++ {
			if f[index][i] {
				arr = append(arr, s[index:i+1])
				backTrace(i + 1)
				arr = arr[:len(arr)-1]
			}
		}
	}
	backTrace(0)
	return
}
