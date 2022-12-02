/*
 * @lc app=leetcode.cn id=240 lang=golang
 *
 * [240] 搜索二维矩阵 II
 */
package main

// @lc code=start
// z字形搜索
// https://leetcode.cn/problems/search-a-2d-matrix-ii/solutions/1062538/sou-suo-er-wei-ju-zhen-ii-by-leetcode-so-9hcx/
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])

	for x, y := n-1, 0; x >= 0 && y < m; {
		num := matrix[y][x]
		if num == target {
			return true
		}

		if num < target {
			y++
		} else if num > target {
			x--
		}
	}

	return false
}

// @lc code=end
