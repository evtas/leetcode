package _55

import "math"

type MinStack struct {
	Stack    []int
	MinStack []int
}

func Constructor() MinStack {
	return MinStack{
		Stack:    []int{},
		MinStack: []int{math.MaxInt},
	}
}

func (this *MinStack) Push(val int) {
	this.Stack = append(this.Stack, val)
	this.MinStack = append(this.MinStack, min(this.MinStack[len(this.MinStack)-1], val))
}

func (this *MinStack) Pop() {
	this.Stack = this.Stack[:len(this.Stack)-1]
	this.MinStack = this.MinStack[:len(this.MinStack)-1]
}

func (this *MinStack) Top() int {
	return this.Stack[len(this.Stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.MinStack[len(this.MinStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
