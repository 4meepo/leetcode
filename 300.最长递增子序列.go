/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长递增子序列
 */
package leetcode

// @lc code=start
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))

	rtn := 1
	for i, v := range nums {
		if i == 0 {
			dp[0] = 1
			continue
		}

		max := 1
		for j := 0; j < i; j++ {
			if nums[j] < v {
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
