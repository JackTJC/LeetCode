package hashtable

import "container/list"

type LRUCache struct {
	l    *list.List
	m    map[int]*list.Element
	c    int
	data map[int]int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		l:    list.New(),
		m:    make(map[int]*list.Element),
		c:    capacity,
		data: make(map[int]int),
	}
}

func (this *LRUCache) Get(key int) int {
	if ele, ok := this.m[key]; ok {
		this.l.Remove(ele)
		this.l.PushFront(ele.Value.(int))
		this.m[key] = this.l.Front()
		return this.data[ele.Value.(int)]
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	defer func() {
		this.data[key] = value
	}()
	if this.Get(key) != -1 {
		ele := this.m[key]
		this.l.Remove(ele)
		this.l.PushFront(ele.Value.(int))
		this.m[key] = this.l.Front()
		return
	}
	if this.l.Len() >= this.c {
		//drop
		ele := this.l.Back()
		this.l.Remove(ele)
		delete(this.m, ele.Value.(int))
		delete(this.data, ele.Value.(int))
	}
	this.l.PushFront(key)
	this.m[key] = this.l.Front()
}
