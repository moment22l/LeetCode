package main

// 在方法一中有用到, 所以在本文件题解中也没用
func max_2003(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func smallestMissingValueSubtree(parents []int, nums []int) []int {
	n := len(parents)
	children := make([][]int, n)
	for i := 1; i < n; i++ {
		children[parents[i]] = append(children[parents[i]], i)
	}

	geneSet, visited := make(map[int]bool), make([]bool, n)
	var dfs func(int)
	dfs = func(node int) {
		if visited[node] {
			return
		}
		visited[node] = true
		geneSet[nums[node]] = true
		for _, child := range children[node] {
			dfs(child)
		}
	}

	ans, node, iNode := make([]int, n), -1, 1
	for i := 0; i < n; i++ {
		ans[i] = 1
		if nums[i] == 1 {
			node = i
		}
	}
	for node != -1 {
		dfs(node)
		for geneSet[iNode] {
			iNode++
		}
		ans[node] = iNode
		node = parents[node]
	}
	return ans
}

// 找到一个数组中第一个缺失的正整数的方法（在我的超时算法中用到, 所以没用在最后的题解中
func findFirstMissingNumber(nums []int) int {
	temp := make([]int, len(nums))
	copy(temp, nums)
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] > 0 && nums[i] <= n {
			if temp[nums[i]-1] > 0 {
				temp[nums[i]-1] = -1 * temp[nums[i]-1]
			}
		}
	}
	for i := 0; i < n; i++ {
		if temp[i] > 0 {
			return i + 1
		}
	}
	return n + 1
}
