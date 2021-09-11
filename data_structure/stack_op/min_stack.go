package stack_op

/*
剑指 Offer 30. 包含min函数的栈
定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。
*/

type MinStack struct {
	arr []int
	min int
}

/** initialize your data structure here. */
func Constructor1() MinStack {
	return MinStack{
		arr: []int{},
		min: 0,
	}
}

func (this *MinStack) Push(x int) {
	if len(this.arr) == 0 {
		this.min = x
	}
	diff := x - this.min
	if diff < 0 {
		this.min = x
	}
	this.arr = append(this.arr, diff)
}

func (this *MinStack) Pop() {
	if len(this.arr) == 0 {
		return
	}
	topDiff := this.arr[len(this.arr)-1]
	if topDiff < 0 {
		this.min = this.min - topDiff
	}
	this.arr = this.arr[:len(this.arr)-1]
}

func (this *MinStack) Top() int {
	if len(this.arr) == 0 {
		return -1
	}
	topDiff := this.arr[len(this.arr)-1]
	if topDiff > 0 {
		return this.min + topDiff
	}
	return this.min
}

func (this *MinStack) Min() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
