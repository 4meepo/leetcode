/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */

package leetcode

// @lc code=start
func removeDuplicates(nums []int) int {
	n := len(nums)
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			continue
		}

		if nums[i] == nums[i-1] {
			if i+1 < len(nums) {
				nums = append(nums[:i], nums[i+1:]...)
			} else {
				nums = nums[:i]
			}
			n--
			i--
		}

	}

	return n
}

// @lc code=end
