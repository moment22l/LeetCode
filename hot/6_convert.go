package main

// 按周期划分字符串，逐行遍历行中每个周期中的字符，依次加入到字符串中
func convert(s string, numRows int) string {
	n := len(s)
	if numRows == 1 || numRows >= n { // 如果只有一行或只有一列，那输出串与输入串相等，直接返回
		return s
	}

	outputS := make([]byte, 0, n)
	x := 2 * (numRows - 1)
	for i := 0; i < numRows; i++ { // 枚举行
		for j := 0; j+i < n; j += x { // 枚举每个周期的第一个字母索引
			outputS = append(outputS, s[j+i])           // 将该行第一个字母加入串中
			if i > 0 && i < numRows-1 && j+x-i <= n-1 { // 判断该行是否不为第一或最后行，再判断第二个字母索引是否超出范围
				outputS = append(outputS, s[j+x-i]) // 满足要求则将周期内该行第二个字母也加入串中
			}
		}
	}
	return string(outputS)
}

/*
[N=1]
ABCDEF

[N=2]
ACEG
BDFH

[N=3]
A   E   I   M
B D F H J L N
C   G   K   O

[N=4]
P     I    N
A   L S  I G
Y A   H R
P     I

[N=5]
A   I   Q   Y
B  HJ  PR  XZ
C G K O S W
DF  LN  TV
E   M   U

*/
