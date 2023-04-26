package main

// p1指向第一个为1的索引, p2指向第一个为2的索引
// 时间O(n) 空间O(1)
//func sortColors(nums []int) {
//	p1, p2 := 0, 0
//	for i, num := range nums {
//		if num == 0 {
//			nums[i], nums[p1] = nums[p1], nums[i]
//			if p1 < p2 {
//				nums[i], nums[p2] = nums[p2], nums[i]
//			}
//			p1++
//			p2++
//		} else if num == 1 {
//			nums[i], nums[p2] = nums[p2], nums[i]
//			p2++
//		}
//	}
//}

// p0, p2分别指向0, 2该去的位置
// 时间O(n) 空间O(1)
func sortColors(nums []int) {
	p0, p2 := 0, len(nums)-1
	for i := 0; i <= p2; i++ {
		for ; i <= p2 && nums[i] == 2; p2-- {
			nums[i], nums[p2] = nums[p2], nums[i]
		}
		if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			p0++
		}
	}
}
