package _22

type MyCircularQueue struct {
	q    []int
	n    int
	head int
	tail int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		q:    make([]int, k+1),
		n:    k + 1,
		head: 0,
		tail: 0,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if (this.tail+1)%this.n == this.head {
		return false
	}
	this.q[this.tail] = value
	this.tail += 1
	this.tail %= this.n
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.head += 1
	this.head %= this.n
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.q[this.head]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.q[(this.tail+this.n-1)%this.n]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.head == this.tail
}

func (this *MyCircularQueue) IsFull() bool {
	return (this.tail+1)%this.n == this.head
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
