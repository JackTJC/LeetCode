package array

import (
	"fmt"
	"sort"
	"testing"
)

func TestFR(t *testing.T) {
	t.Log(findRepeatNumber([]int{3, 2, 1, 0, 0}))
}

func TestSearch(t *testing.T) {
	arr := []int{1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5, 5, 5, 6, 7, 8, 9, 10}
	t.Log(search(arr, 1))
	t.Log(search(arr, 2))
	t.Log(search(arr, 3))
	t.Log(search(arr, 4))
	t.Log(search(arr, 5))
	t.Log(search(arr, 6))
	t.Log(search(arr, 7))
	t.Log(search(arr, 8))
	t.Log(search(arr, 9))
	t.Log(search(arr, 11))
	t.Log(search(arr, 0))
}

func TestMN(t *testing.T) {
	t.Log(missingNumber([]int{0, 1, 2, 3, 4, 5, 6, 7, 9}))
	t.Log(missingNumber([]int{0, 2}))
	t.Log(missingNumber([]int{0, 1, 3}))
	t.Log(missingNumber([]int{0}))
	t.Log(missingNumber([]int{1}))
	t.Log(missingNumber([]int{1, 2}))
}

func TestMyMap(t *testing.T) {
	m := &MyMap{}
	data := []int{100, 45, 67, 23, 45, 64, 76, 78, 98, 44, 23, 44, 55, 66, 77, 100}
	m.build(data)
	m.transverse()
	sort.Slice(data, func(i, j int) bool {
		return data[i] < data[j]
	})
	fmt.Println(data)
}

func TestSpiral(t *testing.T) {
	arr1 := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	arr2 := append(arr1, []int{13, 14, 15, 16})
	_ = spiralOrder(arr2)
	res1 := spiralOrder(arr1)
	t.Log(res1)
}

func TestDengcha(t *testing.T) {
	longestArithSeqLength([]int{3, 6, 9, 12})
}
