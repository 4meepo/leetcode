/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 */
package leetcode

// @lc code=start
func findTargetSumWays(nums []int, target int) int {

	count := 0
	var dfs func(idx int, sum int)
	dfs = func(idx int, sum int) {
		if idx == len(nums) {
			if sum == target {
				count++
			}
			return
		}

		dfs(idx+1, sum+nums[idx])
		dfs(idx+1, sum-nums[idx])
	}

	dfs(0, 0)

	return count
}

// @lc code=end
