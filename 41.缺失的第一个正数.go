/*
 * @lc app=leetcode.cn id=41 lang=golang
 *
 * [41] 缺失的第一个正数
 */
package leetcode

// @lc code=start
func firstMissingPositive(nums []int) int {
	m := make(map[int]struct{})

	for _, v := range nums {
		m[v] = struct{}{}
	}

	for i := 1; ; i++ {
		if _, ok := m[i]; !ok {
			return i
		}
	}
}

// @lc code=end
