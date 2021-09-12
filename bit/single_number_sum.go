package bit

/*
给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。
*/
//解法一：针对每个比特位新建状态机器
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

//解法2，统计每个bit，然后模3
func singleNumber2(nums []int) int {
	var arr [32]int
	for _, num := range nums {
		base := 1
		for i := 0; i < 32; i++ {
			if num&base > 0 {
				arr[i]++
			}
			base <<= 1
		}
	}
	ret := 0
	for i, n := range arr {
		if n%3 == 1 {
			ret |= 1 << i
		}
	}
	return ret
}

//一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。请写程序找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。
func singleNumbers(nums []int) []int {
	res := 0
	for _, num := range nums {
		res ^= num
	}
	mask := 1
	for mask&res == 0 {
		mask <<= 1
	}
	//分组异或
	//两个数字不同，那么异或结果一定不为0，那么用异或结果作为分组依据，按以下规则分组
	//首先容易知道异或结果为1代表两个数在该位置不同
	//那么以所有数字异或结果的最低位1为标准，该位为0为一组，不为0为另一组,就一定能把两个需要的数分到两个组
	var r1, r2 int
	for _, num := range nums {
		if mask&num > 0 {
			r1 ^= num
		} else {
			r2 ^= num
		}
	}
	return []int{r1, r2}
}
