/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */
package main

import "container/heap"

// @lc code=start

type Element struct {
	val   int
	index int
}

type maxHeap2 []*Element

func (h maxHeap2) Len() int {
	return len(h)
}

func (h maxHeap2) Less(i int, j int) bool {
	return h[i].val > h[j].val

}

func (h maxHeap2) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxHeap2) Push(x interface{}) {
	*h = append(*h, x.(*Element))
}

func (h *maxHeap2) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func maxSlidingWindow(nums []int, k int) []int {
	_k := k
	var rtn []int
	h := make(maxHeap2, 0, len(nums))

	for i := 0; i < len(nums); i++ {
		e := &Element{val: nums[i], index: i}
		heap.Push(&h, e)
		_k--
		if _k > 0 {
			continue
		}

		var max *Element
		for {
			max = h[0]
			if max.index >= i-k+1 {
				// 非窗口外的值
				break
			}

			// 窗口外的值. 继续pop then loop
			heap.Pop(&h)
		}

		rtn = append(rtn, max.val)
	}

	return rtn
}

// @lc code=end
