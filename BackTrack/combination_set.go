package BackTrack

import (
	"sort"
)

type backTracker struct {
	candidates []int
	arr        []int
	result     [][]int
	sum        int
	target     int
	curIdx     int
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] > candidates[j]
	})
	bt := backTracker{
		candidates: candidates,
		arr:        make([]int, 0),
		result:     make([][]int, 0),
		sum:        0,
		target:     target,
		curIdx:     0,
	}
	bt.run()
	return bt.result
}

func (bt *backTracker) run() {
	if bt.sum > bt.target {
		return
	}
	if bt.sum == bt.target {
		r := make([]int, 0)
		r = append(r, bt.arr...)
		bt.result = append(bt.result, r)
		return
	}
	for i := bt.curIdx; i < len(bt.candidates); i++ {
		//select
		bt.arr = append(bt.arr, bt.candidates[i])
		bt.sum += bt.candidates[i]
		bt.curIdx = i
		//search
		bt.run()
		//pop
		bt.arr = bt.arr[0 : len(bt.arr)-1]
		bt.sum -= bt.candidates[i]
	}
}
