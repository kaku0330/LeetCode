package implementstackusingqueues

type MyStack struct {
	stack []int
	count int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.count++
	this.stack = append(this.stack, x)
}

func (this *MyStack) Pop() int {
	result := 0
	if this.count != 0 {
		result = this.stack[this.count-1]
		this.stack = this.stack[:this.count-1]
		this.count--
		return result
	}
	return result
}

func (this *MyStack) Top() int {
	return this.stack[this.count-1]
}

func (this *MyStack) Empty() bool {
	return this.count == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
