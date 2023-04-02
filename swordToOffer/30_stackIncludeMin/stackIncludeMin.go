package main

import (
	"fmt"
	"math"
)

type MinStack struct {
	s   []int
	min int
}

// Constructor initialize your data structure here.
func Constructor() MinStack {
	return MinStack{[]int{}, 0}
}

func (this *MinStack) Push(x int) {
	if len(this.s) == 0 {
		this.s = append(this.s, 0)
		this.min = x
	} else {
		this.s = append(this.s, x-this.min)
		this.min = int(math.Min(float64(this.min), float64(x)))
	}
}

func (this *MinStack) Pop() {
	if this.s[len(this.s)-1] < 0 {
		this.min = this.min - this.s[len(this.s)-1]
	}
	this.s = this.s[:len(this.s)-1]
}

func (this *MinStack) Top() int {
	if this.s[len(this.s)-1] > 0 {
		return this.s[len(this.s)-1] + this.min
	} else {
		return this.min
	}
}

func (this *MinStack) Min() int {
	return this.min
}

func main() {
	minS := Constructor()
	minS.Push(-2)
	minS.Push(0)
	minS.Push(-3)
	fmt.Println(minS.Min())
	minS.Pop()
	fmt.Println(minS.Top())
	fmt.Println(minS.Min())
}
