/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */
package leetcode

// @lc code=start
func maxArea(height []int) int {
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	maxArea := 0
	for i, j := 0, len(height)-1; i < j; {
		_area := min(height[i], height[j]) * (j - i)
		maxArea = max(maxArea, _area)
		// 移动短板, 因为短的那一侧会限制盛水量, 而不能移动长版
		if height[i] < height[j] {
			// 左边为短板, 左边往右移
			i++
		} else {
			// 右边为短板, 右边左移
			j--
		}
	}

	return maxArea
}

// @lc code=end
