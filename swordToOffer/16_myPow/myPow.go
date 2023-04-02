package main

func myPow(x float64, n int) float64 {
	// 暴力解法(超时)
	var res float64 = 1
	if n == 0 {
		return 1
	} else if n > 0 {
		//for i := 0; i < n; i++ {
		//	res = res * x
		//}
		for ; n > 0; n >>= 1 {
			// 快速幂算法
			if n&1 == 1 {
				res = res * x
			}
			x = x * x
		}
		return res
	} else {
		//for i := 0; i > n; i-- {
		//	res = res / x
		//}
		n = -n
		for ; n > 0; n >>= 1 {
			// 快速幂算法
			if n&1 == 1 {
				res = res * x
			}
			x = x * x
		}
		return 1 / res
	}
}

func main() {

}
