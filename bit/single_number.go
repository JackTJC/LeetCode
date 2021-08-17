package bit

/*
给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。
*/
func singleNumber(nums []int) int {
	machines := make([]machine, 64)
	for i := 0; i < len(nums); i++ {
		idx := 63
		for nums[i] != 0 {
			machines[idx].transfer(getCurInput(nums[i]))
			nums[i] <<= 1
			idx--
		}
	}
	var ret int
	base := 1
	for i := 0; i < 64; i++ {
		if machines[i].state == acc1 {
			ret += base
		}
		base <<= 1
	}
	return ret
}

//构造状态机，输入三个1回到初始状态

const (
	start = 0
	acc1  = 1
	acc2  = 2
)

type machine struct {
	state int
}

func (m *machine) transfer(input int) {
	if input == 0 {
		return
	}
	switch m.state {
	case start:
		m.state = acc1
	case acc1:
		m.state = acc2
	case acc2:
		m.state = start
	}
}

func getCurInput(x int) int {
	if x < 0 {
		return 1
	}
	return 0
}
