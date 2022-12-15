/*
 * @lc app=leetcode.cn id=719 lang=golang
 *
 * [719] 找出第 K 小的数对距离
 */
package leetcode

import (
	"container/heap"
	"math"
	"sort"
)

// @lc code=start
type heap719 struct {
	sort.IntSlice
}

func (h *heap719) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}
func (h *heap719) Pop() interface{} {
	n := len(h.IntSlice)
	x := h.IntSlice[n-1]
	h.IntSlice = h.IntSlice[:n-1]
	return x
}

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	hp := &heap719{make(sort.IntSlice, 0, len(nums))}
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			heap.Push(hp, int(math.Abs(float64(nums[i])-float64(nums[j]))))
		}
	}
	var rtn int
	for k > 0 {
		k--
		rtn = heap.Pop(hp).(int)
	}

	return rtn
}

// @lc code=end
