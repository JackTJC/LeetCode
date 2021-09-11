package backtrack_search

type permuter struct {
	nums     []int
	selected map[int]struct{}
	ret      [][]int
	cur      []int
}

func (p *permuter) run() {
	if len(p.selected) == len(p.nums) {
		p.ret = append(p.ret, p.cur)
	}
}

//n * (n-1) * (n-2) * ... * 1
func permute(nums []int) [][]int {
	selected := make(map[int]struct{})
	ret := make([][]int, 0)
	cur := make([]int, 0)
	//填充每一个位置
	for i := 0; i < len(nums); i++ {
		//选择一个数
		for j := 0; j < len(nums); j++ {
			if _, ok := selected[nums[j]]; ok {
				continue
			}
			cur = append(cur, nums[j])
		}
	}
	return ret
}
