package stack_op

/*
给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
求在该柱状图中，能够勾勒出来的矩形的最大面积。
*/

//构造一个栈底到栈定递增的单调栈
func largestRectangleArea(heights []int) int {
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)
	stack := make([]int, 0)
	ret := 0
	for idx, height := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > height {
			//pop
			cur := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			//cal
			ret = max(ret, (idx-stack[len(stack)-1]-1)*heights[cur])
		}
		stack = append(stack, idx)
	}
	return ret
}
