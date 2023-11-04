package main

type trie struct {
	left, right *trie
}

func (t *trie) add(num int) {
	p := t
	for k := 30; k >= 0; k-- {
		if num>>k&1 == 0 {
			if p.left == nil {
				p.left = &trie{nil, nil}
			}
			p = p.left
		} else {
			if p.right == nil {
				p.right = &trie{nil, nil}
			}
			p = p.right
		}
	}
}

func (t *trie) find(num int) int {
	ans := 0
	p := t
	for k := 30; k >= 0; k-- {
		bit := num >> k & 1
		if bit == 0 {
			if p.right != nil {
				ans = ans*2 + 1
				p = p.right
			} else {
				ans = ans * 2
				p = p.left
			}
		} else {
			if p.left != nil {
				ans = ans*2 + 1
				p = p.left
			} else {
				ans = ans * 2
				p = p.right
			}
		}
	}
	return ans
}

func findMaximumXOR(nums []int) (x int) {
	n := len(nums)
	root := &trie{}
	for i := 1; i < n; i++ {
		root.add(nums[i-1])
		x = max421(x, root.find(nums[i]))
	}
	return
}

func max421(x, y int) int {
	if x > y {
		return x
	}
	return y
}
