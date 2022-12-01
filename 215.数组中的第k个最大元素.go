/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */
package main

import "container/heap"

// @lc code=start

// maxHeap 大顶堆
type maxHeap []int

func (h maxHeap) Len() int {
	return len(h)
}
func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h maxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	rtn := old[n-1]
	*h = old[:n-1]
	return rtn
}

func findKthLargest(nums []int, k int) int {

	hp := make(maxHeap, len(nums))

	for i, v := range nums {
		hp[i] = v
	}

	heap.Init(&hp)

	var rtn *int
	for k > 0 {
		n := heap.Pop(&hp).(int)
		rtn = &n
		k--
	}

	return *rtn
}

// @lc code=end
