package main

import (
	"fmt"
	"math"
	"strconv"
)

// 1. 首先找到n所属的区间(0~9,10~99,100~999...), 即找到n的所在数字的位数
// 	  方法: n-9-1是否小于9*1, 不小于则-9*1继续判断是否小于90*2, 不小于则-90*2继续判断是否小于900*3, 以此类推
// 2. 找到在当前区间中n为一个完整数字的第几位
//    例如: n在100~999时, 如果(1)中判断后剩余数字为14
//         则n在一个完整数字的14%3=2位(即n所指向位置为一个完整数字的第2位(位数从0开始)
// 3. 找到在当前区间中n所处的那个完整数字
//    方法: n在100~999时, 如果(1)中判断后剩余数字x = 14, 则指向的是第x/3=4个(个数从0开始)
//         那么从100开始数第4个数字的第2位'4'即为所需数字

// 变量含义:
// numDigit: 当前数字区间的数字位数
// count: 当前数字区间的数字个数
// num: n在当前数字区间的第num位
// digit: n在当前区间的一个完整数字中的第digit位(位数从0开始)
// rank: n在当前区间的第rank个完整数字
// fullNumber: 当前区间的第rank个完整数字
func findNthDigit(n int) int {
	//if n >= 0 && n <= 9 {
	//	return n
	//}
	//numDigit := 1
	//count := 9
	//num := n - 1
	//for num >= count*numDigit {
	//	num -= count * numDigit
	//	count *= 10
	//	numDigit++
	//}
	//digit := num % numDigit
	//rank := num / numDigit
	//fullNumber := count/9 + rank
	//divisor := int(math.Pow10(numDigit - 1))
	//res := 0
	//for i := 0; i <= digit; i++ {
	//	res = fullNumber / divisor
	//	fullNumber -= res * divisor
	//	divisor /= 10
	//}
	//return res

	k := n
	for digit := 1; ; digit++ {
		if k < digit*int(math.Pow10(digit)) {
			return int(strconv.Itoa(k / digit)[k%digit] - '0')
		}
		k += int(math.Pow10(digit))
	}
}

func main() {
	fmt.Println(findNthDigit(13))
}
