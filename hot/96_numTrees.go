package main

func numTrees(n int) int {
	G := make([]int, n+1)
	G[0], G[1] = 1, 1
	for i := 2; i < n+1; i++ {
		for j := 1; j < i+1; j++ {
			G[i] += G[j-1] * G[i-j]
		}
	}
	return G[n]
}
