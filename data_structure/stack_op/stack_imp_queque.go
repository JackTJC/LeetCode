package stack_op

type Stack interface {
	Pop() int
	Push(int)
	Len() int
}

type stackImpl struct {
	arr []int
}

func newStackImpl() *stackImpl {
	return &stackImpl{
		arr: []int{},
	}
}

func (s *stackImpl) Pop() int {
	if len(s.arr) == 0 {
		return -1
	}
	ret := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return ret
}

func (s *stackImpl) Push(value int) {
	s.arr = append(s.arr, value)
}

func (s *stackImpl) Len() int {
	return len(s.arr)
}

//两个栈实现队列
type CQueue struct {
	s1 Stack
	s2 Stack
}

//一致性约束：每次操作完成，只有s1有数据
func Constructor() CQueue {
	return CQueue{
		s1: newStackImpl(),
		s2: newStackImpl(),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.s1.Push(value)
}

func (this *CQueue) DeleteHead() int {
	if this.s1.Len() == 0 {
		return -1
	}
	for this.s1.Len() != 0 {
		this.s2.Push(this.s1.Pop())
	}
	ret := this.s2.Pop()
	for this.s2.Len() != 0 {
		this.s1.Push(this.s2.Pop())
	}
	return ret
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
