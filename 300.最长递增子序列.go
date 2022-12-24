/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长递增子序列
 */
package leetcode

// @lc code=start
func lengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	dp := make([]int, len(nums))

	var rtn = 0

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			dp[0] = 1
			continue
		}

		max := 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if dp[j]+1 > max {
					max = dp[j] + 1
				}
			}
		}
		dp[i] = max

		if max > rtn {
			rtn = max
		}
	}

	return rtn
}

// @lc code=end
