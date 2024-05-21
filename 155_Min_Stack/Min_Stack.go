package minstack

type MinStack struct {
	stack  []int
	top    int
	getMin []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	this.top = val
	if len(this.getMin) == 0 {
		this.getMin = append(this.getMin, val)
	} else {
		if this.getMin[len(this.getMin)-1] >= val {
			this.getMin = append(this.getMin, val)
		}
	}
}

func (this *MinStack) Pop() {
	if this.getMin[len(this.getMin)-1] == this.top {
		this.getMin = this.getMin[:len(this.getMin)-1]
	}
	if len(this.stack) != 0 {
		this.stack = this.stack[:len(this.stack)-1]
		if len(this.stack) != 0 {
			this.top = this.stack[len(this.stack)-1]
		} else {
			this.top = 0
		}
	}
}

func (this *MinStack) Top() int {
	return this.top
}

func (this *MinStack) GetMin() int {
	return this.getMin[len(this.getMin)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
