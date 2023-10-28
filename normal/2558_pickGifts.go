package main

import (
	"container/heap"
	"math"
)

type bigHeap []int

func (h *bigHeap) Len() int {
	return len(*h)
}

func (h *bigHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *bigHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *bigHeap) Push(v any) {
	*h = append(*h, v.(int))
}

func (h *bigHeap) Pop() any {
	n := h.Len()
	res := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return res
}

func pickGifts(gifts []int, k int) int64 {
	q := &bigHeap{}
	for _, gift := range gifts {
		q.Push(gift)
	}
	heap.Init(q)
	for k > 0 {
		gift := heap.Pop(q)
		heap.Push(q, int(math.Floor(math.Sqrt(float64(gift.(int))))))
		k--
	}
	var res int64 = 0
	for q.Len() > 0 {
		res += int64(q.Pop().(int))
	}
	return res
}
