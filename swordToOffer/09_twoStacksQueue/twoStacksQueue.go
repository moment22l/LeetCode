package main

type CQueue struct {
	instack, outstack []int
}

func Constructor() CQueue {
	return CQueue{[]int{}, []int{}}
}

func (this *CQueue) AppendTail(value int) {
	this.instack = append(this.instack, value)
}

func (this *CQueue) DeleteHead() int {
	// 最简单的方法: 将instack中除了栈底元素全部都扔进outstack，取出栈底元素后再将outstack的数按顺序扔回来
	// if len(this.instack) == 0 {
	// 	return -1
	// }
	// var head int
	// for len(this.instack) > 1 {
	// 	this.outstack = append(this.outstack, this.instack[len(this.instack)-1])
	// 	this.instack = this.instack[:len(this.instack)-1]
	// }
	// head = this.instack[len(this.instack)-1]
	// this.instack = this.instack[:len(this.instack)-1]
	// for len(this.outstack) > 0 {
	// 	this.instack = append(this.instack, this.outstack[len(this.outstack)-1])
	// 	this.outstack = this.outstack[:len(this.outstack)-1]
	// }
	// return head

	if len(this.outstack) == 0 && len(this.instack) == 0 {
		return -1
	}
	// 如果outstack中没有元素，则将instack中的元素扔进来
	if len(this.outstack) == 0 {
		for len(this.instack) > 0 {
			this.outstack = append(this.outstack, this.instack[len(this.instack)-1])
			this.instack = this.instack[:len(this.instack)-1]
		}
	}
	// 取出outstack的栈顶元素
	head := this.outstack[len(this.outstack)-1]
	this.outstack = this.outstack[:len(this.outstack)-1]
	return head
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

func main() {

}
