/*
* 用到最小栈数据结构的算法
* */
package minstack

func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0)
	ret := make([]int, len(temperatures))
	for idx, t := range temperatures {
		if len(stack) == 0 {
			stack = append(stack, idx)
			continue
		}
		if t <= temperatures[stack[len(stack)-1]] {
			stack = append(stack, idx)
		} else {
			for len(stack) > 0 && t > temperatures[stack[len(stack)-1]] {
				popV := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				ret[popV] = idx - popV
			}
			stack = append(stack, idx)
		}
	}
	return ret
}
