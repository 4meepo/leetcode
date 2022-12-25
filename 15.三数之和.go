/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */
package leetcode

import "sort"

// @lc code=start
func threeSum(nums []int) [][]int {
	var rtn [][]int

	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1

		target := -1 * nums[i]
		for left < right {
			sum := nums[left] + nums[right]
			if sum > target {
				right--
			} else if sum < target {
				left++
			} else {
				duplicated := false
				for _, v := range rtn {
					if v[0] == nums[i] && v[1] == nums[left] && v[2] == nums[right] {
						duplicated = true
						break
					}
				}
				if !duplicated {
					rtn = append(rtn, []int{nums[i], nums[left], nums[right]})
				}
				
				left++
			}
		}
	}

	return rtn
}

// @lc code=end
