package main

// 经典超时
//func maxArea(height []int) int {
//	if len(height) <= 1 {
//		return 0
//	}
//	if len(height) == 2 {
//		return int(math.Min(float64(height[0]), float64(height[1])))
//	}
//	heightmap := make(map[int]int)
//	ans := 0
//	for i, h := range height {
//		if _, ok := heightmap[h]; !ok {
//			heightmap[h] = i
//		}
//		for key := range heightmap {
//			area := (i - heightmap[key]) * int(math.Min(float64(key), float64(h)))
//			ans = int(math.Max(float64(ans), float64(area)))
//		}
//	}
//	return ans
//}

func maxArea(height []int) int {
	ans := 0
	for left, right := 0, len(height)-1; left < right; {
		//area := (right - left) * int(math.Min(float64(height[left]), float64(height[right])))
		//ans = int(math.Max(float64(area), float64(ans)))
		var area int
		if height[left] <= height[right] {
			area = (right - left) * height[left]
			left++
		} else {
			area = (right - left) * height[right]
			right--
		}
		if area > ans {
			ans = area
		}
	}
	return ans
}
