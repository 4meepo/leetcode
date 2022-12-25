/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */
package leetcode

import "sort"

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	counter := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		n := nums[i]
		if v, ok := counter[n]; ok {
			counter[n] = v + 1
		} else {
			counter[n] = 1
		}
	}
	type entry struct {
		K int
		V int
	}
	entries := make([]entry, 0, len(counter))
	for k, v := range counter {
		entries = append(entries, entry{k, v})
	}

	// sort
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].V > entries[j].V
	})

	var rtn []int
	for i := 0; i < k; i++ {
		rtn = append(rtn, entries[i].K)
	}
	return rtn

}

// @lc code=end
