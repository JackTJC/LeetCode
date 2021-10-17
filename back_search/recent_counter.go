package backtrack_search

import "sort"

type RecentCounter struct {
	arr []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		arr: make([]int, 0),
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.arr = append(this.arr, t)
	ret := sort.Search(len(this.arr), func(i int) bool {
		return i >= t-3000
	})
	return len(this.arr) - ret
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
