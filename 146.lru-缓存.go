/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存
 */
package main

import "container/list"

// @lc code=start

type LRUCache struct {
	capacity int

	lru   *list.List
	cache map[int]*list.Element
}

type entry struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		lru:      list.New(),
		cache:    make(map[int]*list.Element, capacity),
	}

}

func (this *LRUCache) Get(key int) int {
	if e, ok := this.cache[key]; ok {
		this.lru.MoveToFront(e)
		return e.Value.(*entry).value
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {

	if val, ok := this.cache[key]; ok {
		// exists
		this.lru.MoveToFront(val)
		val.Value.(*entry).value = value
		return
	}

	// not hit
	e := this.lru.PushFront(&entry{key: key, value: value})
	this.cache[key] = e

	if this.lru.Len() > this.capacity {
		// evit
		// remove oldest
		oldest := this.lru.Back()
		if oldest != nil {
			this.removeElement(oldest)
		}
	}

}

func (this *LRUCache) removeElement(e *list.Element) {
	this.lru.Remove(e)
	delete(this.cache, e.Value.(*entry).key)
}

// ["LRUCache","put",  "put","get","put", "get","put","get","get","get"]
// [[2]        ,[1,1],[2,2], [1],  [3,3],  [2],  [4,4],[1],  [3],  [4]]

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
