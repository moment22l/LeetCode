package main

// 设置一个二维字符数组，用以记录每行字符
// 设置flag记录遍历过程中行+还是行-,碰到第一行或最后一行时反向
// 遍历完字符串后，再将二维数组中的所有一维数组转变为字符串并依次拼接，即得答案
func convert2(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	rowBytes := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		rowBytes[i] = make([]byte, 0)
	}
	i := 0
	flag := -1
	for _, ch := range s {
		rowBytes[i] = append(rowBytes[i], byte(ch))
		if i == 0 || i == numRows-1 {
			flag = -flag
		}
		i += flag
	}
	outputS := ""
	for i := 0; i < numRows; i++ {
		outputS = outputS + string(rowBytes[i])
	}
	return outputS
}
