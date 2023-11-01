package main

func maximumInvitations(favorite []int) int {
	n := len(favorite)
	indeg := make([]int, n) // 入度
	for i := 0; i < n; i++ {
		indeg[favorite[i]]++
	}

	f := make([]int, n) // 每个节点的最长游走路径
	for i := 0; i < n; i++ {
		f[i] = 1
	}

	q := make([]int, 0) // 入度为0的节点队列
	for i := 0; i < n; i++ {
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}

	used := make([]bool, n) // 节点是否已经被遍历过了
	for len(q) != 0 {
		u := q[0]
		used[u] = true
		q = q[1:]
		v := favorite[u]
		f[v] = max_2127(f[v], f[u]+1)
		indeg[v]--
		if indeg[v] == 0 {
			q = append(q, v)
		}
	}

	ring, total := 0, 0 // 最长环和最长双向游走路径
	for i := 0; i < n; i++ {
		if !used[i] {
			j := favorite[i]
			if favorite[j] == i {
				// 环为2
				total += f[i] + f[j]
				used[i], used[j] = true, true
			} else {
				u, count := i, 0
				for {
					count++
					u = favorite[u]
					used[u] = true
					if u == i {
						break
					}
				}
				ring = max_2127(ring, count)
			}
		}
	}
	return max_2127(ring, total)
}

func max_2127(x, y int) int {
	if x > y {
		return x
	}
	return y
}
