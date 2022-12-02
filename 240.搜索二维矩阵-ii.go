/*
 * @lc app=leetcode.cn id=240 lang=golang
 *
 * [240] 搜索二维矩阵 II
 */
package main

import "sort"

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	for x, y := n-1, 0; x >= 0 && y < m; x, y = x-1, y+1 {
		cur := matrix[y][x]
		if target == cur {
			return true
		}
		if target < cur {
			index := sort.SearchInts(matrix[y], target)
			if index >= x {
				// not found
			} else if matrix[y][index] == target {
				return true
			}
			// not found
		} else if target > cur {
			for _y := y; _y < m; _y++ {
				if matrix[_y][x] == target {
					return true
				}
			}
		}
	}

	return false
}

// @lc code=end
