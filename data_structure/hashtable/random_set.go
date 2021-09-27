package hashtable

import "math/rand"

type RandomizedSet struct {
	hashMap map[int]int
	arr     []int
	length  int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		hashMap: make(map[int]int, 0),
		arr:     make([]int, 0),
		length:  0,
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.hashMap[val]; ok {
		return false
	}
	var idx int
	if len(this.arr) == this.length {
		this.arr = append(this.arr, val)
		idx = len(this.arr) - 1
	} else {
		this.arr[this.length] = val
		idx = this.length
	}
	this.hashMap[val] = idx
	this.length++
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.hashMap[val]
	if !ok {
		return false
	}
	delete(this.hashMap, val)
	this.arr[idx] = this.arr[this.length-1]
	if idx != this.length-1 {
		this.hashMap[this.arr[idx]] = idx
	}
	this.length--
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	return this.arr[rand.Intn(this.length)]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
