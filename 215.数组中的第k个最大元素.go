/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */
package leetcode

import "container/heap"

// @lc code=start

type maxheap []int

func (m maxheap) Less(i, j int) bool {
	return m[i] > m[j]
}

func (m maxheap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m maxheap) Len() int {
	return len(m)
}

func (m *maxheap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *maxheap) Pop() interface{} {
	old := *m
	n := len(old)

	x := old[n-1]
	*m = old[:n-1]

	return x
}

func findKthLargest(nums []int, k int) int {
	hp := &maxheap{}

	for i := 0; i < len(nums); i++ {
		hp.Push(nums[i])
	}
	heap.Init(hp)

	var n int
	for i := 0; i < k; i++ {
		n = heap.Pop(hp).(int)
	}

	return n
}

// @lc code=end
