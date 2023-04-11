package main

// 快排, 并利用随机数进行加速, 使选择算法的时间变为线性, 时间O(n), 空间O(logn)
//func findKthLargest(nums []int, k int) int {
//	rand.Seed(time.Now().UnixNano())
//	return quickSort(nums, 0, len(nums)-1, k)
//}
//
//func quickSort(nums []int, l, r int, k int) int {
//	index := randomPartition(nums, l, r)
//	if index == k-1 {
//		return nums[index]
//	} else if index < k-1 {
//		return quickSort(nums, index+1, r, k)
//	} else {
//		return quickSort(nums, l, index-1, k)
//	}
//}
//
//func randomPartition(nums []int, l, r int) int {
//	i := rand.Int()%(r-l+1) + l
//	nums[i], nums[l] = nums[l], nums[i]
//	return partition(nums, l, r)
//}
//
//func partition(nums []int, l, r int) int {
//	pivot := nums[l]
//	i := l
//	for j := l + 1; j <= r; j++ {
//		if nums[j] >= pivot {
//			i++
//			nums[i], nums[j] = nums[j], nums[i]
//		}
//	}
//	nums[i], nums[l] = nums[l], nums[i]
//	return i
//}

// 堆排序 时间O(n+klogn)=O(nlogn) 空间O(logn)(递归的栈空间代价)
func findKthLargest(nums []int, k int) int {
	tail := len(nums)
	buildMaxHeap(nums, tail)
	for i := len(nums) - 1; i > len(nums)-k; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		tail--
		maxHeapify(nums, 0, tail)
	}
	return nums[0]
}

func buildMaxHeap(nums []int, size int) {
	for i := size / 2; i >= 0; i-- {
		maxHeapify(nums, i, size)
	}
}

func maxHeapify(nums []int, i, size int) {
	l, r, largest := 2*i+1, 2*i+2, i
	if l < size && nums[l] > nums[largest] {
		largest = l
	}
	if r < size && nums[r] > nums[largest] {
		largest = r
	}
	if i != largest {
		nums[i], nums[largest] = nums[largest], nums[i]
		maxHeapify(nums, largest, size)
	}
}
