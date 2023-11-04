package main

type Node117 struct {
	Val   int
	Left  *Node117
	Right *Node117
	Next  *Node117
}

func connect(root *Node117) *Node117 {
	start := root
	for start != nil {
		var nextStart, last *Node117
		handle := func(cur *Node117) {
			if cur == nil {
				return
			}
			if nextStart == nil {
				nextStart = cur
			}
			if last != nil {
				last.Next = cur
			}
			last = cur
		}
		for p := start; p != nil; p = p.Next {
			handle(p.Left)
			handle(p.Right)
		}
		start = nextStart
	}
	return root
}
