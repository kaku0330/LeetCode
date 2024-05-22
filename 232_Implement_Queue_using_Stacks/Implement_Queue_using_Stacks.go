package implementqueueusingstacks

type MyQueue struct {
	queue []int
	count int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.queue = append(this.queue, x)
	this.count++
}

func (this *MyQueue) Pop() int {
	result := 0
	if this.count != 0 {
		result = this.queue[0]
		this.queue = this.queue[1:]
		this.count--
	}
	return result
}

func (this *MyQueue) Peek() int {
	return this.queue[0]
}

func (this *MyQueue) Empty() bool {
	return this.count == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
